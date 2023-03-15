package main

import (
	"strings"
	"fmt"
	// "bytes"
	
)
// 多种字符串拼接的方式
func main() {
	// a := [] string{"a","b","c"}

	// 方式1: +
	// ret := a[0] + a[1] + a[2];
	// 方式2: fmt.Sprintf
	// ret := fmt.Sprintf("%s%s%s",a[0],a[1],a[2]);
	// 方式3: strings.Builder
	// var sb strings.Builder

	// sb.WriteString(a[0])
	// sb.WriteString(a[1])
	// sb.WriteString(a[2])

	// ret := sb.String()

	// 方式4: bytes.Buffer

	// buf := new(bytes.Buffer)
	// buf.WriteString(a[0])
	// buf.WriteString(a[1])
	// buf.WriteString(a[2])
	
	// ret := buf.String()


	//方式5: strings.join 

	ret := strings.Join(a,"") // a 是一个string的切片

	fmt.Println(ret)
}


