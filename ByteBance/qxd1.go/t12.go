package main 

import "fmt"

type user2 struct {
	name string
	passwd string
}

func (u user2) checkpasswd(pd string) bool {
	return u.passwd == pd
}

func (u *user2) changepasswd(pd string) {
	u.passwd = pd
}

func main() {
	a := user2{name: "ass", passwd : "123"}
	if  ok := a.checkpasswd("122"); !ok {
		a.changepasswd("122")
	}
	// 这里的 %+v 和 %#v 都是限于结构体的
	fmt.Printf("%+v", a)
}