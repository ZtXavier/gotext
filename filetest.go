package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main(){
	file,err := os.Open("words.txt")
	if err != nil {
		fmt.Printf("open error\n")
	}
	
	out,err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Read error\n")
	}
	lines := strings.Split(string(out),"\n")
	for _,l:= range lines { 
		fmt.Println(l)
	}
	fmt.Println(lines)
}
