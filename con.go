package main

//常量的声明与变量类似，只不过是使用 const 关键字
//常量可以是字符、字符串、布尔值或数值。常量不能用 := 语法声明。
import "fmt"

const Pi = 3.14

func main(){
	const world = "world"
	fmt.Println("hello",world)
	fmt.Println("happy",Pi,"Day")

	const Truth = true
	fmt.Println("Go rules?",Truth)
}