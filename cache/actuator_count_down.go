package utils

import (
	"sync"
	"time"

	"github.com/pkg/errors"
)

/*
* @Author: hxk
* @Description: 任务计时器工具，目前使用到促销和秒杀，以便当活动开始的时候将活动的状态改编为已开始或者已结束
* @File:
* Date: 2021/09/02 11:00
 */

var (
	countDownOnce     sync.Once
	countDownActuator *CountDownActuator
)

//CountDownActuator 倒时计数器
type CountDownActuator struct {
	TickerMap *sync.Map
	Key       []string
}

//GetCountDownActuator 获取计数器对象
func GetCountDownActuator() *CountDownActuator {
	countDownOnce.Do(func() {
		countDownActuator = new(CountDownActuator)
		countDownActuator.TickerMap = &sync.Map{}
	})

	return countDownActuator
}

//SetTicker 设置任务
func (c *CountDownActuator) SetTicker(key string, ticker *time.Ticker) error {
	if key == "" {
		return errors.Wrap(nil, "SetTicker failed, key is empty")
	}

	// 1.先移除相同的任务
	c.RemoveTicker(key)
	// 2.添加任务
	c.TickerMap.Store(key, ticker)
	c.Key = append(c.Key, key)

	return nil
}

//RemoveTicker 移除任务
func (c *CountDownActuator) RemoveTicker(key string) {
	ticker, ok := c.GetTicker(key)
	if !ok {
		return
	}

	ticker.Stop()
	c.TickerMap.Delete(key)
	for k, v := range c.Key {
		if v == key {
			c.Key = append(c.Key[:k], c.Key[k:]...)
			break
		}
	}

}

//GetTicker 获取任务
func (c *CountDownActuator) GetTicker(key string) (ticker *time.Ticker, ok bool) {
	v, ok := c.TickerMap.Load(key)
	if ok {
		ticker = v.(*time.Ticker)
		return ticker, ok
	}

	return
}

//Close 关闭所有的任务
func (c *CountDownActuator) Close() {
	for _, v := range c.Key {
		t, ok := c.GetTicker(v)
		if ok {
			t.Stop()
		}
	}
}
