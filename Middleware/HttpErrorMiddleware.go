package Middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"mini-wallet-exercise/Model"
	"net/http"
)

func httpHandler(err error, c echo.Context) {
	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Logger().Error(report)

	var (
		code    = report.Code
		message = report.Message
	)

	switch report.Message {
	case "not found":
		code = http.StatusNotFound
		break
	case "400 Bad Request":
		code = http.StatusBadRequest
		message = "bad request"
	}

	if c.JSON(code, Model.NewErrorResponse(message)) != nil {
		fmt.Println(err)
	}
}

func notFoundError(c echo.Context) error {
	return c.JSON(http.StatusNotFound, Model.NewErrorResponse("not found"))
}
