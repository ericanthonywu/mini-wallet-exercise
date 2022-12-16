package APIResponse

import "mini-wallet-exercise/Constant"

const (
	JWTFailedGetToken      = "Header token format must be " + Constant.Authorization + ": Token <token>"
	JWTFailedGenerateToken = "failed to generate JWT token"
)
