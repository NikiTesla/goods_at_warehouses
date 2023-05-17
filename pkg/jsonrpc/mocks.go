package jsonrpc

import (
	"fmt"

	"github.com/NikiTesla/goods_at_warehouses"
)

type WarehouseGood struct {
	goodCode        int
	warehouseID     int
	availableAmount int
	reservedAmount  int
}

type MockDB struct {
	goods          []goods_at_warehouses.Good
	warehouses     []goods_at_warehouses.Warehouse
	waregouseGoods []WarehouseGood
}

func (m *MockDB) CreateGood(good goods_at_warehouses.Good) error {
	if good.Name == "" || good.Code == 0 || good.Amount < 0 {
		return fmt.Errorf("check constraint")
	}
	for _, v := range m.goods {
		if good.Code == v.Code {
			return fmt.Errorf("good alrady exists")
		}
	}
	m.goods = append(m.goods, good)

	return nil
}

func (m *MockDB) AddGood(goodCode, warehouseID, amount int) error {
	var availableAmount int
	for _, v := range m.goods {
		if v.Code == goodCode {
			availableAmount = v.Amount
			if amount <= availableAmount {
				return nil
			}
		}
	}
	return fmt.Errorf("there is not enough %d good in %d warehouse. Available only %d",
		goodCode, warehouseID, availableAmount)
}

func (m *MockDB) ReserveGood(goodCode, warehouseID, amount int) error {
	var availableAmount int
	for _, v := range m.waregouseGoods {
		if v.goodCode == goodCode && v.warehouseID == warehouseID {
			availableAmount = v.availableAmount
			if amount <= availableAmount {
				availableAmount -= amount
				v.reservedAmount += amount
				return nil
			}
		}
	}
	return fmt.Errorf("there is not enough %d good in %d warehouse. Available only %d",
		goodCode, warehouseID, availableAmount)
}

func (m *MockDB) CancelGoodReservation(goodCode, warehouseID, amount int) error {
	var reservedAmount int
	for i, v := range m.waregouseGoods {
		if v.goodCode == goodCode && v.warehouseID == warehouseID {
			reservedAmount = v.reservedAmount
			if amount <= reservedAmount {
				m.waregouseGoods[i].reservedAmount -= amount
				m.waregouseGoods[i].availableAmount += amount
				return nil
			}
		}
	}
	return fmt.Errorf("there is not enough reserved %d goods at the %d warehouse. Reserved only %d",
		goodCode, warehouseID, reservedAmount)
}

func (m *MockDB) CreateWarehouse(warehouse goods_at_warehouses.Warehouse) error {
	if warehouse.Name == "" {
		return fmt.Errorf("check constraint")
	}

	for _, v := range m.warehouses {
		if warehouse.Name == v.Name {
			return fmt.Errorf("warehouse already exists")
		}
	}
	m.warehouses = append(m.warehouses, warehouse)

	return nil
}

func (m *MockDB) GetAmount(goodCode, warehouseID int) (int, error) {
	for _, v := range m.waregouseGoods {
		if goodCode == v.goodCode && warehouseID == v.warehouseID {
			return v.availableAmount, nil
		}
	}
	return 0, fmt.Errorf("there is no %d goods at the %d warehouse", goodCode, warehouseID)
}
