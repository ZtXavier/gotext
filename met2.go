package main 

import (
	"fmt"
	"math"
)

type vertex struct{
	x, y float64
}

func(v vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func Absfunc(v vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
func main(){
	v := vertex{3,4}
	fmt.Println(v.Abs())
	fmt.Println(Absfunc(v))

	p := &vertex{4,3}
	fmt.Println(p.Abs())
	fmt.Println(Absfunc(*p))
}