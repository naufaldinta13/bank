package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/naufaldinta13/bank/config"
	"github.com/naufaldinta13/bank/utils"
)

func RegisterHandler(e *echo.Echo) {
	e.POST("/daftar", HandlerRegister, utils.Restricted())
	e.POST("/tabung", HandlerDeposit, utils.Restricted())
	e.POST("/tarik", HandlerWithdraw, utils.Restricted())
	e.GET("/saldo/:account_number", HandlerShow, utils.Restricted())
}

func HandlerRegister(c echo.Context) (e error) {
	var req registerRequest
	var res interface{}

	if e = c.Bind(&req); e == nil {
		if e = req.Validate(); e == nil {
			res, e = req.Execute()
		}
	}

	return config.Response(c, res, e)
}

func HandlerDeposit(c echo.Context) (e error) {
	var req depositRequest
	var res interface{}

	if e = c.Bind(&req); e == nil {
		if e = req.Validate(); e == nil {
			res, e = req.Execute()
		}
	}

	return config.Response(c, res, e)
}

func HandlerWithdraw(c echo.Context) (e error) {
	var req withdrawRequest
	var res interface{}

	if e = c.Bind(&req); e == nil {
		if e = req.Validate(); e == nil {
			res, e = req.Execute()
		}
	}

	return config.Response(c, res, e)
}

func HandlerShow(c echo.Context) (e error) {
	res := ShowCustomerByAccountNumber(c.Param("account_number"))

	return config.Response(c, res, nil)
}
