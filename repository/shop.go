package repository

import (
	"github.com/algonacci/backend-evermos/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(dsn string) error {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto Migrate Table
	db.AutoMigrate(&model.User{}, &model.Shop{}, &model.Product{}, &model.Category{}, &model.Address{}, &model.Transaction{})

	return nil
}

// Fungsi-fungsi lainnya di sini
