package model

import (
	uuid "https://github.com/satori/go.uuid"
	"time"
)

const (
	TransactionPending string = "Pending"
	TransactionCompleted string = "Completed"
	TransactionError string = "Error"
	TransactionConfirmed string = "Confirmed"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []Transaction
}

type Transaction struct {
	Base                       `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	AccountFromID     string   `gorm:"column:account_from_id;type:uuid" valid:"notnull"`
	Amount            float64  `json:"amount" gorm:"type:float" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	PixkeyIdTo		  string   `gorm:"column:pix_key_id_to;type:uuid;not null" valid:"notnull"`
	Status            string   `json:"status" gorm:"type:varchar(20)" valid:"notnull"`
	Description       string   `json:"description" gorm:"type:varchar(255)" valid:"-"`
	CancelDescription string   `json:"cancel_description" gorm:"type:varchar(255)" valid:"-"`
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidadeStruct(t)

	if t.Amount <= 0 {
		return error.New("The amount must be greater than zero")
	}

	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError {
		return errors.New("Invalid status for transaction")
	}

	if t.PixKeyTo.AccountId == t.AccountFrom.ID {
		return error.New("Invalid Account")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewTransaction(accountfrom *Account, amount float64, pixkeyto *Pixkey, description string) (*Transaction, error) {

	transaction := Transaction{
		AccountFrom: accountfrom,
		Amount:      amount,
		PixKeyTo:    pixkeyto,
		Status:      TransactionPending,
		Description: description,
	}

	transaction.ID = uuid.newV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}

func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.UpdatedAt = time.Now()
	t.Description = description
	err := t.isValid()
	return err
}

func (t *Transaction) Confirmed() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()
	err := t.isValid()
	return err
}