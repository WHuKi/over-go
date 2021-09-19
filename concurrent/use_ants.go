package concurrent

/*
@brief 并发相关
ants 包的使用
*/

import (
	"fmt"

	"github.com/panjf2000/ants"
)

func Send() {
	p, _ := ants.NewPool(10, ants.WithPreAlloc(true))
	defer p.Release()

	for i := 0; i < 100; i++ {
		p1 := P{
			num: i,
		}
		_ = p.Submit(p1.Do)
	}

	//ants.Submit(Do)
}

type P struct {
	num int
}

func (p *P) Do() {
	fmt.Println("。。。。。", p.num)
}
