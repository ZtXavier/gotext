package main 
import (
		"fmt"
		"time"
		
)

func main() {
	now := time.Now()
	fmt.Println(now)
	t := time.Date(2023, 5, 14, 10, 27, 0, 0, time.UTC)
	t1 := time.Date(2023, 5, 14, 10,29,0, 0, time.UTC)
	fmt.Println(t)
	fmt.Println(t.Year(),t.Month(),t.Day(),t.Hour(),t.Minute())
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	diff := t1.Sub(t)
	fmt.Println(diff)
	fmt.Println(diff.Minutes(),diff.Seconds())
	t3, err := time.Parse("2006-01-02 15:04:05", "2023-05-14 10:27:00")
	if err != nil {
		panic(err)
	}

	//所以
	// t := time.Date(2023, 5, 14, 10, 27, 0, 0, time.UTC)
	// t.Format("2006-01-02 15:04:05")
	
	// 相当于 t3, err := time.Parse("2006-01-02 15:04:05", "2023-05-14 10:27:00")
	fmt.Println(t3 == t) 
	
	// 获取时间戳
	fmt.Println(now.Unix())

}