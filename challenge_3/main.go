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

	part_2(lines)
}

func part_1(lines []string) {
	var zeroes [12]int
	var ones [12]int
	var oxygen string
	var co2 string

	// iterate through each "column" of the bytes, count the ones and zeroes
	for i := 0; i < 12; i++ {
		for _, line := range lines {
			// line = int as ASCII encoded integer
			if string(line[i]) == "0" {
				zeroes[i] += 1
			} else {
				ones[i] += 1
			}
		}
		fmt.Println(zeroes[i], ones[i])
		if zeroes[i] > ones[i] {
			oxygen += "0"
			co2 += "1"
		} else {
			oxygen += "1"
			co2 += "0"
		}
	}

	fmt.Println(oxygen, co2)
	//stuff for later
	oxygenvalue, err := strconv.ParseInt(oxygen, 2, 64)
	if err != nil {
		return
	}
	co2value, err := strconv.ParseInt(co2, 2, 64)

	fmt.Println(oxygenvalue * co2value)
}

func part_2(lines []string) {
	var zeroes [12]int
	var ones [12]int
	var oxygen string
	var co2 string

	// iterate through each "column" of the bytes, count the ones and zeroes
	for i := 0; i < 12; i++ {
		for bytepos, line := range lines {
			bytepos = bytepos

			// line = int as ASCII encoded integer
			if i == 0 {
				if string(line[i]) == "0" {
					zeroes[i] += 1
				} else {
					ones[i] += 1
				}
			} else {
				// too lazy to write something graceful, replace co2 with oxygen to check on oxygen order.
				if line[0:i] == co2[0:i] {
					fmt.Println(line, co2)
					if string(line[i]) == "0" {
						zeroes[i] += 1
					} else {
						ones[i] += 1
					}
				}
			}
		}
		fmt.Println(zeroes[i], ones[i])
		if zeroes[i] < ones[i] {
			oxygen += "1"
			co2 += "0"
		} else if zeroes[i] > ones[i] {
			oxygen += "0"
			co2 += "1"
		} else {
			oxygen += "1"
			co2 += "0"
		}
	}

	//o2 = 100111101011
	//co2 = 001011000101
}
