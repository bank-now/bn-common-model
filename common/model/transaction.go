package model

import "github.com/golang/protobuf/ptypes/timestamp"

type Transaction struct {
	ID         string              `json:"id"`
	SystemCode string              `json:"sys_code"`
	Amount     float32             `json:"amt"`
	Timestamp  timestamp.Timestamp `json:"ts"`
	AccountID  string              `json:"account_id"`
}
