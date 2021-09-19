package main

import (
	rate "over-go/rate_use"
	"time"
)

func main() {
	for i := 0; i < 4; i++ {
		rate.MsgSendLimitRate("helo", 1, time.Second)
	}

}
