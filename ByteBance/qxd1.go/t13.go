package main 
import (
		"errors"
		"fmt"
		// "strconv"
)

type user3 struct {
	name string
	passwd string 
}

func findUser(users []user3, name string) (*user3, error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {
	// u := make([]user3, 3)
	// i := 0
	// for i < 3 {
	// 	u[i].name = "s" + "w"
	// 	u[i].passwd = "123" + strconv.Itoa(i)
	// 	i++
	// }

	u, err := findUser([]user3{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name)
	if u, err := findUser([]user3{{"wang","1024"}}, "li"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(u.name)
	}
}