package main

import "fmt"

func swap(x,y string)(string,string){
	return y,x
}

func main(){
	//这里需要注意的是语法上的不同
	a,b := swap("hello","world")
	fmt.Println(a,b)
}