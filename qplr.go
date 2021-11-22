package main

import "fmt"

func main(){
	s := []int{2,3,5,7,11,13}
	printSlice(s)
    //每次改變上限時不會改變容量，但是當改變下限會改變大小
	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}
func printSlice(s []int){
	fmt.Printf("len = %d cap = %d %v\n",len(s),cap(s),s)
}