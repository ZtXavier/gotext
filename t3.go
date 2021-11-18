package main

//这里的按格式输出

import (
	"fmt"
	"math/cmplx"
)

var(
	ToBe bool = true
	MaxInt uint64 = 1<<64-1
	z    complex128 = cmplx.Sqrt(-5 + 12i)
)

func main(){
	fmt.Printf("Type: %T value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T value %v\n", z, z)
}