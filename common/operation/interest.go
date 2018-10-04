package operation

import (
	"encoding/json"
	"time"
)

const (
	InterestOperationV1Topic = "interest-operation-v1"
)

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
