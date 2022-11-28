package main

import (
	"fmt"
	"strconv"
)
// strconv 包中有基础的字符串的操作方法

func addupater() func (x int) int {
	var n int = 0
	var str string = "hello"
	return func (x int) int {
		 n += x
		//  str = str + strconv.FormatInt(int64(n),10) // string(x)将int转为x->ascii 而不是直接转为字符串
		 str = str + strconv.Itoa(n)
		 fmt.Println("str =",str)
		 return n
	} 
}


func main() {
	n := addupater()
	fmt.Println(n(1))
	fmt.Println(n(2))
	fmt.Println(n(3))
}