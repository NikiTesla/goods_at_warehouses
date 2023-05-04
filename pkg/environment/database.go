package environment

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func NewDataBase(cfg DBConfig) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, os.Getenv("POSTGRES_PASSWORD"), cfg.DBname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("can't connect database, err: %s", err.Error())
	}
	log.Printf("connected database, ping: error is %v", db.Ping())

	return db, nil
}
