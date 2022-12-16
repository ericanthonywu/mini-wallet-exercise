package Route

import (
	"github.com/labstack/echo/v4"
	RouteV1 "mini-wallet-exercise/Route/V1"
)

func Init(e *echo.Echo) {
	api := e.Group("/api")
	RouteV1.Init(api)
}
