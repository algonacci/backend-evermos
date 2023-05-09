package model

import "time"

type Shop struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255"`
	Address   string    `gorm:"size:255"`
	UserID    uint      `gorm:"uniqueIndex"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
