package main
import (
	"fmt"
	"sync"
	"runtime"
)
// func calc(index string, a, b int) int {
//     ret := a + b
//     fmt.Println(index, a, b, ret)
//     return ret
// }

// // 将defer可以看作为入栈形式的数据结构
// // 参数和defer没关系，只是运行的时间变化了

// func main() {
//     a := 1
//     b := 2
//     defer calc("1", a, calc("10", a, b))
//     a = 0
//     defer calc("2", a, calc("20", a, b))
//     b = 1
// }

func main() {
	runtime.GOMAXPROCS(2)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 20; i++ {
			go func() {
					fmt.Println("A: ", i)
					wg.Done()
			}()
	}
	// for i := 0; i < 10; i++ {
	// 		go func(i int) {
	// 				fmt.Println("B: ", i)
	// 				wg.Done()
	// 		}(i)
	// }
	fmt.Println("start----------------------")
	wg.Wait()
	fmt.Println("end----------------------")
}