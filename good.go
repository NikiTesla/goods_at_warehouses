package goods_at_warehouses

type Good struct {
	Name   string  `json:"name"`
	Code   int     `json:"code"`
	Size   float64 `json:"size"`
	Amount int     `json:"amount"`
}
