package main

import "fmt"
var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	fmt.Printf(a)
}

func m() {
	// a := "O"
	a = "O"
	fmt.Printf(a)
}