package v1

import (
	"fmt"
	"os"
)

// Config - represents configurtion for v1 services.
type Config struct {
	DBName string
	DBPass string
	DBUser string
}

// NewConfig - returns an new configurtion initialized with environment variables.
func NewConfig() (Config, error) {
	var (
		cfg Config
		err error
	)
	cfg, err = cfg.parseEnv()
	if err != nil {
		return cfg, fmt.Errorf("new config: unable to parse environment config: %w", err)
	}

	return cfg, nil
}

func (c Config) parseEnv() (Config, error) {
	var (
		dbName = os.Getenv("DB_NAME")
		dbUser = os.Getenv("DB_USER")
		dbPass = os.Getenv("DB_PASS")
	)

	switch "" {
	case dbUser, dbPass, dbName:
		return c, fmt.Errorf("parse env: invalid config provided")
	}

	c.DBName = dbName
	c.DBPass = dbPass
	c.DBUser = dbUser
	return c, nil
}
