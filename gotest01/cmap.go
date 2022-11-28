package main

import "fmt"

func main(){
	var couCapmap map [string]string /* 创建集合 */
	couCapmap = make(map[string]string)

	couCapmap ["france"]  = "巴黎"
	couCapmap ["j"] = "日本"
	
	for count := range couCapmap {
		fmt.Println(count,"cap is ",couCapmap[count])
	}
	// 查看元素在集合中是否存在
	capt,ok := couCapmap["english"]
	if(ok) {
		fmt.Println("cap is",capt)
	}else {
		fmt.Println("not exist")
	}
}