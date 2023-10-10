package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Account, hesap modeli
type Account struct {
	gorm.Model
	FirstName string
	LastName  string
	Balance   float64
}

// CreateAccount, yeni bir hesap oluşturur
func CreateAccount(firstName string, lastName string, Balance float64) *Account {
	db, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		return nil
	}

	// Tabloları oluşturun.
	if err := db.AutoMigrate(&Account{}); err != nil {
		return nil
	}

	// Yeni bir hesap nesnesi oluşturun.
	account := Account{
		FirstName: firstName,
		LastName:  lastName,
		Balance:   Balance,
	}

	// Hesabı veritabanına kaydedin.
	if err := db.Create(&account).Error; err != nil {
		return nil
	}

	return &account
}
