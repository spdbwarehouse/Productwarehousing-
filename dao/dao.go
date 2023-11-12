package dao

import (
	"time"
)

type User struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255;not null;unique" json:"name"`
	Description string    `gorm:"size:1024" json:"description"`
	StockCount  int       `gorm:"not null" json:"stock_count"`
	Inflow      int       `gorm:"not null" json:"inflow"`
	Outflow     int       `gorm:"not null" json:"outflow"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
