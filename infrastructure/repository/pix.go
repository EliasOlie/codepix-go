package repository

import (
	"fmt"
	"https://github.com/EliasOlie/codepix-go/tree/master/domain/model"
	"github.com/jinzhu/gorm"
)

// type PixKeyRepositoryInterface interface {
// 	RegisterKey(pixKey *PixKey) (*PixKey, error)
// 	FindKeyByKind(key string, kind string) (*PixKey, error)
// 	AddBank(bank *Bank) error
// 	AddAccount(account *Account) error
// 	FindAccount(account string) (*Account, error)
// }

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := r.Db.Create(bank).Error

	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := r.Db.Create(account).Error

	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := r.Db.Create(pixKey).Error

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

func (r PixKeyRepositoryDb) FindKeyById(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("No keys was found")
	}

	return &pixKey, nil
}

func (r PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	r.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("No account found")
	}

	return &account, nil
}

func (r PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.bank

	r.Db.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("No bank found")
	}

	return &bank, nil
}
