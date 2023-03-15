package main 
import "fmt"


// rune 就是在ASCII码的基础上多表示了很多字符
// 其大小为32bit,byte表示为8bit
func main() {
	s := "我爱GO"
	runes := []rune(s)
	runes[0] = '你'
	fmt.Println(string(runes))
	fmt.Println(len(runes))
}