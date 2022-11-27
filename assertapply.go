package main

import "fmt"

type Student struct {

}


func typetest(item ...interface{}) {
	for i, v := range item {
		switch v.(type) {
		case int, int32, int64:
			fmt.Printf("第%d 个数据类型是 %T value is %v\n", i, v, v)
		case float64:
			fmt.Printf("第%d 个数据类型是 %T value is %v\n", i, v, v)
		case float32:
			fmt.Printf("第%d 个数据类型是 %T value is %v\n", i, v, v)
		case string:
			fmt.Printf("第%d 个数据类型是 %T value is %v\n", i, v, v)
		case bool:
			fmt.Printf("第%d 个数据类型是 %T value is %v\n", i, v, v)
		case Student:
			fmt.Printf("第%d 个数据类型是 %T value is %v\n", i, v, v)
		case *Student:
			fmt.Printf("第%d 个数据类型是 %T value is %v\n", i, v, v)
		default:
			fmt.Printf("there is no the type\n")
		}
	}
}


func main() {
	a := int64(1)
	b := int32(2)
	c := int(3)
	d := string("dawdaw")
	e := Student{}
	e1 := &Student{}
	f := float64(1.1)
	g := float32(2.2)
	typetest(a,b,c,d,e,f,g,e1)	
}