package operation

type Write struct {
	Table  string `json:"table"`
	Method string `json:"Method"`
	Item   []byte `json:"Item"`
}
