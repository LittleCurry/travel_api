package misc

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func KeyAuth() echo.MiddlewareFunc {
	return middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
		Validator: Validator,
	})
}
