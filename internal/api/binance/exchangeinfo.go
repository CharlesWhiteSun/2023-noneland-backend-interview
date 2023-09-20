package binance

import (
	"encoding/json"
	"errors"
	"nonelandBackendInterview/internal/lib"
	"strconv"
	"time"
)

type ExchangeInfo struct {
	Apikey      string
	Path        string
	Timestamp   string
	RecvWindows string
	Signature   string
}

type ExchangeInfoOptions struct {
	Apikey      string
	Path        string
	Timestamp   string
	RecvWindows string
	Signature   string
}

type IExchangeInfoOption interface {
	apply(*ExchangeInfoOptions)
}

type FuncExchangeInfoOption struct {
	f func(*ExchangeInfoOptions)
}

func (fo FuncExchangeInfoOption) apply(option *ExchangeInfoOptions) {
	fo.f(option)
}

func ExchangeInfoWithApikey(apikey string) IExchangeInfoOption {
	return FuncExchangeInfoOption{
		f: func(option *ExchangeInfoOptions) {
			option.Apikey = apikey
		},
	}
}

func ExchangeInfoWithPath(path string) IExchangeInfoOption {
	return FuncExchangeInfoOption{
		func(option *ExchangeInfoOptions) {
			option.Path = path
		},
	}
}

func ExchangeInfoWithTimestamp(timestamp string) IExchangeInfoOption {
	return FuncExchangeInfoOption{
		func(option *ExchangeInfoOptions) {
			option.Timestamp = timestamp
		},
	}
}

func ExchangeInfoWithRecvWindows(recvWindows string) IExchangeInfoOption {
	return FuncExchangeInfoOption{
		func(option *ExchangeInfoOptions) {
			option.RecvWindows = recvWindows
		},
	}
}

func ExchangeInfoWithSignature(signature string) IExchangeInfoOption {
	return FuncExchangeInfoOption{
		func(option *ExchangeInfoOptions) {
			option.Signature = signature
		},
	}
}

// NewExchangeInfoObj 新增 API Request 物件
// examples:
// - NewExchangeInfoObj()
// - NewExchangeInfoObj(ExchangeInfoWithApikey("8888"), ExchangeInfoWithPath("/api/v3/account"))
func NewExchangeInfoObj(opts ...IExchangeInfoOption) *ExchangeInfo {
	init := ExchangeInfoOptions{
		Apikey:      defaultApiKey,
		Path:        defaultPath,
		Timestamp:   strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
		RecvWindows: defaultRecvWindows,
		Signature:   "",
	}

	for _, opt := range opts {
		opt.apply(&init)
	}

	return &ExchangeInfo{
		Apikey:      init.Apikey,
		Path:        init.Path,
		Timestamp:   init.Timestamp,
		RecvWindows: init.RecvWindows,
		Signature:   init.Signature,
	}
}

var spotExchangeInfoResp string = `
{
	"timezone": "UTC",
	"serverTime": 1565246363776,
	"rateLimits": [
		{
			"rateLimitType": "REQUEST_WEIGHT",
			"interval": "MINUTE",
			"intervalNum": 1,
			"limit": 1200
		},
		{
			"rateLimitType": "RAW_REQUESTS",
			"interval": "MINUTE",
			"intervalNum": 5,
			"limit": 6100
		}
	]
}
`

type SpotExchangeInfoResponse struct {
	Timezone   string          `json:"timezone"`
	ServerTime int64           `json:"serverTime"`
	RateLimits []SpotRateLimit `json:"rateLimits"`
}

type SpotRateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
}

// GetSpotExchangeInfo 透過第三方 API 取得現貨交易規範
func GetSpotExchangeInfo(param ExchangeInfo) ([]byte, error) {
	urlStr := binanceSpotBaseURL + param.Path
	params := map[string]interface{}{
		"timestamp":  param.Timestamp,
		"recvWindow": param.RecvWindows,
		"signature":  param.Signature,
	}
	headers := map[string]interface{}{
		"X-MBX-APIKEY": param.Apikey,
	}
	// TODO: get third party api data
	_, _ = lib.DoGet(urlStr, params, headers)
	return []byte(spotExchangeInfoResp), nil
}

// SpotExchangeInfo 取得現貨交易規範
func SpotExchangeInfo() (*SpotExchangeInfoResponse, error) {
	obj := NewExchangeInfoObj()
	data, err := GetSpotExchangeInfo(*obj)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("spot exchange info is empty")
	}

	var resp SpotExchangeInfoResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.New("spot exchange info json.Unmarshal error:" + err.Error())
	}
	return &resp, nil
}

var futuresExchangeInfoResp string = `
{
	"timezone": "UTC",
	"serverTime": 1565246363776,
	"rateLimits": [
	   {
		  "rateLimitType": "REQUEST_WEIGHT",
		  "interval": "MINUTE",
		  "intervalNum": 1,
		  "limit": 1200
	   },
	   {
		  "rateLimitType": "RAW_REQUESTS",
		  "interval": "MINUTE",
		  "intervalNum": 5,
		  "limit": 6100
	   }
	]
}
`

type FuturesExchangeInfoResponse struct {
	Timezone   string             `json:"timezone"`
	ServerTime int64              `json:"serverTime"`
	RateLimits []FuturesRateLimit `json:"rateLimits"`
}

type FuturesRateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
}

// GetFuturesExchangeInfo 透過第三方 API 取得合約交易規範
func GetFuturesExchangeInfo(param ExchangeInfo) ([]byte, error) {
	urlStr := binanceFuturesBaseURL + param.Path
	params := map[string]interface{}{
		"timestamp":  param.Timestamp,
		"recvWindow": param.RecvWindows,
		"signature":  param.Signature,
	}
	headers := map[string]interface{}{
		"X-MBX-APIKEY": param.Apikey,
	}
	// TODO: get third party api data
	_, _ = lib.DoGet(urlStr, params, headers)
	return []byte(futuresExchangeInfoResp), nil
}

// FuturesExchangeInfo 取得合約交易規範
func FuturesExchangeInfo() (*FuturesExchangeInfoResponse, error) {
	obj := NewExchangeInfoObj()
	data, err := GetFuturesExchangeInfo(*obj)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("spot exchange info is empty")
	}

	var resp FuturesExchangeInfoResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.New("spot exchange info json.Unmarshal error:" + err.Error())
	}
	return &resp, nil
}
