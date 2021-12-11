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
	var crabs = make(map[int]int)

	// split stuff up and put into an array
	crabbies := strings.Split(string(content), ",")
	fmt.Println(len(crabbies))

	for _, mr_crabs := range crabbies {
		crab, _ := strconv.Atoi(mr_crabs)

		crabs[crab]++
	}
	sum := make(map[int]int)
	// for i := 0; i < 2000; i++ {
	// 	for k, v := range crabs {
	// 		if i > k {
	// 			sum[i] += (i - k) * v
	// 		} else if i < k {
	// 			sum[i] += (k - i) * v

	// 		}
	// 	}
	// }

	for i := 0; i < 2000; i++ {
		for k, v := range crabs {
			fuel := 0
			if i > k {
				for j := k; j < i; j++ {
					fuel += j - k + 1
				}
				sum[i] += v * fuel
			} else if i < k {
				for j := i; j < k; j++ {
					fuel += j - i + 1
				}
				sum[i] += v * fuel
			}
		}
	}
	min_value := 10000000000
	for k, v := range sum {
		if v < min_value {
			min_value = v
			fmt.Println(k, min_value)
		}

	}

}
