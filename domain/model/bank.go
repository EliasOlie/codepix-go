package model

import (
	uuid "https://github.com/satori/go.uuid"
	"time"
)

type Bank struct {
	Base `valid:required`
	Code      string 	`json:"code" gorm:"type:varchar(20)" valid:"notnull"`
	Name      string 	`json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Accounts [] *Account `gorm:"ForeignKey:BankID" valid:"-"` 
	
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidadeStruct(bank)
	if err 	!= nil {
		return err
	}
	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank {
		Code: code,
		Name: name,
	}
	bank.ID = uuid.newV4().String()
	bank.CreatedAt = time.Now()

	err := bank.isValid()
	if err != nil {
		return nil, err
	}

	return &bank, nil
}
