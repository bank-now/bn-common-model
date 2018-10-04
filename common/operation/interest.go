package operation

import (
	"encoding/json"
	"time"
)

const (
	InterestOperationV1Topic = "interest-operation-v1"
)

func NewInterestOperationV1(account string) *InterestOperationV1 {
	i := InterestOperationV1{
		Account: account,
		DateFor: time.Now()}
	return &i
}

type InterestOperationV1 struct {
	Account string    `json:"account"`
	DateFor time.Time `json:"dateFor"`
}

func (i *InterestOperationV1) ToJsonBytes() ([]byte, error) {
	return json.Marshal(i)
}

func GetInterestOperation(b []byte) (i *InterestOperationV1, err error) {
	err = json.Unmarshal(b, i)
	return

}
