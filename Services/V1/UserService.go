package V1

import (
	"github.com/google/uuid"
	"mini-wallet-exercise/Lib"
	"mini-wallet-exercise/Model/Database"
	"mini-wallet-exercise/Utils"
)

func CreateIfNotExists(customerXid string) uuid.UUID {
	userData := Database.User{}

	if err := Lib.DB.
		Where("id = ?", customerXid).
		Select("id").
		First(&userData).Error; err != nil {
		if Utils.IsDBNotFound(err) {
			err = Lib.DB.Create(&userData).Error
		}
		if err != nil {
			panic(err)
		}
	}

	return userData.ID
}

func GetWalletData(userUUID uuid.UUID) (error, Database.Wallet) {
	var wallet Database.Wallet

	if err := Lib.DB.Model(&wallet).
		Where("user_id = ?", userUUID).
		Select("id", "status", "balance", "enabled_at").
		First(&wallet).Error; err != nil {
		return Utils.DatabaseErrorResponse(err), Database.Wallet{}
	}

	return nil, wallet
}
