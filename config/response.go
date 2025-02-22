package config

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ResponseBody struct {
	Status  int               `json:"status"`
	Total   int64             `json:"total,omitempty"`
	Data    interface{}       `json:"data,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
	Message string            `json:"message,omitempty"`
}

func Response(c echo.Context, data interface{}, err error) (e error) {
	code := 200
	response := new(ResponseBody)

	if err != nil {
		code = 400
	}

	if rb, ok := data.(*ResponseBody); ok {
		if rb == nil {
			rb = &ResponseBody{Status: code}
		}

		response = rb
	} else if code == 200 {
		response = &ResponseBody{Status: code, Data: data}
	}

	validError, ok := err.(validator.ValidationErrors)
	if ok {
		response = errorValidation(validError)
	}

	httpError, ok := err.(*echo.HTTPError)
	if ok {
		response = errorHTTP(httpError)
	}

	return c.JSON(code, response)
}

func errorValidation(errors validator.ValidationErrors) (mx *ResponseBody) {
	mx = &ResponseBody{Status: 400}
	result := make(map[string]string)

	for _, i := range errors {
		field := strings.ToLower(i.Field())
		tag := i.Tag()
		msg := fmt.Sprintf("%s is not valid", field)

		switch tag {
		case "required":
			msg = fmt.Sprintf("%s is required", field)
		}

		result[field] = msg
	}

	mx.Errors = result

	slog.Error("Body request error validation.")

	return
}

func errorHTTP(errors *echo.HTTPError) (mx *ResponseBody) {
	mx = &ResponseBody{Status: 400, Message: "Body request invalid."}

	slog.Warn("Body request invalid.")

	return
}
