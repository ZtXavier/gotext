package main

import (
	"fmt"

	"time"
)
// 总结如下,两者的区别是timer需要手动reset每次的超时时间,而ticker不需要,但是经过简单的测试发现ticker有可能不准确

func main() {
	myTimer := time.NewTimer(time.Second * 2) // 启动定时器
	var i int = 0
	for {
		select {
		case t := <-myTimer.C:
			i++
			fmt.Println("time is: ", t)
			myTimer.Reset(time.Second * 2) // 每次使用完后需要人为重置下
		}
	}


}
