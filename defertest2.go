package main
import (
	"fmt"
	"sync"
	// "net/http"
)



func main() {
	var wg sync.WaitGroup
	var urls = []string {
		"http://www.baidu.com/",
		"https://www.qiqu.com/",
		"http://www.github.com/",
	}
	
	for _,url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			fmt.Println(url)
			urls[i] = urls[i] + "xiugai"
			fmt.Println(urls[i])
		}(url)
	}
	wg.Wait()
	fmt.Println("--------------")
	for _,url := range urls {
		fmt.Println(url)
	}
	fmt.Println("over")
}