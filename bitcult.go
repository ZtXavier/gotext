package main 

// &^ 该作用是位清除,将指定位置上的值设置为0


import "fmt"

func main () {
	var x uint8 = 15
	var y uint8 = 12
	// 将x上对应的y的bit置为零
	fmt.Printf("%08b\n",x &^ y)

	// go 中^ 按位补足
	// 如果是无的符号数则用 ^n, 如果是有符号数则该数与-1^
	var n int8 = -1
	var m int8 = 2
	
	fmt.Printf("%08b\n",n)
	fmt.Printf("%08b\n",m)
	fmt.Printf("%08b\n",n ^ m)
}