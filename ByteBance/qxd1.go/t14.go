package main 
import (
		"fmt"
		"strings"
)


func main() {
	a := "hello"
	fmt.Println(strings.Contains(a,"ll"))
	fmt.Println(strings.Count(a,"l"))
	// 求前缀
	fmt.Println(strings.HasPrefix(a,"he"))
	// 求后缀
	fmt.Println(strings.HasSuffix(a, "llo"))
	// 找到匹配到的字符串的首个下标
	fmt.Println(strings.Index(a,"ll"))
	// 在该字符串中间加入第二个输入的字符
	fmt.Println(strings.Join([]string{"he","llo"}, "-$"))
	fmt.Println(strings.Repeat(a, 2))
	fmt.Println(strings.Split("a-b-c", "-"))
	fmt.Println(strings.ToLower(a))
	fmt.Println(strings.ToUpper(a))
	fmt.Println(len(a))
	// 在go中一个汉字占用3个字节
	b := "你"
	fmt.Println(len(b))
}