package main
// import ("fmt")

func Generate(ch chan<- int) {
	for i := 2;;i++ {
		// print(i," =i ")
		ch <- i
	}
}

func Filter(in <-chan int, out chan <-int, prime int) {
	for {
		i := <-in
		// print("F:i=",i,"<")
		if i % prime != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < 10; i++ {
		print("start")
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		// num := <-ch1
		// print(num, "\n")
		ch = ch1
	}
}