package main

import "fmt"

func main(){
	countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	fmt.Println("origin map")
	
	for cou := range countryCapitalMap {
		fmt.Println(cou,"首都是",countryCapitalMap[cou])
	}
	delete(countryCapitalMap,"France")
	fmt.Println("France had delete")
	fmt.Println("删除后的map")
	for cou := range countryCapitalMap {
		fmt.Println(cou,"首都是",countryCapitalMap[cou])
	}	
}