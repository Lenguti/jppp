package v1

import "github.com/rs/zerolog"

type controller struct {
	cfg Config
	log zerolog.Logger
}
