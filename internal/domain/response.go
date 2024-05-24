package domain

type Response struct {
	Error  []interface{}         `json:"error"`
	Result map[string]TickerInfo `json:"result"`
}

// TickerInfo represents the structure for each ticker information
type TickerInfo struct {
	A []string `json:"a"`
	B []string `json:"b"`
	C []string `json:"c"`
	V []string `json:"v"`
	P []string `json:"p"`
	T []int    `json:"t"`
	L []string `json:"l"`
	H []string `json:"h"`
	O string   `json:"o"`
}
