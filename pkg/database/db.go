package database

import (
	"github.com/NikiTesla/goods_at_warehouses/pkg/core"
	"github.com/jackc/pgx"
)

type DataBase interface {
	CreateGood(good core.Good) error
	AddGood(goodCode int, warehouseID int, amount int) error
	ReserveGood(goodCode int, warehouseID int, amount int) error
	CancelGoodReservation(goodCode int, warehouseID int, amount int) error
	CreateWarehouse(warehouse core.Warehouse) error
	GetAmount(goodCode int, warehouseID int) (int, error)
}

type PostgresDB struct {
	DB *pgx.Conn
}
