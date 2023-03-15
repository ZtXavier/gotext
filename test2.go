package main

import (
	// "strings"
	"fmt"
	// "bytes"
	
)


func test1(args ...string) { // 这里...是指可以接受任意个string参数
	for _ , v := range args {
		fmt.Println(v)
	}
}


func main() {
	var strss = []string {
		"qwr",
		"234",
		"yui",
		"cvbc",
	}
	// var strss1 = []string {
	// 	"111",
	// 	"222",
	// 	"333",
	// 	"444",
	// }
	a1 := "adw"
	strss = append(strss,a1) // 这里接受的是打散切片类型或者切片的元素类型
	fmt.Println(strss)
}