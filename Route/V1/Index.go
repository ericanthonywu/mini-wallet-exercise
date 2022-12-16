package Route

import (
	"github.com/labstack/echo/v4"
	"mini-wallet-exercise/Controller/V1"
	"mini-wallet-exercise/Middleware"
)

func Init(e *echo.Group) {
	api := e.Group("/v1")

	api.POST("/init", V1.InitUser)
	api.POST("/wallet", V1.EnableWallet, Middleware.UserMiddleware)
	api.GET("/wallet", V1.ViewWallet, Middleware.UserMiddleware, Middleware.WalletMiddleware)
	api.POST("/wallet/deposits", V1.DepositsWallet, Middleware.UserMiddleware, Middleware.WalletMiddleware)
	api.POST("/wallet/withdrawals", V1.WithdrawalsWallet, Middleware.UserMiddleware, Middleware.WalletMiddleware)
	api.PATCH("/wallet", V1.DisableWallet, Middleware.UserMiddleware, Middleware.WalletMiddleware)
}
