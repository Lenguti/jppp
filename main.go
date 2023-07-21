package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lenguti/jppp/app/api/handlers"
	v1 "github.com/lenguti/jppp/app/api/handlers/v1"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

func main() {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := v1.NewConfig()
	if err != nil {
		log.Error().Err(err).Msg("Unable to create new config.")
		os.Exit(1)
	}

	srv := &http.Server{
		Addr:    ":8000",
		Handler: handlers.NewV1Handler(log, cfg),
	}

	go func() {
		log.Info().Msg("Starting web server.")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Unable to start webserver.")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	sig := <-quit
	log.Error().Str("signal", sig.String()).Msg("Received signal, shutting down server.")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Error shutting down server.")
	}

	log.Info().Msg("Server gracefully shut down.")
}
