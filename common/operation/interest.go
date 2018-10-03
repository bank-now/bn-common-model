package operation

import "time"

type InterestOperation struct {
	Account string    `json:"account"`
	DateFor time.Time `json:"dateFor"`
}
