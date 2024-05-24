package environment

import (
	"fmt"

	"github.com/jackc/pgx"
)

type Environment struct {
	Config *Config
	DB     *pgx.Conn
}

func NewEnvironment(configFile string) (*Environment, error) {
	cfg, err := NewConfig(configFile)
	if err != nil {
		return nil, fmt.Errorf("cannot create new config, err: %w", err)
	}

	db, err := NewDatabase(cfg.DBConfig)
	if err != nil {
		return nil, fmt.Errorf("cannot create database connection, err: %w", err)
	}

	return &Environment{
		Config: cfg,
		DB:     db,
	}, nil
}
