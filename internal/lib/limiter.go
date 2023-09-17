package lib

import (
	"go.uber.org/ratelimit"
)

//	{
//		"rateLimitType": "REQUEST_WEIGHT",
//		"interval": "MINUTE",
//		"intervalNum": 1,
//		"limit": 1200
//	 },
const ratePerSecond int = 20

var BinanceRequestLimiter = ratelimit.New(ratePerSecond)
