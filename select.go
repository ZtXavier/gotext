package main
import "fmt"
import "time"

func Chann(ch chan int, stp chan bool) {
	for j := 0; j < 10; j++ {
		ch <- j
		time.Sleep(time.Second)
	}
	stp <- true
}


func main() {
	ch := make(chan int)
	c := 0
	stp := make(chan bool)

	go Chann(ch, stp)

	// select 是作为随机触发其中收到chan并满足条件的事件
	for {
		select {
		case c = <- ch:
			fmt.Println("Receive C",c)
		case s := <- ch:
			fmt.Println("Receive S",s)
		case h := <-ch:
			fmt.Println("Receive H",h) 
		case _ = <- stp:
			goto end
		}
	}
end:
}