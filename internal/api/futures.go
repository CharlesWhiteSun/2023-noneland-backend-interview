package api

import (
	"net/http"
	"nonelandBackendInterview/internal/api/binance"
	"nonelandBackendInterview/internal/lib"

	"github.com/gin-gonic/gin"
)

// [GET] BinanceFuturesExchangeInfoHandler 取得合約交易規範
func BinanceFuturesExchangeInfoHandler(c *gin.Context) {
	lib.BinanceRequestLimiter.Take()
	resp, err := binance.FuturesExchangeInfo()
	if err != nil {
		errResponseWithStatus(c, http.StatusOK, ErrorCodeIntFailure, ErrMsgStr(err.Error()), nil)
		return
	}
	errResponseWithStatus(c, http.StatusOK, ErrorCodeIntSuccess, ErrMsgStrSuccess, resp)
}

// [GET] BinanceFuturesBalanceHandler 取得合約帳戶餘額
func BinanceFuturesBalanceHandler(c *gin.Context) {
	lib.BinanceRequestLimiter.Take()
	resp, err := binance.FuturesBalance()
	if err != nil {
		errResponseWithStatus(c, http.StatusOK, ErrorCodeIntFailure, ErrMsgStr(err.Error()), nil)
		return
	}
	errResponseWithStatus(c, http.StatusOK, ErrorCodeIntSuccess, ErrMsgStrSuccess, resp)
}
