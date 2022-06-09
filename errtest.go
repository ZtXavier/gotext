package main

import (
	"fmt"
	"math"
)

// 注意: 在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环。
// 可以通过先转换 e 来避免这个问题：fmt.Sprint(float64(e))。这是为什么呢？
// 修改 Sqrt 函数，使其接受一个负数时，返回 ErrNegativeSqrt 值。 


type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Eorror() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v",float64(e))
}

func Sqrt(x float64) (float64, error) {
	return 0,nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}