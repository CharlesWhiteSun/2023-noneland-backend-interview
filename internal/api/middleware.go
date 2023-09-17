package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MiddlewareValidateHeaderToken 檢查 Token 邏輯
func MiddlewareValidateHeaderToken(c *gin.Context) {
	tokenString := c.GetHeader("token")
	if tokenString == "" {
		errResponseFailure(c, http.StatusUnauthorized, ErrorCodeIntUnauthorized, ErrMsgStr("header token is empty"))
		return
	}
	// TODO: validate user token data
}
