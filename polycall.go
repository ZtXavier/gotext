// 类型断言

package main

import "fmt"

type Point struct {
	x int 
	y int
}

func main() {
	var a interface {}
	var p Point = Point{1,2}
	a = p // 这里是将a 指向 p(point)类型,所以下面的语句才会成功
	var b Point
	// b = a // 不可以直接赋值
	b = a.(Point) // 类型断言,类似类型强转
	fmt.Println(b)

	var b2 float64
	b2 = 1.1
	a = b2
	y ,ok := a.(float32)
	// if y, ok := a.(float32);ok {} 第二种判断条件,感觉很反人类 
	if ok == false {
		fmt.Println("default")
	}
	fmt.Printf("y 的类型是 %T %.2f", y, y)
}