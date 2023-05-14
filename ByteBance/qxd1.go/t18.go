package main 
import (
		"fmt"
		"os"
		"os/exec"
)


func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Getenv("PATH"))
	fmt.Println(os.Setenv("AA","BB"))
	fmt.Println(os.Getenv("AA"))
	buf, err := exec.Command("ls", "./").CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}