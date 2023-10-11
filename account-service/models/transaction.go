package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Transaction struct {
	Type    string
	Balance float64
}

func TransactionMoney(id uint, transaction Transaction) (*Account, error) {
	db, err := gorm.Open(sqlite.Open("db1.sqlite"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var account Account
	if err := db.First(&account, id).Error; err != nil {
		return nil, err
	}

	if transaction.Type == "withdraw" && account.Balance >= transaction.Balance {
		account.Balance -= transaction.Balance
	} else if transaction.Type == "deposit" {
		account.Balance += transaction.Balance
	} else {
		return nil, nil
	}

	if err := db.Save(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
