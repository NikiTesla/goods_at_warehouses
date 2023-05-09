package database

import (
	"database/sql"

	lamodatest "github.com/NikiTesla/lamoda_test"
)

type DataBase interface {
	CreateGood(lamodatest.Good) error
	AddGood(int, int, int) error
	ReserveGood(int, int, int) error
	CancelGoodReservation(int, int, int) error
	CreateWarehouse(lamodatest.Warehouse) error
	GetAmount(int, int) (int, error)
}

type PostgresDB struct {
	DB *sql.DB
}
