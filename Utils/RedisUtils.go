package Utils

import (
	"mini-wallet-exercise/Constant"
	"mini-wallet-exercise/Lib"
	"time"
)

func generateUserWalletStatus(userId string) string {
	return Constant.UserWalletStatus + "-" + userId
}

func SetUserWalletStatus(userId string, status bool) {
	RedisExpireSec := GetEnvInt("REDIS_EXPIRE_SECONDS")
	Lib.RDBSet(generateUserWalletStatus(userId), status, time.Duration(RedisExpireSec)*time.Second)
}

func GetUserWalletStatus(userId string) (string, bool) {
	return Lib.RDBGet(generateUserWalletStatus(userId))
}

func DelUserWalletStatus(userId string) {
	Lib.RDBDel(generateUserWalletStatus(userId))
}
