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
	var lanternfishes = make(map[int]int)

	// split stuff up and put into an array
	fishies := strings.Split(string(content), ",")
	fmt.Println(len(fishies))

	for _, fish := range fishies {
		fish, err := strconv.ParseInt(fish, 10, 8)
		if err != nil {
			fmt.Println("woospie")
		}
		switch fish {
		case 0:
			lanternfishes[0]++
		case 1:
			lanternfishes[1]++
		case 2:
			lanternfishes[2]++
		case 3:
			lanternfishes[3]++
		case 4:
			lanternfishes[4]++
		case 5:
			lanternfishes[5]++
		case 6:
			lanternfishes[6]++
		case 7:
			lanternfishes[7]++
		case 8:
			lanternfishes[8]++
		}
	}

	fmt.Println(lanternfishes)
	for day := 0; day < 256; day++ {
		tmp := lanternfishes[0]
		lanternfishes[0] = lanternfishes[1]
		lanternfishes[1] = lanternfishes[2]
		lanternfishes[2] = lanternfishes[3]
		lanternfishes[3] = lanternfishes[4]
		lanternfishes[4] = lanternfishes[5]
		lanternfishes[5] = lanternfishes[6]
		lanternfishes[6] = lanternfishes[7]
		lanternfishes[7] = lanternfishes[8]
		lanternfishes[8] = tmp
		lanternfishes[6] += tmp
	}

	var count = 0
	for i := range lanternfishes {
		count += lanternfishes[i]

	}
	fmt.Println(count)
}
