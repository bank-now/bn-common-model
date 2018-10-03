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
