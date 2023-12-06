package main

import (
	"os"
	"log"
	"bufio"
	"strconv"
	"fmt"
	"strings"
)

func skim_substrings(a string) {
	for _, r := range (1 << len(a)) {
		substring := ""
		for _, i := range len(a) {
			substring += string(i)
		}
		for mums, m := range numbers {
			if substring == mums {
				return substring
			}
		}
	}
	return nil
}

func main() {
	//returned solution value
	var solution int

	//array of numbers in their string form
	numbers := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven":7, "eight":8, "nine":9}

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

		for _, r := range current_line {
			if _, err := strconv.Atoi(r); err == nil {
    			new_line = new_line + r
				new_line2 = ""
				//fmt.Println("\n\nHere is new line")
				//fmt.Println(new_line)
			}else {
				new_line2 = new_line2 + r
				for n, num := range numbers {
					//for all substrings of new_line2, 
					if skim_substrings(new_line2) != nil
						new_line = new_line + numbers[skim_substrings(new_line2)]
						new_line2 = ""
						//fmt.Println("\nHere is new line after finding a text number")
						//fmt.Println(new_line)
					}
				}
			}
		}

		current_line = strings.Split(new_line, "")
		if len(new_line) > 2 {
			fmt.Println("Bigger than 2")
			current_line = append(current_line[:1], current_line[len(current_line)-1:]...)
			s, err:= strconv.Atoi((strings.Join(current_line, "")))
			if err != nil {}
			solution = solution + s
		}else if len(new_line) == 1 {
			fmt.Println("equal to 1")
			s, err:= strconv.Atoi((strings.Join(current_line, "")))
			if err != nil {}
			solution = solution + s*11
		}else {
			fmt.Println("Equal to 2")
			s, err:= strconv.Atoi((strings.Join(current_line, "")))
			if err != nil {}
			solution = solution + s
		}
	}
	fmt.Println(solution)
}