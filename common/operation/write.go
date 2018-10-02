package operation

type Write struct {
	Table  string `json:"table"`
	Method string `json:"method"`
	Item   []byte `json:"item"`
}
