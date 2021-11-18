package main
//这里要注意的是不用某个函数包就不能在import中引用它
import (
	"fmt"
)

func add(x,y int) int {
	return x - y
}

func main(){
	fmt.Println(add(76,67))
}