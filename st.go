package main

import "fmt"

type vertex struct{
	x int
	y int
}

func main(){
	v := vertex{1,2}
	v.x = 5
	fmt.Println(v.x,v.y)
}