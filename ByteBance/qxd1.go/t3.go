package main 
import "fmt"

// go 中使用for作为唯一的循环
func main() {
	i := 1
	// 第一种死循环
	for {
		fmt.Printf("loop")
		break
	}
	for j := 7; j < 7; j++ {
		fmt.Printf("%d\n", j)
	}

	for n := 0; n < 5; n++ {
		if n %2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	for i <= 3 {
		fmt.Println(i)
		i++
	}


}