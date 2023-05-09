package jsonrpc

import (
	"fmt"

	lamodatest "github.com/NikiTesla/lamoda_test"
)

type MockDB struct {
	goods []lamodatest.Good
}

func (m *MockDB) CreateGood(good lamodatest.Good) error {
	for _, v := range m.goods {
		if good.Code == v.Code {
			return fmt.Errorf("good alrady exists")
		}
	}
	m.goods = append(m.goods, good)

	return nil
}

func (m *MockDB) AddGood(good_code, warehouse_id, amount int) error {
	return nil
}

func (m *MockDB) ReserveGood(good_code, warehouse_id, amount int) error {
	return nil
}

func (m *MockDB) CancelGoodReservation(good_code, warehouse_id, amount int) error {
	return nil
}

func (m *MockDB) CreateWarehouse(warehouse lamodatest.Warehouse) error {
	return nil
}

func (m *MockDB) GetAmount(good_code, warehouse_id int) (int, error) {
	return 0, nil
}
