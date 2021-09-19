package rate

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/ratelimit"
)

/*
golang 漏斗限流使用
*/

const execTimes = 1

var (
	limitersMap     *sync.Map
	once            sync.Once
	limiterCountMap *sync.Map
)

func init() {
	once.Do(func() {
		limitersMap = &sync.Map{}
		limiterCountMap = &sync.Map{}
	})
}

/*
MsgSendLimitRate
@brief 漏斗限流
@params key 该内容的键
@params rps 在该时间单位能执行的次数
@params period 时间单位 如 1s、1m、1h
*/
func MsgSendLimitRate(key string, rps int, period time.Duration) (limitExec bool) {
	// 1. 查看该时间窗口已经执行的次数
	v, ok := limiterCountMap.Load(key)
	if ok {
		if v.(int) > rps-execTimes {
			fmt.Println("超过次数")
			limitExec = true
			return
		}
	}

	// 2. 漏斗执行
	l, _ := limitersMap.LoadOrStore(key, ratelimit.New(rps, ratelimit.Per(period)))
	now := l.(ratelimit.Limiter).Take()
	count := execTimes
	if v != nil {
		count = v.(int) + execTimes
	}
	limiterCountMap.Store(key, count)
	fmt.Printf("key: %s, %v\n", key, now)

	// 3. 删除过期的漏斗
	go delLimiterMap(key, period)

	return
}

func delLimiterMap(key string, period time.Duration) {
	if period <= 0 {
		return
	}
	ticker := time.NewTicker(period)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			limitersMap.Delete(key)
			limiterCountMap.Delete(key)
			return
		}
	}
}
