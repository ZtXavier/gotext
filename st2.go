package main

import "fmt"

type vertex struct{
	x,y int
}

var (
	v1 = vertex{1,2}
	v2 = vertex{x: 1}	//x = 0;y = 0
	v3 = vertex{}        //x,y = 0
	p = &vertex{1,2}
)
func main(){
	fmt.Println(v1,p,v2,v3)
}