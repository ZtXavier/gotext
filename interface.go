package main

import (
	"fmt"
	"math"
)

type Abser interface{
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := vertex{3,4}
	a = f
	a = &v
	// a= v
	fmt.Println(a.Abs())
}

type MyFloat float64
func (f MyFloat) Abs() float64 {
	if f < 0{
		return float64(-f)
	}
	return float64(f)
}
 type vertex struct {
	 x, y float64
 }

 func (v * vertex) Abs() float64 {
	 return math.Sqrt(v.x*v.x + v.y*v.y)
 }