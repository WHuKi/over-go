/*
context 使用场景1 -> 超时控制
*/
package concurrent

import (
	"context"
	"fmt"
	"time"
)

// 模拟一个耗时的操作
func rpc() (string, error) {
	time.Sleep(2 * time.Second)
	return "rpc done", nil
}

type result struct {
	data string
	err  error
}

func Handle(ctx context.Context, ms int) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	r := make(chan result)
	go func() {
		data, err := rpc()
		r <- result{
			data: data,
			err:  err,
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("timeout: %d ms, context exit: %+v\n", ms, ctx.Err())
	case res := <-r:
		fmt.Printf("result: %s, err: %+v\n", res.data, res.err)
	}

}
