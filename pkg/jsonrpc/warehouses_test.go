package jsonrpc

// import (
// 	"fmt"
// 	"testing"

// )

// func TestWarehouseCreate(t *testing.T) {
// 	mock := &Warehouses{db: &MockDB{}}

// 	args := [][]goods_at_warehouses.Warehouse{
// 		{{Name: "Main", Availability: true}, {Name: "Old", Availability: false}},
// 		{{}},
// 	}
// 	expected := []string{
// 		fmt.Sprintf("Created warehouses: %v", args[0]),
// 		"Created warehouses: []",
// 		"",
// 	}

// 	reply := &Reply{}
// 	for i := 0; i < len(args); i++ {
// 		if err := mock.Create(args[i], reply); err != nil {
// 			t.Errorf("cannot create warehouses %v, error: %s", args, err.Error())
// 		}
// 		if reply.Data != expected[i] {
// 			t.Errorf("Expected: %s. Got: %s", expected[i], reply.Data)
// 		}
// 		reply.Data = ""
// 	}
// }

// func TestGetAmount(t *testing.T) {
// 	mock := &Warehouses{db: &MockDB{
// 		waregouseGoods: []WarehouseGood{
// 			{goodCode: 123, warehouseID: 1, availableAmount: 100},
// 			{goodCode: 121, warehouseID: 1, availableAmount: 200},
// 		},
// 	}}

// 	args := []map[string]int{
// 		{"goodCode": 123, "warehouseID": 1},
// 		{"goodCode": 122, "warehouseID": 2},
// 		{},
// 	}
// 	expected := []map[string]string{
// 		{"error": "", "reply": "amount: 100"},
// 		{"error": "there is no 122 goods at the 2 warehouse", "reply": ""},
// 		{"error": "request is incorrect, good_code is not presented", "reply": ""},
// 	}

// 	reply := &Reply{}
// 	for i := 0; i < len(args); i++ {
// 		if err := mock.GetAmount(args[i], reply); err != nil {
// 			if err.Error() != expected[i]["error"] {
// 				t.Errorf("Expected error: %s. Got: %s", expected[i]["error"], reply.Data)
// 			}
// 		}
// 		if reply.Data != expected[i]["reply"] {
// 			t.Errorf("Expected: %s. Got: %s", expected[i]["reply"], reply.Data)
// 		}
// 		reply.Data = ""
// 	}
// }
