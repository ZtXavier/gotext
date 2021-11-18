package main


//命名返回，go的返回值可以被命名，它们会被视为定义在函数顶部的变量
//尽量在简短的代码中使用这样的语句
import "fmt"

func split(sum int) (x, y int){
	x = sum * 4 /9
	y = sum - x
	return
}
func main(){
	fmt.Println(split(17))
}
