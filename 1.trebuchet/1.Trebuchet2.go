package main

import (
	"os"
	"log"
	"bufio"
	"strconv"
	"fmt"
	"strings"
)

//array of numbers in their string form
var numbers = map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven":7, "eight":8, "nine":9}

//Check all substrings for string number 1-9
func skim_substrings(a string) string {
	integers_found := ""
	for i:=0; i<len(a); i++ {
		for l:=i+1; l<=len(a); l++ { 
			for n := range numbers {
				if a[i:l] == n {
					integers_found += strconv.Itoa(numbers[a[i:l]])
					//fmt.Println("adding in skim")
				}
			}
		}
	}
	if integers_found == "" {
		return "nothin"
	}else {
		//fmt.Println("skimming found some stuff")
		//fmt.Println(integers_found)
		return integers_found
	}
}

func main() {
	//returned solution value
	var solution int

	//assign text file
	file, err := os.Open("trebuchet.txt")

	//log error if error
	if err != nil {
		log.Fatal(err)
	}

	//ensure file current_lineoses
	defer file.Close()

	//create scanner from the file to read text
	scanner := bufio.NewScanner(file)

	//assign values line by line
	for scanner.Scan() {

		//current_line
		current_line := strings.Split(scanner.Text(), "")
		new_line := ""
		new_line2 := ""

		//loop every character of line
		for _, r := range current_line {
			//if the character is an int, add it to newline
			if _, err := strconv.Atoi(r); err == nil {
				//fmt.Println("Current character is an int. NEXT!")
    			new_line += r
				new_line2 = ""
			}else {
				//If its not an int, add it to a new string new_line2
				new_line2 += r
				fmt.Println("Added another letter, not an int.")
				fmt.Println(new_line2)
				//if any substring of new_line2 spells a string integer 1-9, add to new_line
				if (skim_substrings(new_line2) != "nothin") {
					fmt.Println("integer found")
					fmt.Println("new line 1:  " + new_line + "\nnew_line2  " + new_line2)
					new_line += skim_substrings(new_line2)
					fmt.Println(new_line)
					fmt.Println("Is the above 0?")
					new_line2 = ""
				}
			}
		}
		fmt.Println("Whats coming")
		current_line = strings.Split(new_line, "")
		if len(new_line) > 2 {
			current_line = append(current_line[:1], current_line[len(current_line)-1:]...)
			fmt.Println(current_line)
			s, err:= strconv.Atoi((strings.Join(current_line, "")))
			if err != nil {}
			fmt.Println(s)
			solution = solution + s
		}else if len(new_line) == 1 {
			s, err:= strconv.Atoi((strings.Join(current_line, "")))
			if err != nil {}
			solution = solution + s*11
		}else {
			s, err:= strconv.Atoi((strings.Join(current_line, "")))
			if err != nil {}
			solution += s
		}
	}
	fmt.Println(solution)
}