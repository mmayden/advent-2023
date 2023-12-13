package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//assign text file
	file, err := os.Open("cubes.txt")
	//log error if error
	if err != nil {
		panic(err)
	}
	//ensure file closes
	defer file.Close()

	//create scanner to scan file
	scanner := bufio.NewScanner(file)

	//keep track of which game
	game := 1
	//games that are impossible
	var games []int
	//possible games
	var possible_games []int
	//tallys for red/green/blue
	var numbers = map[string]int{"r": 0, "g": 0, "b": 0}
	//total to return
	total := 0

	//loop through scanner file line by line
	for scanner.Scan() {
		//current_line is string of entire current line being scanned
		current_line := strings.Split(scanner.Text(), "")
		//character number to avoid bound error
		char_num := -1

		//loop through every char in line
		for _, char := range current_line {
			char_num += 1
			//check
			check := ""
			fmt.Println("This is the char in play")
			fmt.Println(char)
			fmt.Println(numbers["r"])
			if char == ";" {
				if numbers["r"] > 12 || numbers["g"] > 13 || numbers["b"] > 14 {
					fmt.Println("Apparently this is a semicolon")
					games = append(games, game)
					fmt.Println("ffs")
				}
				numbers["b"] = 0
				numbers["g"] = 0
				numbers["r"] = 0
			}

			if char == "r" || char == "b" {
				if char_num > 2 {
					fmt.Println(current_line[char_num-1] + "wiuefhwiehfiuwehfiwe")
					if current_line[char_num-1] == "g" {
						check = strings.Fields(strings.Join(current_line[char_num-4:char_num-2]), "")
						fmt.Printf("Goofy %s", check)
						counter, err := strconv.Atoi(check)
						if err != nil {
						}
						fmt.Println("green counter")
						fmt.Println(check)
						fmt.Println(counter)
						//fmt.Println(check)
						fmt.Println(numbers[current_line[char_num-1]])
						numbers["g"] = numbers["g"] + counter
						fmt.Println(numbers[current_line[char_num-1]])
						fmt.Println("lawd")
					} else {
						check = strings.Join(current_line[char_num-3:char_num-1], "")
						fmt.Println("Not goofy, else")
						fmt.Println(check)
						counter, err := strconv.Atoi(check)
						if err != nil {
						}
						//fmt.Println(check)
						//strconv.Atoi(check)
						//fmt.Println(check)
						fmt.Println("r/b counter")
						fmt.Println(counter)
						numbers[char] += counter
						fmt.Println(numbers[char])
					}
				}
			}
		}
		game += 1
		numbers["r"] = 0
		numbers["g"] = 0
		numbers["b"] = 0
	}
	for i := 1; i <= 100; i++ {
		possible_games = append(possible_games, i)
	}
	for i, ig := range possible_games {
		for _, rg := range games {
			if ig == rg {
				possible_games = append(possible_games[0:i], possible_games[i+1:len(possible_games)]...)
			}
		}
	}
	for _, pg := range possible_games {
		total += pg
	}
	fmt.Println(possible_games)
	fmt.Println(games)
	fmt.Println(total)
}
