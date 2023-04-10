package main

import "fmt"

func main() {
	var a int16 = 34
	var b int32
	
	b = int32(a) // 通过这种显式转换方式来进行强转

	fmt.Println("32 int is: %d\n",b)
	fmt.Println("16 int is: %d\n",a)
}

// 这里需要 %g 用于格式化浮点型(%f输出浮点数.%e输出科学计数表示法)
// %0nd用于规定输出长度为n的整数

// %n.mg 表示数字n并精确到小数点后m位