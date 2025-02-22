package utils

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
)

type (
	JWTConfig struct {
		SigningKey interface{}
	}
)

// Restricted middleware for authorized check
// with jwt token
func Restricted() echo.MiddlewareFunc {
	c := JWTConfig{}
	c.SigningKey = []byte(os.Getenv("JwtSecret"))
	return JWTWithConfig(c)
}

func JWTWithConfig(config JWTConfig) echo.MiddlewareFunc {
	if config.SigningKey == nil {
		slog.Warn("echo: jwt middleware requires signing key")

		panic("echo: jwt middleware requires signing key")
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// jwt disini biasanya menggunakan bearer dari token login
			// karna ini tidak ada login memakai singinkey saja
			return next(c)
		}
	}
}
