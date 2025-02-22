package handler

import (
	"github.com/naufaldinta13/bank/model"
	"github.com/naufaldinta13/bank/repository"

	"github.com/go-playground/validator/v10"
)

type registerRequest struct {
	NIK   string `json:"nik" validate:"required,len=16,unique_nik"`
	Name  string `json:"name" validate:"required"`
	Phone string `json:"phone" validate:"required,e164,unique_phone"`
}

func (r *registerRequest) Validate() (e error) {
	v := validator.New()

	repo := repository.NewCustomerRepository()

	v.RegisterValidation("unique_nik", func(fl validator.FieldLevel) bool {
		if r.NIK != "" {
			if cusnik, _ := repo.FindByNik(r.NIK); cusnik != nil {
				return false
			}
		}

		return true
	})

	v.RegisterValidation("unique_phone", func(fl validator.FieldLevel) bool {
		if r.Phone != "" {
			if cusphone, _ := repo.FindByPhone(r.Phone); cusphone != nil {
				return false
			}
		}

		return true
	})

	return v.Struct(r)
}

func (r *registerRequest) Execute() (an *model.Customer, e error) {
	repo := repository.NewCustomerRepository()

	mx := &model.Customer{
		NIK:           r.NIK,
		Name:          r.Name,
		PhoneNumber:   r.Phone,
		AccountNumber: repo.GenerateAccountNumber(),
	}

	if e = repo.Create(mx); e == nil {
		// an = mx.AccountNumber
		an = mx
	}

	return
}
