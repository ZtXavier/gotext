package main 

import "fmt"

func main() {
	nums := []int{2,3,4}
	sum := 0
	// 第二种遍历数组的方式,会返还数组的 索引 和 值
	for i,num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:",num)

		}
	}
	fmt.Println(sum)


	// 注意映射表初始化的写法
	m := map[string]string{"a": "A", "b" : "B"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println("key:", k)
	}

}