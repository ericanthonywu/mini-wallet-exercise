package Utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"mini-wallet-exercise/Constant"
	"os"
)

func GenerateJwtToken(id string) (string, error) {
	secretToken := os.Getenv("JWTUSERSECRETTOKEN")

	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims[Constant.UserClaimsId] = id

	t, err := token.SignedString([]byte(secretToken))

	if err != nil {
		panic(err)
	}
	return t, nil
}

func GetUserIdJWTClaims(c echo.Context) uuid.UUID {
	userUUID, _ := uuid.Parse(fmt.Sprintf("%v", c.Get(Constant.UserClaimsId)))
	return userUUID
}
