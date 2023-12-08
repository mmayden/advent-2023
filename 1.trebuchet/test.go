package main

import (
	"fmt"
)

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven":7, "eight":8, "nine":9}

	var str = "oseveno"
	var integers_found []string
	for i:=0; i<len(str); i++ {
		for l:=i+1; l<len(str); l++ { 
			for n := range numbers {
				if str[i:l] == n {
					integers_found = append(integers_found, str[i:l])
				}
			}
		}
	}
	fmt.Println(integers_found)
}