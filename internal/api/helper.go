package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type okResponse struct {
	OK bool `json:"ok"`
}

func errResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, okResponse{
		OK: false,
	})
}

// ErrCodeInt 定義 API 返回前端的狀態代碼
type ErrCodeInt int

const (
	ErrorCodeIntUnauthorized ErrCodeInt = -1
	ErrorCodeIntFailure      ErrCodeInt = iota // 失敗
	ErrorCodeIntSuccess                        // 成功
	ErrorCodeIntPending                        // 等待
)

// ErrMsgStr 定義 API 返回前端的錯誤訊息字串
type ErrMsgStr string

const (
	ErrMsgStrFailure ErrMsgStr = "failure" // 失敗
	ErrMsgStrSuccess ErrMsgStr = "success" // 成功
	ErrMsgStrPending ErrMsgStr = "pending" // 等待
)

type okResponseWithStatus struct {
	ErrCode ErrCodeInt  `json:"errcode"` // API 返回前端的狀態代碼
	ErrMsg  ErrMsgStr   `json:"errmsg"`  // API 返回前端的錯誤訊息字串
	Data    interface{} `json:"data"`    // API 返回前端的內容
}

func errResponseFailure(c *gin.Context, httpStatus int, errCode ErrCodeInt, errMsg ErrMsgStr) {
	c.AbortWithStatusJSON(
		httpStatus,
		okResponseWithStatus{
			ErrCode: errCode,
			ErrMsg:  errMsg,
			Data:    nil,
		},
	)
}

func errResponseWithStatus(c *gin.Context, httpStatus int, errCode ErrCodeInt, errMsg ErrMsgStr, data interface{}) {
	c.JSON(
		httpStatus,
		okResponseWithStatus{
			ErrCode: errCode,
			ErrMsg:  errMsg,
			Data:    data,
		},
	)
}
