package main

import "fmt"

type Person struct  {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)",p.Name,p.Age)
}
func main() {
	a := Person{"dawd",123}
	z := Person{"hdtfh",6534}
	fmt.Println(a, z)
}