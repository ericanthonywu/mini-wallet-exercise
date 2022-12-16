package V1

import (
	"github.com/labstack/echo/v4"
	"mini-wallet-exercise/Lib"
	"mini-wallet-exercise/Model"
	"mini-wallet-exercise/Model/Database"
	"mini-wallet-exercise/Services/V1"
	"mini-wallet-exercise/Utils"
	"net/http"
	"strconv"
	"time"
)

func InitUser(c echo.Context) error {
	customerXid := c.FormValue("customer_xid")

	if customerXid == "" {
		return Utils.BadRequestResponse(map[string][1]string{
			"customer_xid": {"Missing data for required field."},
		})
	}

	userId := V1.CreateIfNotExists(customerXid)

	token, err := Utils.GenerateJwtToken(userId.String())

	if err != nil {
		return err
	}

	return Utils.OkResponseMessage(c, Model.InitializeAccountResponse{Token: token}, http.StatusCreated)
}

func EnableWallet(c echo.Context) error {
	userUUID := Utils.GetUserIdJWTClaims(c)

	var isEnable bool
	var needToCreate bool
	var err error = nil

	err = Lib.DB.Model(&Database.Wallet{}).
		Where("user_id = ?", userUUID).
		Select("status").
		First(&isEnable).Error

	if Utils.IsDBNotFound(err) {
		isEnable = false
		needToCreate = true
	}

	if err != nil && !Utils.IsDBNotFound(err) {
		return Utils.DatabaseErrorResponse(err)
	}

	if isEnable {
		return Utils.BadRequestResponse("Already enabled")
	}

	var wallet = Database.Wallet{
		Balance:   0,
		Status:    true,
		EnabledAt: time.Now(),
	}

	if needToCreate {
		wallet.UserId = userUUID
		err = Lib.DB.Create(&wallet).Error
	} else {
		err = Lib.DB.Model(&Database.Wallet{}).
			Where("user_id = ?", userUUID).
			Updates(&wallet).
			Error
	}

	if err != nil {
		return Utils.DatabaseErrorResponse(err)
	}

	return Utils.OkResponseMessage(c, &map[string]interface{}{
		"wallet": &Model.WalletResponse{
			ID:        wallet.ID,
			OwnedBy:   userUUID,
			Status:    Utils.ReturnWalletStatus(wallet.Status),
			EnabledAt: wallet.EnabledAt,
			Balance:   wallet.Balance,
		},
	}, http.StatusCreated)
}

func ViewWallet(c echo.Context) error {
	userUUID := Utils.GetUserIdJWTClaims(c)

	err, wallet := V1.GetWalletData(userUUID)
	if err != nil {
		return err
	}

	return Utils.OkResponseMessage(c, &map[string]interface{}{
		"wallet": &Model.WalletResponse{
			ID:        wallet.ID,
			OwnedBy:   userUUID,
			Status:    Utils.ReturnWalletStatus(wallet.Status),
			EnabledAt: wallet.EnabledAt,
			Balance:   wallet.Balance,
		},
	}, http.StatusCreated)
}

func DepositsWallet(c echo.Context) error {
	amount := c.FormValue("amount")
	amountUint64, _ := strconv.ParseUint(amount, 10, 64)
	referenceId := c.FormValue("reference_id")

	userUUID := Utils.GetUserIdJWTClaims(c)

	var err error = nil

	err = Lib.DB.Model(&Database.ReferenceLog{}).
		Where("reference_id = ?", referenceId).
		Select("1").Error

	if err != nil && !Utils.IsDBNotFound(err) {
		return Utils.DatabaseErrorResponse(err)
	}

	if Utils.IsDBNotFound(err) {
		return Utils.BadRequestResponse("reference_id has already used")
	}

	var walletModel Database.Wallet

	err = Lib.DB.Model(&walletModel).
		Where("user_id = ?", userUUID).
		UpdateColumn("balance", Lib.DB.Raw("balance + "+amount)).
		Error

	if err != nil {
		return Utils.DatabaseErrorResponse(err)
	}

	err, wallet := V1.GetWalletData(userUUID)
	if err != nil {
		return err
	}

	return Utils.OkResponseMessage(c, &map[string]interface{}{
		"deposit": &Model.DepositResponse{
			ID:          wallet.ID,
			DepositedBy: userUUID,
			Status:      "success",
			DepositedAt: time.Now(),
			Amount:      amountUint64,
			ReferenceId: referenceId,
		},
	}, http.StatusOK)
}

func WithdrawalsWallet(c echo.Context) error {
	amount := c.FormValue("amount")
	amountUint64, _ := strconv.ParseUint(amount, 10, 64)
	referenceId := c.FormValue("reference_id")

	userUUID := Utils.GetUserIdJWTClaims(c)

	var err error = nil

	err = Lib.DB.Model(&Database.ReferenceLog{}).
		Where("reference_id = ?", referenceId).
		Select("1").Error

	if err != nil && !Utils.IsDBNotFound(err) {
		return Utils.DatabaseErrorResponse(err)
	}

	if Utils.IsDBNotFound(err) {
		return Utils.BadRequestResponse("reference_id has already used")
	}

	var balance uint64

	err = Lib.DB.Model(&Database.Wallet{}).
		Select("balance").
		Where("user_id = ?", userUUID).
		First(balance).Error

	if balance > amountUint64 {
		return Utils.BadRequestResponse(map[string][1]string{
			"balance": {"balance is not enough to make withdrawals"},
		})
	}

	var walletModel Database.Wallet

	err = Lib.DB.Model(&walletModel).
		Where("user_id = ?", userUUID).
		UpdateColumn("balance", Lib.DB.Raw("balance - "+amount)).
		Error

	if err != nil {
		return Utils.DatabaseErrorResponse(err)
	}

	err, wallet := V1.GetWalletData(userUUID)
	if err != nil {
		return err
	}

	return Utils.OkResponseMessage(c, map[string]interface{}{
		"withdrawal": Model.WithdrawalResponse{
			ID:          wallet.ID,
			WithdrawnBy: userUUID,
			Status:      "success",
			WithdrawnAt: time.Now(),
			Amount:      amount,
			ReferenceId: referenceId,
		},
	}, http.StatusOK)
}

func DisableWallet(c echo.Context) error {
	userUUID := Utils.GetUserIdJWTClaims(c)

	var walletModel Database.Wallet
	err := Lib.DB.Model(&walletModel).Where("user_id = ?", userUUID).UpdateColumn("status", false).Error

	if err != nil {
		return Utils.DatabaseErrorResponse(err)
	}

	err, wallet := V1.GetWalletData(userUUID)
	if err != nil {
		return err
	}

	return Utils.OkResponseMessage(c, map[string]interface{}{
		"wallet": Model.WalletDisabledResponse{
			ID:         wallet.ID,
			OwnedBy:    userUUID,
			Status:     "disabled",
			DisabledAt: time.Now(),
			Balance:    wallet.Balance,
		},
	}, http.StatusOK)
}
