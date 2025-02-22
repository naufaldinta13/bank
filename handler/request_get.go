package handler

import (
	"github.com/naufaldinta13/bank/config"
	"github.com/naufaldinta13/bank/repository"
)

func ShowCustomerByAccountNumber(an string) (res *config.ResponseBody) {
	res = &config.ResponseBody{Status: 200}
	if mx, e := repository.NewCustomerRepository().FindByAccountNumber(an); e != nil {
		res.Status = 400
		result := make(map[string]string)
		result["remarks"] = "account number is not valid."

		res.Errors = result
	} else {
		res.Data = mx
	}

	return
}
