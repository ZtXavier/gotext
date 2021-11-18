package main


//没有确定的初始值的变量会被赋予他们零值
import "fmt"

func main(){
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q",i,f,b,s)
}