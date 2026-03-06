package models

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	AccountID uint      `json:"account_id"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"` // credit / debit
	Note      string    `json:"note"`
	CreatedAt time.Time
}