package main 

import (
		"fmt"
		"time"
)

func main() {
	a := 2
	// go 的switch有两种写法
	// 第一种是switch后带变量的
	// 注意:在go的switch中到达某个case中不需要加break
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}



	t := time.Now()
	switch {
	case t.Hour() < 12 :
		fmt.Println("It`s before noon")
	default:
		fmt.Println("It`s after noon")
	}
}