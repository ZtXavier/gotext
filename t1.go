package main

import(
	"fmt"
)

func main(){
	var stockcode=123
	var endate="2020-1-20"
	var url="code=%d&endate=%s"
	var target_url = fmt.Sprintf(url,stockcode,endate)
	fmt.Println(target_url)
}