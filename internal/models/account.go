package models

import "time"

type Account struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}