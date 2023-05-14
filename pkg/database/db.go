package database

import (
	lamodatest "github.com/NikiTesla/lamoda_test"
	"github.com/jackc/pgx"
)

type DataBase interface {
	CreateGood(good lamodatest.Good) error
	AddGood(goodCode int, warehouseID int, amount int) error
	ReserveGood(goodCode int, warehouseID int, amount int) error
	CancelGoodReservation(goodCode int, warehouseID int, amount int) error
	CreateWarehouse(warehouse lamodatest.Warehouse) error
	GetAmount(goodCode int, warehouseID int) (int, error)
}

type PostgresDB struct {
	DB *pgx.Conn
}
