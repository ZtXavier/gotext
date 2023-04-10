package main 
func main() {
	var a int
	var b int32
	a = 15
	b =  a + a // 这里不能强转,两者的数据类型不同
	b = b + 5 // 因为5是常量,所以可以通过编译
}