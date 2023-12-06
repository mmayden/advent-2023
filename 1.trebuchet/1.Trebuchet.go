//Dec 1 2023, Advent Calendar Day 1

package main

import (
	"os"
	"log"
	"bufio"
	"strconv"
	"fmt"
	"strings"
	//"reflect"
)

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
		//new_line is used to keep track of integers and produce final integer
		new_line := ""
		//for every slice in current_line
		for _, r := range current_line {
			//if attempting to convert slice to integer doesn't error add it to new_line
			if _, err := strconv.Atoi(r); err == nil {
    			new_line = new_line + r
			}
		}
		//reuse current_line by equaling it to an array of slices for each new_line value
		current_line = strings.Split(new_line, "")
		//check current_line 
		if len(new_line) > 2 {
			current_line = append(current_line[:1], current_line[len(current_line)-1:]...)
			s, err:= strconv.Atoi((strings.Join(current_line, "")))
			if err != nil {}
			fmt.Println(solution)
			solution = solution + s
			fmt.Println("into")
			fmt.Println(solution)

		}else if len(new_line) == 1 {
			s, err:= strconv.Atoi((strings.Join(current_line, "")))
			if err != nil {}
			solution = solution + s*11
		}else {
			fmt.Println(solution)
			fmt.Println("else")
			s, err:= strconv.Atoi((strings.Join(current_line, "")))
			if err != nil {}
			solution = solution + s
			fmt.Println(solution)
		}
	}
	//log error if scanner error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(solution)
}