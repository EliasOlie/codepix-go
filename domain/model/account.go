package model

import (
	uuid "https://github.com/satori/go.uuid"
	"time"
)

type Account struct {
	Base      `valid:required`
	OwnerName string `json:"owner_name" gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	Bank      *Bank  `valid:"-"`
	BankID    string `gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Number    string `json:"number" gorm:"type:varchar(20)" valid:"notnull"`
	PixKeys [] *PixKey `gorm:"ForeignKey:AccountID" valid:"-"`

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
