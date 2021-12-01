package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	// Hello world, the web server
	fmt.Println("Hello AdventofCode 2021!")

	// read stuff
	content, err := ioutil.ReadFile("resources/data.txt")

	// mandatory error catching
	if err != nil {
		log.Fatal(err)
	}

	// split stuff up and put into an array
	lines := strings.Split(string(content), "\n")

	//convert strings to int, easier for part 2
	var int_lines [2001]int

	for i := 0; i < len(lines); i++ {
		int_lines[i], err = strconv.Atoi(lines[i])

	}
	part_2(int_lines)
}

func part_1(lines []string) {

	//some random variables to keep track of the amount of HIGHER/LOWER stuff
	higher_count := 0
	lower_count := 0

	for i := 0; i < len(lines); i++ {
		// nothing to see at the beginning
		if i == 0 {
			fmt.Println("WOOP")
			continue
		}
		if lines[i] > lines[i-1] {
			higher_count++
		}
		if lines[i] < lines[i-1] {
			lower_count++
		}
	}

	fmt.Println(higher_count)
	fmt.Println(lower_count)
}

func part_2(lines [2001]int) {
	higher_count := 0
	lower_count := 0

	for i := 0; i < len(lines); i++ {
		// nothing to see at the beginning
		if i == 0 || i == 1 || i == 2 {
			fmt.Println("BOOP")
			continue
		}
		if lines[i]+lines[i+1]+lines[i+2] > lines[i-1]+lines[i]+lines[i+1] {
			higher_count++
		}
		if lines[i]+lines[i+1]+lines[i+2] < lines[i-1]+lines[i]+lines[i+1] {
			lower_count++
		}
		fmt.Println(higher_count)
		fmt.Println(lower_count)
	}

}
