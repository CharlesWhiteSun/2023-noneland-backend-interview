package lib

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// DoGet 使用 Get 呼叫 API
func DoGet(urlStr string, headers, params map[string]interface{}) ([]byte, error) {
	httpURL, _ := url.Parse(urlStr)
	queryInfo := httpURL.Query()
	for k, v := range params {
		queryInfo.Set(k, fmt.Sprintf("%v", v))
	}
	httpURL.RawQuery = queryInfo.Encode()
	request, _ := http.NewRequest(http.MethodGet, httpURL.String(), nil)

	for k, v := range headers {
		request.Header.Set(k, fmt.Sprintf("%v", v))
	}
	request.Header.Set("Content-Type", "application/json;charset=utf-8")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.New("DoGet error: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("DoGet StatusCode != 200, error: " + err.Error())
		}
		return nil, errors.New("DoGet StatusCode != 200, response.body: " + string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("DoGet io.ReadAll response.body error: " + err.Error())
	}
	return body, nil
}
