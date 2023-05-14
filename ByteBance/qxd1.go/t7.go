package main 
import "fmt"
// 这里是用make初始化了一个map
// go中的map索引是随即的,所以一般不会遍历map的key
func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(m["one"])
	fmt.Println(m["unknow"])

	r, ok := m["unknow"]
	fmt.Println(r, ok)
	delete(m, "one")

	m2 := map[string]int {"one" : 1,"two" : 2}
	var m3 = map[string]int {"one" : 1, "two" : 2}
	fmt.Println(m2, m3)
}