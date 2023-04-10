package main
/*
// C 标志io头文件，你也可以使用里面提供的函数
#include <stdio.h> 

void pri(){
    printf("hey");
}

int add(int a,int b){
    return a+b;
}
*/
import "C"  // 切勿换行再写这个

import "fmt"
/* const (
	_ = iota
	KB = 1 << (10 *  iota)
	MB
	// 使用_ 来忽略不需要的iota
) */

func main() {
	fmt.Println(C.add(2, 1))
	a := 1
	b := 2
	fmt.Println(a,b)
	a,b = b,a
	fmt.Println(a,b)
}

/* var a, b *int
var a int 
var b bool 
var str string 

var (
	a int 
	b int
	str string
) */

