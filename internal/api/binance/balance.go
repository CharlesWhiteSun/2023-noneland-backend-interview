package binance

import (
	"encoding/json"
	"errors"
	"nonelandBackendInterview/internal/lib"
	"strconv"
	"time"
)

type Balance struct {
	Apikey      string
	Path        string
	Timestamp   string
	RecvWindows string
	Signature   string
}

type BalanceOptions struct {
	Apikey      string
	Path        string
	Timestamp   string
	RecvWindows string
	Signature   string
}

type IBalanceOption interface {
	apply(*BalanceOptions)
}

type FuncBalanceOption struct {
	f func(*BalanceOptions)
}

func (fo FuncBalanceOption) apply(option *BalanceOptions) {
	fo.f(option)
}

func BalanceWithApikey(apikey string) IBalanceOption {
	return FuncBalanceOption{
		f: func(option *BalanceOptions) {
			option.Apikey = apikey
		},
	}
}

func BalanceWithPath(path string) IBalanceOption {
	return FuncBalanceOption{
		f: func(option *BalanceOptions) {
			option.Path = path
		},
	}
}

func BalanceWithTimestamp(timestamp string) IBalanceOption {
	return FuncBalanceOption{
		f: func(option *BalanceOptions) {
			option.Timestamp = timestamp
		},
	}
}

func BalanceWithRecvWindows(recvWindows string) IBalanceOption {
	return FuncBalanceOption{
		f: func(option *BalanceOptions) {
			option.RecvWindows = recvWindows
		},
	}
}

func BalanceWithSignature(signature string) IBalanceOption {
	return FuncBalanceOption{
		f: func(option *BalanceOptions) {
			option.Signature = signature
		},
	}
}

// NewBalanceObj 新增 API Request 物件
// examples:
// - NewBalanceObj()
// - NewBalanceObj(BalanceWithApikey("8888"), BalanceWithPath("/api/v3/account"))
func NewBalanceObj(opts ...IBalanceOption) *Balance {
	init := BalanceOptions{
		Apikey:      defaultApiKey,
		Path:        defaultPath,
		Timestamp:   strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
		RecvWindows: defaultRecvWindows,
		Signature:   "",
	}

	for _, opt := range opts {
		opt.apply(&init)
	}

	return &Balance{
		Apikey:      init.Apikey,
		Path:        init.Path,
		Timestamp:   init.Timestamp,
		RecvWindows: init.RecvWindows,
		Signature:   init.Signature,
	}
}

var spotBalanceResp string = `
{
	"free": "10.12345"
}
`

type SpotBalanceResponse struct {
	Free string `json:"free"`
}

// GetSpotBalance 透過第三方 API 取得現貨帳戶餘額
func GetSpotBalance(param Balance) ([]byte, error) {
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
	return []byte(spotBalanceResp), nil
}

// SpotBalance 取得現貨帳戶餘額
func SpotBalance() (*SpotBalanceResponse, error) {
	obj := NewBalanceObj()
	data, err := GetSpotBalance(*obj)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("spot balance is empty")
	}

	var resp SpotBalanceResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.New("spot balance json.Unmarshal error:" + err.Error())
	}
	return &resp, nil
}

var futuresBalanceResp string = `
{
	"free": "10.12345"
}
`

type FuturesBalanceResponse struct {
	Free string `json:"free"`
}

// GetFuturesBalance 透過第三方 API 取得合約帳戶餘額
func GetFuturesBalance(param Balance) ([]byte, error) {
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
	return []byte(futuresBalanceResp), nil
}

// FuturesBalance 取得合約帳戶餘額
func FuturesBalance() (*FuturesBalanceResponse, error) {
	obj := NewBalanceObj()
	data, err := GetFuturesBalance(*obj)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("futures balance is empty")
	}

	var resp FuturesBalanceResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.New("futures balance json.Unmarshal error:" + err.Error())
	}
	return &resp, nil
}
