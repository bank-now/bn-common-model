package model

import (
	"time"
)

type Transaction struct {
	ID         string    `json:"id"`
	SystemCode string    `json:"sys_code"`
	Amount     float32   `json:"amt"`
	Timestamp  time.Time `json:"ts"`
	AccountID  string    `json:"account_id"`
}
