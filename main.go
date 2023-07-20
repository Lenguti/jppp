package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	DBPass  string
	DBTable string
	DBUser  string
}

func NewConfig() (Config, error) {
	var (
		cfg Config
		err error
	)
	cfg, err = cfg.ParseEnv()
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (c Config) ParseEnv() (Config, error) {
	var (
		dbUser  = os.Getenv("DB_USER")
		dbPass  = os.Getenv("DB_PASS")
		dbTable = os.Getenv("DB_TABLE")
	)

	switch "" {
	case dbUser, dbPass, dbTable:
		return Config{}, fmt.Errorf("parse env: invalid config provided")
	}

	return Config{
		DBPass:  dbPass,
		DBTable: dbTable,
		DBUser:  dbUser,
	}, nil
}

type Router struct {
	cfg Config
}

func main() {
	cfg, err := NewConfig()
	if err != nil {
		log.Println("unable to create new config")
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr: ":8000",
	}

	r := Router{
		cfg: cfg,
	}

	http.HandleFunc("/", r.root)
	http.HandleFunc("/dbping", r.dbping)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("err")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	sig := <-quit
	log.Fatal("Received signal, shutting down server: ", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Error shutting down server.")
	}

	log.Println("Server gracefully shut down.")
}

func (rr Router) root(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello world"))
}

func (rr Router) dbping(w http.ResponseWriter, r *http.Request) {
	q := make(url.Values)
	q.Set("sslmode", "disable")
	q.Set("timezone", "utc")

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(rr.cfg.DBUser, rr.cfg.DBPass),
		Host:     "db:5432",
		Path:     rr.cfg.DBTable,
		RawQuery: q.Encode(),
	}

	log.Println("db string: ", u.String())
	if _, err := sqlx.Connect("postgres", u.String()); err != nil {
		log.Println("unable to connect db")
		log.Fatal(err)
	}

	_, _ = w.Write([]byte("DB OK"))
}
