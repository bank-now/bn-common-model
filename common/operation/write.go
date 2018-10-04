package operation

import "encoding/json"

const (
	WriteOperationV1Topic = "write-operation-v1"
)

type WriteOperationV1 struct {
	Table  string `json:"table"`
	Method string `json:"method"`
	Item   []byte `json:"item"`
}

func (i *WriteOperationV1) ToJsonBytes() ([]byte, error) {
	return json.Marshal(i)
}
