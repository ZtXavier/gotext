package main

import "fmt"

type vertex struct{
	lat, long float64
}
//映射文法
// var m = map[string]vertex {
// 	"bell labs": vertex{
// 		40.68433,-74.9988,
// 	},
// 	"google": vertex{
// 		37.2133,0,
// 	},
// }

//简略写法

var m = map[string]vertex{
	"bell": {1313.22131,1232.44323,},
	"google":{666.666,777.777},
}



func main(){
	fmt.Println(m)
}