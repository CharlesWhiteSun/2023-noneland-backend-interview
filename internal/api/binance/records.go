package binance

import (
	"encoding/json"
	"errors"
	"nonelandBackendInterview/internal/lib"
	"strconv"
	"time"
)

type Records struct {
	Apikey      string
	Path        string
	Timestamp   string
	RecvWindows string
	Signature   string
	StartTime   int64
	EndTime     int64
	Current     int64
	Size        int64
}

type RecordsOptions struct {
	Apikey      string
	Path        string
	Timestamp   string
	RecvWindows string
	Signature   string
	StartTime   int64
	EndTime     int64
	Current     int64
	Size        int64
}

type IRecordsOption interface {
	apply(*RecordsOptions)
}

type FuncRecordsOption struct {
	f func(*RecordsOptions)
}

func (fo FuncRecordsOption) apply(option *RecordsOptions) {
	fo.f(option)
}

func RecordsWithApikey(apikey string) IRecordsOption {
	return FuncRecordsOption{
		f: func(option *RecordsOptions) {
			option.Apikey = apikey
		},
	}
}

func RecordsWithPath(path string) IRecordsOption {
	return FuncRecordsOption{
		f: func(option *RecordsOptions) {
			option.Path = path
		},
	}
}

func RecordsWithTimestamp(timestamp string) IRecordsOption {
	return FuncRecordsOption{
		f: func(option *RecordsOptions) {
			option.Timestamp = timestamp
		},
	}
}

func RecordsWithRecvWindows(recvWindows string) IRecordsOption {
	return FuncRecordsOption{
		f: func(option *RecordsOptions) {
			option.RecvWindows = recvWindows
		},
	}
}

func RecordsWithSignature(signature string) IRecordsOption {
	return FuncRecordsOption{
		f: func(option *RecordsOptions) {
			option.Signature = signature
		},
	}
}

func RecordsWithStartTime(startTime int64) IRecordsOption {
	return FuncRecordsOption{
		f: func(option *RecordsOptions) {
			option.StartTime = startTime
		},
	}
}

func RecordsWithEndTime(endTime int64) IRecordsOption {
	return FuncRecordsOption{
		f: func(option *RecordsOptions) {
			option.EndTime = endTime
		},
	}
}

func RecordsWithCurrent(current int64) IRecordsOption {
	return FuncRecordsOption{
		f: func(option *RecordsOptions) {
			option.Current = current
		},
	}
}

func RecordsWithSize(size int64) IRecordsOption {
	return FuncRecordsOption{
		f: func(option *RecordsOptions) {
			option.Size = size
		},
	}
}

// NewRecordsObj 新增 API Request 物件
// examples:
// - NewRecordsObj()
// - NewRecordsObj(RecordsWithApikey("8888"), RecordsWithPath("/api/v3/account"))
func NewRecordsObj(opts ...IRecordsOption) *Records {
	init := RecordsOptions{
		Apikey:      defaultApiKey,
		Path:        defaultPath,
		Timestamp:   strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
		RecvWindows: defaultRecvWindows,
		Signature:   "",
		StartTime:   0,
		EndTime:     0,
		Current:     defaultCurrent,
		Size:        defaultSize,
	}

	for _, opt := range opts {
		opt.apply(&init)
	}

	return &Records{
		Apikey:      init.Apikey,
		Path:        init.Path,
		Timestamp:   init.Timestamp,
		RecvWindows: init.RecvWindows,
		Signature:   init.Signature,
		StartTime:   init.StartTime,
		EndTime:     init.EndTime,
		Current:     init.Current,
		Size:        init.Size,
	}
}

var spotRecordsResp string = `
{
	"rows": [
	   {
		  "amount": "0.10000000",
		  "asset": "BNB",
		  "status": "CONFIRMED",
		  "timestamp": 1566898617,
		  "txId": 5240372201,
		  "type": "IN"
	   },
	   {
		  "amount": "5.00000000",
		  "asset": "USDT",
		  "status": "CONFIRMED",
		  "timestamp": 1566888436,
		  "txId": 5239810406,
		  "type": "OUT"
	   },
	   {
		  "amount": "1.00000000",
		  "asset": "EOS",
		  "status": "CONFIRMED",
		  "timestamp": 1566888403,
		  "txId": 5239808703,
		  "type": "IN"
	   }
	],
	"total": 3
}
`

type SpotRecordsResponse struct {
	Rows  []SpotRecordRow `json:"rows"`
	Totol int             `json:"total"`
}

type SpotRecordRow struct {
	Amount    string `json:"amount"`
	Asset     string `json:"asset"`
	Status    string `json:"status"`
	Timestamp int64  `json:"timestamp"`
	TxId      int64  `json:"txId"`
	Type      string `json:"type"`
}

// GetSpotBalance 透過第三方 API 取得現貨帳戶轉入轉出紀錄
func GetSpotRecords(param Records) ([]byte, error) {
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
	return []byte(spotRecordsResp), nil
}

// SpotRecords 取得現貨帳戶轉入轉出紀錄
func SpotRecords() (*SpotRecordsResponse, error) {
	obj := NewRecordsObj()
	data, err := GetSpotRecords(*obj)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, errors.New("spot Records is empty")
	}

	var resp SpotRecordsResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.New("spot Records json.Unmarshal error:" + err.Error())
	}
	return &resp, nil
}
