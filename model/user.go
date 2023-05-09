package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:255"`
	Email     string    `gorm:"uniqueIndex"`
	Password  string    `gorm:"size:255"`
	Phone     string    `gorm:"uniqueIndex"`
	Role      string    `gorm:"size:50"`
	ShopID    uint      `gorm:"index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
