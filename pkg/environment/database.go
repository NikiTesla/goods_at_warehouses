package environment

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx"
)

func NewDataBase(cfg DBConfig) (*pgx.Conn, error) {
	config := pgx.ConnConfig{
		Host:     cfg.Host,
		Port:     uint16(cfg.Port),
		Database: cfg.DBname,
		User:     cfg.Username,
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	db, err := pgx.Connect(config)
	if err != nil {
		return nil, fmt.Errorf("can't connect database, err: %s", err)
	}
	if err := db.Ping(context.Background()); err != nil {
		log.Fatalf("Cannot connect to database, error: %s\n", err)
	}

	return db, nil
}
