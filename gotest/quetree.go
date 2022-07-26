package main
import (
		"fmt"
		"./tree/btree"
)

func Walk(t *btree.Tree,ch chan int){
	visit(t,ch)
	close(ch)
}

// 写递归函数时需要放在另一个函数中(当操作共享资源的话)
func visit(t *btree.Tree,ch chan int) {
	ch <- t.Value
	if t.Left != nil{
		visit(t.Left, ch)
	}
	if t.Right != nil{
		visit(t.Right,ch)
	}
}

func Same(t1,t2 *btree.Tree) bool {
	i := 0
	ch1 := make(chan int,100)
	ch2 := make(chan int,100);
	go Walk(t1,ch1)
	go Walk(t2,ch2)
	for n := range ch1{
		fmt.Println(n)
		i ^= n
	}
	for n := range ch2{
		fmt.Println(n)
		i ^= n
	} 
	return i == 0
}

func main(){
	b1 := Same(btree.New(1),btree.New(1))
	b2 := Same(btree.New(2),btree.New(1));
	fmt.Println(b1,b2)
}