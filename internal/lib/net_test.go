package lib_test

import (
	"nonelandBackendInterview/internal/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoGet_Binance_Time_Success(t *testing.T) {
	if testing.Short() {
		t.Skip("do not need run on short ver.")
	}
	resp, err := lib.DoGet("https://api.binance.com/api/v3/time", nil, nil)
	assert.Nil(t, err)
	assert.NotEmpty(t, resp)
}
