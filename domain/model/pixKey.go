package model

import (
	uuid "https://github.com/satori/go.uuid"
	"time"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(account string) (*Account, error)
}

type PixKey struct {
	Base      `valid:"required`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountId string   `json:"account_id" valid:"notnull"`
	Account   *Account `valid:"notnull"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidadeStruct(pixKey)
	
	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("Invalid type of key")
	}

	if pixKey.Status != "active" && pixKey.Status != "innactive" {
		return errors.New("Invalid status")
	}
	
	if err != nil {
		return err
	}
	return nil
}

func NewPixKey(kind string, account *Account, key string) (*PixKey, error) {

	pixKey := PixKey{
		Kind:    kind,
		Account: account,
		Key:     key,
		Status:  "active",
	}

	pixKey.ID = uuid.newV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}

	return &pixKey, nil
}
