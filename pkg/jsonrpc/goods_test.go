package jsonrpc

// import (
// 	"fmt"
// 	"reflect"
// 	"testing"

// 	lamodatest "github.com/NikiTesla/lamoda_test"
// )

// func TestGoodCreate(t *testing.T) {
// 	mock := &Goods{db: &MockDB{}}

// 	args := [][]lamodatest.Good{
// 		{{Name: "Coffee", Code: 123, Size: 1.1, Amount: 100}, {Name: "Water", Code: 121, Size: 5, Amount: 200}},
// 		{{}},
// 	}
// 	expected := []string{
// 		fmt.Sprintf("Created goods: %v", args[0]),
// 		"Created goods: []",
// 	}

// 	reply := &[]lamodatest.Good{}
// 	for i := 0; i < len(args); i++ {
// 		if err := mock.Create(args[i], reply); err != nil {
// 			t.Errorf("cannot create goods %v, error: %s", args, err.Error())
// 		}
// 		if !reflect.DeepEqual(reply, expected[i]) {
// 			t.Errorf("Expected: %v. Got: %v", expected[i], reply)
// 		}
// 		reply = &[]lamodatest.Good{}
// 	}
// }

// func TestGoodAdd(t *testing.T) {
// 	mock := &Goods{db: &MockDB{
// 		goods: []lamodatest.Good{
// 			{Name: "Coffee", Code: 123, Size: 1.1, Amount: 100},
// 			{Name: "Water", Code: 121, Size: 5, Amount: 200},
// 		},
// 	}}
// 	args := [][]map[string]int{
// 		{{"goodCode": 123, "warehouseID": 1, "amount": 20}},
// 		{{"goodCode": 123, "warehouseID": 1, "amount": 120}},
// 		{{}},
// 		{{"goodCode": 12}},
// 		{{"goodCode": 12, "warehouseID": 1}},
// 	}
// 	expected := []map[string]string{
// 		{"error": "", "reply": fmt.Sprintf("Added goods %v", args[0])},
// 		{"error": "", "reply": "Added goods []"},
// 		{"error": "incorrect request, goodCode not presented", "reply": ""},
// 		{"error": "incorrect request, warehouseID not presented", "reply": ""},
// 		{"error": "incorrect request, amount not presented", "reply": ""},
// 	}

// 	if err := checkFunction(args, expected, mock.Add); err != nil {
// 		t.Error(err)
// 	}
// }

// func TestReserveGood(t *testing.T) {
// 	mock := &Goods{db: &MockDB{
// 		waregouseGoods: []WarehouseGood{
// 			{goodCode: 123, warehouseID: 1, availableAmount: 100},
// 			{goodCode: 121, warehouseID: 1, availableAmount: 200},
// 		},
// 	}}

// 	args := [][]map[string]int{
// 		{{"goodCode": 123, "warehouseID": 1, "amount": 20}},
// 		{{"goodCode": 123, "warehouseID": 1, "amount": 120}},
// 		{{}},
// 		{{"goodCode": 12}},
// 		{{"goodCode": 12, "warehouseID": 1}},
// 	}
// 	expected := []map[string]string{
// 		{"error": "", "reply": fmt.Sprintf("Reserved goods: %v", args[0])},
// 		{"error": "", "reply": "Reserved goods: []"},
// 		{"error": "incorrect request, goodCode not presented", "reply": ""},
// 		{"error": "incorrect request, warehouseID not presented", "reply": ""},
// 		{"error": "incorrect request, amount not presented", "reply": ""},
// 	}

// 	if err := checkFunction(args, expected, mock.Reserve); err != nil {
// 		t.Error(err)
// 	}
// }

// func TestCancelGoodReservation(t *testing.T) {
// 	mock := &Goods{db: &MockDB{
// 		waregouseGoods: []WarehouseGood{
// 			{goodCode: 123, warehouseID: 1, availableAmount: 80, reservedAmount: 20},
// 			{goodCode: 121, warehouseID: 1, availableAmount: 150, reservedAmount: 0},
// 		},
// 	}}

// 	args := [][]map[string]int{
// 		{{"goodCode": 123, "warehouseID": 1, "amount": 20}},
// 		{{"goodCode": 123, "warehouseID": 1, "amount": 10}},
// 		{{}},
// 		{{"goodCode": 12}},
// 		{{"goodCode": 12, "warehouseID": 1}},
// 	}
// 	expected := []map[string]string{
// 		{"error": "", "reply": fmt.Sprintf("Cancel reservations of goods: %v", args[0])},
// 		{"error": "there is not enough 123 goods at the 1 warehouse. Reserved only 0", "reply": "Cancel reservations of goods: []"},
// 		{"error": "incorrect request, goodCode not presented", "reply": ""},
// 		{"error": "incorrect request, warehouseID not presented", "reply": ""},
// 		{"error": "incorrect request, amount not presented", "reply": ""},
// 	}

// 	if err := checkFunction(args, expected, mock.CancelReservation); err != nil {
// 		t.Error(err)
// 	}
// }

// func checkFunction(args [][]WarehouseGoodAction, expected []map[string]string, f func([]WarehouseGoodAction, *Reply) error) error {
// 	reply := &Reply{}
// 	for i := 0; i < len(args); i++ {
// 		if err := f(args[i], reply); err != nil {
// 			if err.Error() != expected[i]["error"] {
// 				return fmt.Errorf("Expected error: %s, Got %s", expected[i]["error"], err)
// 			}
// 		}
// 		if reply.Data != expected[i]["reply"] {
// 			return fmt.Errorf("Expected: %s. Got: %s", expected[i]["reply"], reply.Data)
// 		}
// 		reply.Data = ""
// 	}

// 	return nil
// }
