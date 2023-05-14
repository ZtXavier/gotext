package main 

import (
	"encoding/json"
	"fmt"
)


// 一般结构体中的变量都首字母大写,否则会引起意想不到的bug
type userInfo struct {
	Name string 	`json:"nname"` // 这里是为了在json中改变变量的名
	Age int 	
	Hobby []string
}

func main() {
	a := userInfo{Name : "wang", Age : 18, Hobby: []string{"Golang","c++"}}
	buf, err := json.Marshal(a) // json的序列化1,返回byte类型,用string()来翻译

	if err != nil {
		panic(err)
	}
	fmt.Println(buf)
	fmt.Println(string(buf))

	buf2, err := json.MarshalIndent(a, "", "\t") //序列化2,参数2,3都是调节json的格式
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf2))
	var b userInfo 
	err = json.Unmarshal(buf, &b)//反序列化
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", b)
}