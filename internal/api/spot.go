package api

import (
	"net/http"
	"nonelandBackendInterview/internal/api/binance"
	"nonelandBackendInterview/internal/lib"

	"github.com/gin-gonic/gin"
)

// [GET] BinanceSpotExchangeInfoHandler 取得現貨交易規範
func BinanceSpotExchangeInfoHandler(c *gin.Context) {
	lib.BinanceRequestLimiter.Take()
	resp, err := binance.SpotExchangeInfo()
	if err != nil {
		errResponseWithStatus(c, http.StatusOK, ErrorCodeIntFailure, ErrMsgStr(err.Error()), nil)
		return
	}
	errResponseWithStatus(c, http.StatusOK, ErrorCodeIntSuccess, ErrMsgStrSuccess, resp)
}

// [GET] BinanceSpotBalanceHandler 取得現貨帳戶餘額
func BinanceSpotBalanceHandler(c *gin.Context) {
	lib.BinanceRequestLimiter.Take()
	resp, err := binance.SpotBalance()
	if err != nil {
		errResponseWithStatus(c, http.StatusOK, ErrorCodeIntFailure, ErrMsgStr(err.Error()), nil)
		return
	}
	errResponseWithStatus(c, http.StatusOK, ErrorCodeIntSuccess, ErrMsgStrSuccess, resp)
}

// [GET] BinanceSpotRecordsHandler 取得現貨帳戶轉入轉出紀錄
func BinanceSpotRecordsHandler(c *gin.Context) {
	lib.BinanceRequestLimiter.Take()
	resp, err := binance.SpotRecords()
	if err != nil {
		errResponseWithStatus(c, http.StatusOK, ErrorCodeIntFailure, ErrMsgStr(err.Error()), nil)
		return
	}
	errResponseWithStatus(c, http.StatusOK, ErrorCodeIntSuccess, ErrMsgStrSuccess, resp)
}
