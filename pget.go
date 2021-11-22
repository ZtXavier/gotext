package main 

import(
	"fmt"
	"math"
)

type vertex struct{
	x, y float64
}

func (v  vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x+v.y*v.y)
}

func (v *vertex) Scale(f float64){
	v.x = v.x * f
	v.y = v.y * f
}

func main(){
	v := vertex{3,4}
	v.Scale(10)
	fmt.Println(v.Abs())
}