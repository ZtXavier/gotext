package main
import (
	"fmt"
	// "time"
)


func main() {
	ch := make(chan int , 3)
	go func() {
		ch <- sum(2,3)
		ch <- sum(1,6)
		ch <- sum(2,2)
	}()
	for x := range ch {
		fmt.Println(x)
	}
}

func sum(i , j int ) int {
	return i + j
}