package repository

// type TransactionRepositoryInterface interface {
// 	Register(transaction *Transaction) error
// 	Save(transaction *Transaction) error
// 	Find(id string) (*Transaction, error)
// }

import (
	"fmt"
	"https://github.com/EliasOlie/codepix-go/tree/master/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (t TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	
	err := t.Db.Create(transaction).Error
	if err != nil {
		return err 
	}
	return nil
}
func (t TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	
	err := t.Db.Save(transaction).Error
	if err != nil {
		return err 
	}
	return nil
}
func (t, TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	t.Db.Preload("AccountFrom.Bank").Fisrt(&Transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("No transaction was found") 
	}

	return &transaction, nil
}
