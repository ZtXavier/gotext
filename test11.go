package main 
import (
		"fmt"
		"strings"
)

// 一般strings是用来处理字符串

// strings.HasPrefix(str,dis) 判断字符串str的开头是否是以dis开始的

func main() {
	var str string = "this is an example of a string"
	fmt.Printf("T/F ? the string begin with %t", strings.HasPrefix(str, "th"))
}