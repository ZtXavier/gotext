package main 
import "fmt"

type user struct {
	name string 
	passwd string
}

func checkpasswd(u user, p string) bool {
	return u.passwd == p
}

func checkpasswd2(u *user, p string) bool {
	return u.passwd == p
}

func main() {
	a := user{name : "wang",passwd: "123"}
	b := user{"li", "456"}
	c := user{name : "zhang"}
	c.passwd = "333"
	var d user 
	d.name = "zhao"
	d.passwd = "54324"
	fmt.Println(a,b,c,d)
	fmt.Println(checkpasswd(a, "123"))
	fmt.Println(checkpasswd2(&a, "1234"))
}