package main 
import (
		"fmt"
		"strings"
)


func main() {
	var str string = "hi, I am zhengtian,hi"
	fmt.Printf("The position of zhengtian is %d\n",strings.Index(str, "zhengtian"))
	fmt.Printf("hi the first is %d\n",strings.Index(str, "hi"))
	fmt.Printf("hi last is %d\n",strings.LastIndex(str, "hi"))
	fmt.Printf("haha is %d\n",strings.Index(str,"haha"))
}