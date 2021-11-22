package main

import "fmt"


func Sqrt(x float64) float64{
	var z float64
	z = 1
	for i := 1; i <= 10; i++ {
		z -= (z*z-x)/(2*z)
	}
	return z
}

func main(){
	fmt.Println(Sqrt(4))
}