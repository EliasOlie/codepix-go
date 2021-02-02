package model

import (
	uuid "https://github.com/satori/go.uuid"
	"time"
)

type Account struct {
	Base      `valid:required`
	OwnerName string `json:"owner_name" valid:"notnull"`
	Bank      *Bank  `valid:"-"`
	Number    string `json:"number" valid:"notnull"`
	PixKeys [] *PixKey `valid:"-"`

}

func (account *Account) isValid() error {
	_, err := govalidator.ValidadeStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(bank *Bank, number string, owner_name string) (*Account, error) {

	account := Account{
		OwnerName: ownername,
		Number:    number,
		Bank:      bank,
	}

	account.ID = uuid.newV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &account, nil
}
