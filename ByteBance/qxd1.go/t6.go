package main 
import "fmt"


func main() {
	// 创建一个大小为3的切片
	s := make([]string, 3)
	// go  中没有像c++的字面值常量
	// for i := 0;i < 3; i++ {
	// 	s[i] = "a" +  i
	// }
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	// 切片自动扩容
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	// go 不支持py的负数索引,且区间范围是左闭右开
	fmt.Println(s[2 : 5])
	fmt.Println(s[ : 5])
	fmt.Println(s[2 :])
	good := []string {"g", "o", "o", "d"}
	fmt.Println(good)
}