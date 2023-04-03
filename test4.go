package main
import "fmt"


func main() {
	// var sp map[int]int

	// //	 寻找
	// if _, ok := sp[10]; ok{

	// } else {

	// }

	// 匿名函数
		add := func(i,j int) int{
			return i + j;
		}

		fmt.Println(add(3,4));
}