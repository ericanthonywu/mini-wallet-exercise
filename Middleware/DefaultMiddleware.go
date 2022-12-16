package Middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func defaultMiddleware(e *echo.Echo) {
	e.Use(
		middleware.Recover(),   //recover server on production if it's stop
		middleware.Logger(),    //logging
		middleware.RequestID(), //add request ID in every route
		middleware.Secure(),    //secure
	)
}
