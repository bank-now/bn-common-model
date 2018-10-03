package operation

import (
	"bytes"
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
	r := bytes.NewReader(b)
	decoder := json.NewDecoder(r)
	decoder.Decode(i)
	return
}
