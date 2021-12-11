package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type vent struct {
	start_x int
	start_y int
	end_x   int
	end_y   int
}

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

	var vents [500]vent

	for row, line := range lines {
		line_coords := strings.Split(line, " -> ")
		// iterate twice, 1x for start coords and 1x for end coords
		for iterator, line_coord := range line_coords {
			coords := strings.Split(line_coord, ",")
			if iterator == 0 {
				//start coords
				vents[row].start_x, err = strconv.Atoi(coords[0])
				vents[row].start_y, err = strconv.Atoi(coords[1])
			} else {
				//end coords
				vents[row].end_x, err = strconv.Atoi(coords[0])
				vents[row].end_y, err = strconv.Atoi(coords[1])
			}

		}
	}

	//filter only the straight vents where x1 = x2, y1 = y2, also switch start and end coords if end coord is samller of the non equal vector.
	// makes plotting and life easier
	straight_vents, non_straight_vents := filter_vents(vents)

	//fill grid
	fmt.Println("number of straight vents:", len(straight_vents))
	fmt.Println("number of non straight vents:", len(non_straight_vents))
	grid := plot_straight_vents(straight_vents)
	//get amount of "straight" intersections
	fmt.Println(count(grid))

	//plot the non straight vents now
	grid = plot_non_straight_vents(non_straight_vents, grid)
	fmt.Println(count(grid))
}

//seperate vents
func filter_vents(vents [500]vent) ([]vent, []vent) {
	var straight_vents []vent
	var non_straight_vents []vent

	for _, vent := range vents {
		if vent.start_x == vent.end_x || vent.start_y == vent.end_y {
			vent = switch_coords(vent)
			straight_vents = append(straight_vents, vent)
		} else {
			vent = switch_x_coords(vent)
			non_straight_vents = append(non_straight_vents, vent)
		}
	}
	return straight_vents, non_straight_vents
}

//switch the start and end coords
func switch_coords(vent vent) vent {
	if vent.end_x < vent.start_x {
		tmp := vent.start_x
		vent.start_x = vent.end_x
		vent.end_x = tmp
	}
	if vent.end_y < vent.start_y {
		tmp := vent.start_y
		vent.start_y = vent.end_y
		vent.end_y = tmp
	}
	return vent
}

//only switch x coords, only for the non straight vents
func switch_x_coords(vent vent) vent {
	if vent.end_x < vent.start_x {
		tmp := vent.start_x
		vent.start_x = vent.end_x
		vent.end_x = tmp
	}
	return vent
}

//create a grid with only straight vents plotted given ana array of vents
func plot_straight_vents(straight_vents []vent) [1000][1000]int {
	var grid [1000][1000]int
	for _, vent := range straight_vents {
		if vent.start_y == vent.end_y {
			for x := vent.start_x; x <= vent.end_x; x++ {
				grid[x][vent.start_y]++
			}
		} else if vent.start_x == vent.end_x {
			for y := vent.start_y; y <= vent.end_y; y++ {
				grid[vent.start_x][y]++
			}
		}
	}
	return grid
}

//take existing grid and plot
func plot_non_straight_vents(non_straight_vents []vent, grid [1000][1000]int) [1000][1000]int {
	for _, vent := range non_straight_vents {
		x_pos := vent.start_x
		if vent.start_y < vent.end_y {
			for y := vent.start_y; y <= vent.end_y; y++ {
				fmt.Println(vent, x_pos, y)
				grid[x_pos][y]++
				x_pos++
			}
		} else {
			for y := vent.start_y; y >= vent.end_y; y-- {
				grid[x_pos][y]++
				x_pos++
			}
		}

	}
	return grid
}

func count(grid [1000][1000]int) int {
	var count = 0
	//count "dangerous areas"
	for _, row := range grid {
		for _, column := range row {
			if column > 1 {
				count++
			}
		}
	}
	return count
}
