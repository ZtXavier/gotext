package main 
import "fmt"

type Person struct {
	Name string
} 

type Son struct {
	Sex string
}


func (p *Person) testAddress(son *Son) {
	fmt.Printf("testAddress: son的指针地址为: %p\n",son)
	fmt.Printf("testAddress: person的指针地址是: %p\n",p)
}



func main() {
	p := Person{Name: "jack"}
	s := Son{Sex: "boy"}
	p.testAddress(&s)
	fmt.Printf("main: son的指针地址为: %p\n",&s)
	fmt.Printf("main: person的指针地址是: %p\n",&p)
}