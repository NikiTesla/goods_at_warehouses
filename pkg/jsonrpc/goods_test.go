package jsonrpc

import (
	"testing"

	lamodatest "github.com/NikiTesla/lamoda_test"
)

func TestGoodCreate(t *testing.T) {
	reply := &Reply{}
	mock := &Goods{db: &MockDB{}}

	args := []lamodatest.Good{
		{Name: "Coffee", Code: 123, Size: 1.1, Amount: 100},
		{Name: "Water", Code: 121, Size: 5, Amount: 200},
	}

	if err := mock.Create(args, reply); err != nil {
		t.Errorf("cannot create args %v, error: %s", args, err.Error())
	}
}
