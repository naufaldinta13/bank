package handler

import (
	"time"

	"github.com/naufaldinta13/bank/model"
	"github.com/naufaldinta13/bank/repository"

	"github.com/go-playground/validator/v10"
)

type depositRequest struct {
	Account_Number string  `json:"account_number" validate:"required,show"`
	Nominal        float64 `json:"nominal" validate:"gt=0"`

	Customer *model.Customer `json:"-"`
}

func (r *depositRequest) Validate() (e error) {
	v := validator.New()

	v.RegisterValidation("show", func(fl validator.FieldLevel) bool {
		if r.Account_Number != "" {
			if r.Customer, e = repository.NewCustomerRepository().FindByAccountNumber(r.Account_Number); e != nil {
				return false
			}
		}

		return true
	})

	return v.Struct(r)
}

func (r *depositRequest) Execute() (mx float64, e error) {
	m := &model.Transaction{
		CustomerID: r.Customer.ID,
		Type:       "deposit",
		Nominal:    r.Nominal,
		CreatedAt:  time.Now(),
	}

	if e = repository.NewTransactionRepository().Create(m); e == nil {
		repository.NewCustomerRepository().SyncSaldo(r.Customer.ID.String())

		mx = r.Customer.Saldo + r.Nominal
	}

	return
}
