package Utils

import (
	"errors"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func GetEnvInt(key string) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		panic(err)
	}

	return value
}

func IsDBNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func ReturnWalletStatus(status bool) string {
	if status {
		return "enabled"
	}
	return "disabled"
}
