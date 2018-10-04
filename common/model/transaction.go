package model

import (
	"encoding/json"
	"time"
)

const (
	TransactionTable = "transactions"
)

type Transaction struct {
	ID         string    `json:"id"`
	SystemCode string    `json:"sys_code"`
	Amount     float32   `json:"amt"`
	Timestamp  time.Time `json:"ts"`
	AccountID  string    `json:"account_id"`
}

func GetTransaction(b []byte) (*Transaction, error) {
	var item Transaction
	err := json.Unmarshal(b, &item)
	return &item, err

}
