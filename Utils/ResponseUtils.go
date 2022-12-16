package Utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"mini-wallet-exercise/Constant/APIResponse"
	"mini-wallet-exercise/Model"
	"net/http"
)

func InternalErrorResponse(err error, errorMessage string) error {
	fmt.Println("Internal error occured: " + err.Error())

	if errorMessage == "" {
		errorMessage = "unknown error occurred"
	}

	return echo.NewHTTPError(http.StatusInternalServerError, errorMessage)
}

func DatabaseErrorResponse(err error) error {
	return InternalErrorResponse(err, APIResponse.DatabaseErrorResponse)
}

func BadRequestResponse(data interface{}) error {
	return echo.NewHTTPError(http.StatusBadRequest, data)
}

func OkResponseMessage(c echo.Context, data interface{}, statusCode int) error {
	return c.JSON(statusCode, Model.NewSuccessResponse(data))
}
