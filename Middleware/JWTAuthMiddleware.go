package Middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"mini-wallet-exercise/Constant"
	"mini-wallet-exercise/Constant/APIResponse"
	"mini-wallet-exercise/Lib"
	"mini-wallet-exercise/Model/Database"
	"mini-wallet-exercise/Utils"
	"os"
	"strings"
)

func UserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return getToken(next)
}

func WalletMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var isDisabled bool
		userUUID := Utils.GetUserIdJWTClaims(c)
		err := Lib.DB.Model(&Database.Wallet{}).
			Where("user_id = ?", userUUID).
			Select("status").
			First(&isDisabled).
			Error

		if Utils.IsDBNotFound(err) {
			return Utils.BadRequestResponse("Disabled")
		}

		if err != nil {
			return Utils.DatabaseErrorResponse(err)
		}

		if !isDisabled {
			return Utils.BadRequestResponse("Disabled")
		}

		return next(c)
	}
}

func getToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		tokenHeader := c.Request().Header.Get(Constant.Authorization)

		if tokenHeader == "" {
			return Utils.BadRequestResponse(APIResponse.JWTFailedGetToken)
		}

		tokenString := strings.Split(tokenHeader, " ")

		if tokenString[0] != "Token" || tokenString[1] == "" {
			return Utils.BadRequestResponse(APIResponse.JWTFailedGetToken)
		}

		secretToken := os.Getenv("JWTUSERSECRETTOKEN")

		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				errorMessage := fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				return Utils.InternalErrorResponse(errorMessage, APIResponse.JWTFailedGenerateToken), errorMessage
			}

			return []byte(secretToken), nil
		})

		// check token generation error
		if err != nil {
			return Utils.InternalErrorResponse(err, APIResponse.JWTFailedGenerateToken)
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			return Utils.InternalErrorResponse(fmt.Errorf("failed to get jwt claims"), APIResponse.JWTFailedGenerateToken)
		}

		// set jwt claims
		c.Set(Constant.UserClaimsId, claims[Constant.UserClaimsId].(string))

		return next(c)
	}
}
