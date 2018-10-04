package operation

const (
	WriteOperationV1Topic = "write-operation-v1"
)

type WriteOperationV1 struct {
	Table  string `json:"table"`
	Method string `json:"method"`
	Item   []byte `json:"item"`
}
