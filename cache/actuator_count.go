package utils

/*
* @Author: hxk
* @Description: 延时执行任务或者立刻执行
* @File:
* *Data: 2021/09/02 11:00
 */

import (
	"time"
)

//OpType 操作类型
type OpType int

const (
	_            OpType = iota
	OpTypeNormal        // 普通设置
	OpTypeBegin         // 立刻开始
	OpTypeEnd           // 立马结束
)

const NanoSecond = 1 // 1纳秒

//重试配置
var retryConf = &ReTry{
	Max:   3,
	Delay: time.Second * 1,
}

type execFunc func() error

//CountDownReq 活动计时器
type CountDownReq struct {
	OpType   OpType // 操作类型
	OverTime time.Duration
	Key      string
}

func NewCountdownReq(overTime time.Duration, key string, opType OpType) *CountDownReq {
	return &CountDownReq{
		OverTime: overTime,
		Key:      key,
		OpType:   opType,
	}
}

//Exec 加载任务
func (c *CountDownReq) Exec(do execFunc) {
	// 2. 创建计时器
	ticker := time.NewTicker(c.OverTime)
	if err := GetCountDownActuator().SetTicker(c.Key, ticker); err != nil {
		return
	}

	// 3. 执行任务
	for _ = range ticker.C {
		if err := retryConf.Try(func() error {
			return do()
		}); err != nil {
		}
		GetCountDownActuator().RemoveTicker(c.Key)
		return
	}
}
