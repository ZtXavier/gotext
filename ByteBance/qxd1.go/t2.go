package main 

import(
	"fmt"
	"math"
)


func main() {
	var a = "initial"
	
	var b, c int = 1, 2

	var d  = true

	var e  float64
	fmt.Printf("%f\n", e)

	f := float32(e)
	fmt.Printf("%f\n", e)
	g := a + "foo"
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g)
	const s string = "constant"
	const h = 500000000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}