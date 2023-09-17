package lib_test

import (
	"fmt"
	"testing"
	"time"

	"go.uber.org/ratelimit"
)

func TestLimiter_Per_Second(t *testing.T) {
	if testing.Short() {
		t.Skip("do not need run on short ver.")
	}
	rl := ratelimit.New(2) // per second
	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		if i > 0 {
			fmt.Println(i, now.Sub(prev))
		}
		prev = now
	}
}
