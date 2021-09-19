package singleFight

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var (
	singleFight singleflight.Group
	once        sync.Once
)

func init() {
	once.Do(func() {
		singleFight = singleflight.Group{}
	})
}

/*
FightDemo
@brief 模拟大量请求获取缓存情况下使用singleFight
*/
func FightDemo() {
	// 模拟5个协程同时请求
	for i := 0; i < 5; i++ {
		go func(id int) {
			cacheV, err := getCache(strconv.Itoa(0))
			if err != nil {
				fmt.Println("err")
			}
			fmt.Println("结果：", cacheV)
		}(i)
	}
}

/*
@brief 获取缓存公共方法
*/
func getCache(key string) (string, error) {
	r := &cacheReq{cacheKey: key}
	v, err, _ := singleFight.Do(key, r.GetCache)
	if err != nil {
		return "", err
	}

	return v.(string), nil
}

/*
@brief 获取缓存
*/

type cacheReq struct {
	cacheKey string
}

func (r cacheReq) GetCache() (interface{}, error) {
	fmt.Println("请求：", r.cacheKey)
	time.Sleep(time.Second * 1)
	return r.cacheKey, nil
}
