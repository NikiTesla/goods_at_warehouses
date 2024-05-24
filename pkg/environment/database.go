package environment

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx"
)

func NewDatabase(cfg DBConfig) (*pgx.Conn, error) {
	config := pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     uint16(cfg.Port),
		Database: cfg.DBname,
		User:     cfg.Username,
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	db, err := pgx.Connect(config)
	if err != nil {
		return nil, fmt.Errorf("cannot connect database, err: %s", err)
	}
	if err := db.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("cannot ping database, err: %s", err)
	}
	return db, nil
}
