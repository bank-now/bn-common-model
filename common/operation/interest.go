package operation

import (
	"encoding/json"
	"github.com/bank-now/bn-common-io/zipkin"
	"github.com/google/uuid"
	"time"
)

const (
	InterestOperationV2Topic = "interest-operation-v2"
)

func NewInterestOperationV2(account string) *InterestOperationV2 {
	i := InterestOperationV2{
		Account: account,
		DateFor: time.Now()}

	u, _ := uuid.NewRandom()
	i.TraceId = u.String()
	return &i
}

type InterestOperationV2 struct {
	TraceId string       `json:"traceId"`
	Account string       `json:"account"`
	DateFor time.Time    `json:"dateFor"`
	Ghost   zipkin.Ghost `json:"ghost"`
}

func (i *InterestOperationV2) ToJsonBytes() ([]byte, error) {
	return json.Marshal(i)
}

func GetInterestOperation(b []byte) (*InterestOperationV2, error) {
	var item InterestOperationV2
	err := json.Unmarshal(b, &item)
	return &item, err

}
