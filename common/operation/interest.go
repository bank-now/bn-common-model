package operation

import (
	"encoding/json"
	"time"
)

type InterestOperation struct {
	Account string    `json:"account"`
	DateFor time.Time `json:"dateFor"`
}

func (i *InterestOperation) ToJsonBytes() ([]byte, error) {
	return json.Marshal(i)
}

func GetInterestOperation(b []byte) (i *InterestOperation, err error) {
	err = json.Unmarshal(b, i)
	return

}
