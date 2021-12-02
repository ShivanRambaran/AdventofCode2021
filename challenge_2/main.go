package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	// horizontal := 0
	// vertical := 0

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

	fmt.Println(len(lines))
	var int_lines [1001]int

	for i := 0; i <= len(lines)-2; i++ {
		int_lines[i], err = strconv.Atoi(strings.Fields(lines[i])[1])
		lines[i] = strings.Fields(lines[i])[0]
	}

	fmt.Println(len(lines), len(int_lines))
	part_2(lines, int_lines)
}

func part_1(lines []string, int_lines [1001]int) {

	horizontal := 0
	vertical := 0

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i], int_lines[i])
		if lines[i] == "forward" {
			horizontal += int_lines[i]
		}
		if lines[i] == "up" {
			vertical += int_lines[i]
		}
		if lines[i] == "down" {
			vertical -= int_lines[i]
		}
	}
	fmt.Println(horizontal)
	fmt.Println(vertical)
	fmt.Println(horizontal * vertical)
}

func part_2(lines []string, int_lines [1001]int) {

	horizontal := 0
	vertical := 0
	aim := 0

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i], int_lines[i], aim)
		if lines[i] == "forward" {
			if aim == 0 {
				horizontal += int_lines[i]
			} else {
				horizontal += int_lines[i]
				vertical += aim * int_lines[i]
			}

		}
		if lines[i] == "up" {
			aim -= int_lines[i]
		}
		if lines[i] == "down" {
			aim += int_lines[i]
		}
	}
	fmt.Println(horizontal)
	fmt.Println(vertical)
	fmt.Println(horizontal * vertical)
}
