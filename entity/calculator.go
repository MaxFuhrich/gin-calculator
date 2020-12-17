package entity

type Calc struct {
	A             float64 `json:"a"`
	B             float64 `json:"b"`
	Result        float64 `json:"result"`
	OperationUsed string  `json:"operation_used"`
}
