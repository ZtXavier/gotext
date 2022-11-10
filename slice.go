package main 

import "fmt"


func main() {
	// 再记忆一次, := 与 = 的区别是前者直接声明变量并赋值
	numbers := []int{0,1,2,3,4,5,6,7,8,9};
	printSlice(numbers);
	
	number1 := make([]int,0,5);
	printSlice(number1);

}

func printSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x);
}