package main

import (
	"fmt"
	"strings"
)

func main(){
	//定义一个井字版
	board := [][]string{
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}
	board[0][0] = "x"
	board[2][2] = "0"
	board[1][2] = "x"
	board[1][0] = "0"
	for i := 0;i < len(board);i++{
		//这里将给每一个字符串后拼接一个" "
		fmt.Printf("%s \n",strings.Join(board[i]," "))
	}
}