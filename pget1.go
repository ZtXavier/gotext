package main

import (
	"fmt"
	"math"
)

type vertex struct{
	x, y float64
}

func Abs(v vertex) float64{
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func scale(v *vertex, f float64){
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := vertex{3,4}
	scale(&v,10)
	fmt.Println(Abs(v))
}