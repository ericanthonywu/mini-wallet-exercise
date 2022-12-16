package Middleware

import (
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {

	defaultMiddleware(e)
	corsMiddleware(e)

	e.HTTPErrorHandler = httpHandler
	echo.NotFoundHandler = notFoundError
}
