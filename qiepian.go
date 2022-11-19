package main
import "fmt"


func main() {
	var number = make([]int, 5, 5)
	for i := 0; i < cap(number); i++ {
		number[i] = i
	}
	fmt.Println(number)
	ret1 := number[0:3]
	fmt.Println(ret1)
}