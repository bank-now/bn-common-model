package operation

import (
	"github.com/bank-now/bn-common-io/zipkin"
	"github.com/google/uuid"
)

const (
	WriteOperationV2Topic = "write-operation-v2"
)

type WriteOperationV2 struct {
	TraceId string       `json:"traceId"`
	Table   string       `json:"table"`
	Method  string       `json:"method"`
	Item    []byte       `json:"item"`
	Ghost   zipkin.Ghost `json:"ghost"`
}

func NewWriteOperationV2(table string, method string, b []byte) WriteOperationV2 {
	i := WriteOperationV2{
		Table:  table,
		Method: method,
		Item:   b}
	u, _ := uuid.NewRandom()
	i.TraceId = u.String()
	return i

}
