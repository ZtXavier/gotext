package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second)
	defer timer.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-timer.C:
			fmt.Println("Current time: ", t)
			time.Sleep(4 * time.Second)
			timer.Reset(time.Second * 4)
		}
	}
}
