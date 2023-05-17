package database

import (
	"github.com/NikiTesla/goods_at_warehouses"
	"github.com/jackc/pgx"
)

type DataBase interface {
	CreateGood(good goods_at_warehouses.Good) error
	AddGood(goodCode int, warehouseID int, amount int) error
	ReserveGood(goodCode int, warehouseID int, amount int) error
	CancelGoodReservation(goodCode int, warehouseID int, amount int) error
	CreateWarehouse(warehouse goods_at_warehouses.Warehouse) error
	GetAmount(goodCode int, warehouseID int) (int, error)
}

type PostgresDB struct {
	DB *pgx.Conn
}
