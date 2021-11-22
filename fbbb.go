package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {
	sum := 0
	b := 1
	a := 0
	num := 0
	return func() int {
		// if num == 0{
		// 	num = num + 1
		// 	return 0
		// }
		// 	sum = a + b
		// 	a = b
		// 	b = sum		
		// num = num + 1
		// return a

		switch(num) {
		case 0:
			num = num + 1
			return 0
		default:
			sum = a + b
			a = b
			b = sum		
			num = num + 1
			return a
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
