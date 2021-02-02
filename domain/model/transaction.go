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
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"-"`
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