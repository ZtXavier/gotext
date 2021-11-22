package main

import (
	"fmt"
	"math"
)

func compute (fn func(float64 ,float64) float64) float64 {
	return fn(3, 4)
}

func main(){
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x+y*y)
	}
	//这里是应用传入的参数名来在compute内部调用函数并返回结果
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}