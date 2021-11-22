package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var [][]res uint8
    for i := 0;i < dy;i++{
		for j := 0;j < dx;j++{
			res[i][j] = (i+j)/2
		}
	} 
	return res
}

func main(){
	pic.Show(Pic)
}