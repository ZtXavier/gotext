package main

import "fmt"

type vertex struct{
	lat,long float64
}
//在这里定义一个map用字符串作为他的索引
var m map[string]vertex


func main(){
	m = make(map[string]vertex)
	m["bell labs"] = vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["bell labs"])
}