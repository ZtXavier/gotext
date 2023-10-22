package main
import(
	"fmt"
)

// 闭包具有记忆性

func Accumulate(value int) func() int{
	return func() int {
		value++;
		return value
	}
}


func main() {
	acc1 := Accumulate(1)
	fmt.Println(acc1())
	fmt.Println(acc1())

	fmt.Printf("%p\n",&acc1)

	acc2 := Accumulate(10)
	fmt.Println(acc2())
	fmt.Printf("%p\n",&acc2)
}