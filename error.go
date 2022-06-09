package main

import (
	"fmt"
	"time"
)

type Myerror struct {
	When time.Time
	What string
}

func (e *Myerror) Error() string {
	return fmt.Sprintf("at %v, %s",e.When,e.What)
}
func run() error {
	return &Myerror{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run();err != nil {
		fmt.Println(err)
	}
}