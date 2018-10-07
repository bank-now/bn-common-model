package operation

import "github.com/google/uuid"

const (
	WriteOperationV1Topic = "write-operation-v1"
)

type WriteOperationV1 struct {
	TraceId string `json:"traceId"`
	Table   string `json:"table"`
	Method  string `json:"method"`
	Item    []byte `json:"item"`
}

func NewWriteOperationV1(table string, method string, b []byte) WriteOperationV1 {
	i := WriteOperationV1{
		Table:  table,
		Method: method,
		Item:   b}
	u, _ := uuid.NewRandom()
	i.TraceId = u.String()
	return i

}
