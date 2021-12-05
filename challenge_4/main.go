package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type board struct {
	values [5][5]string
	pulled [5][5]bool
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

	var numbers []string
	var board_index int
	var row_index int
	var boards [100]board

	for pos := range lines {
		if pos == 0 {
			numbers = strings.Split(lines[pos], ",")

			// check for newlines
		} else if len(lines[pos]) == 0 {
			continue
		} else {
			row := strings.Fields(lines[pos])
			fmt.Println(row)
			for charpos, char := range row {
				fmt.Println(board_index, row_index, charpos)
				boards[board_index].values[row_index][charpos] = char
				boards[board_index].pulled[row_index][charpos] = false

			}
			fmt.Println("current board row:", boards[board_index].values[row_index])
			row_index++
			if row_index >= 5 {
				row_index = 0
				board_index++
			}
		}
	}

	//so after filling up our arrays, time to check which board has won

	var board_number = 0
	var last_number_drawn = 0

out:
	for _, value := range numbers {
		for i := 0; i < len(boards); i++ {
			for column := 0; column < 5; column++ {
				for row := 0; row < 5; row++ {
					if boards[i].values[column][row] == value {
						boards[i].pulled[column][row] = true
					}
					// check for columns
					if boards[i].pulled[0][row] == true && boards[i].pulled[1][row] == true && boards[i].pulled[2][row] == true && boards[i].pulled[3][row] == true && boards[i].pulled[4][row] == true {
						board_number = i
						fmt.Println("columns:", boards[i].pulled)
						last_number_drawn, err = strconv.Atoi(value)
						break
					}
					if boards[i].pulled[column][0] == true && boards[i].pulled[column][1] == true && boards[i].pulled[column][2] == true && boards[i].pulled[column][3] == true && boards[i].pulled[column][4] == true {
						board_number = i
						fmt.Println("row:", boards[i].pulled)
						last_number_drawn, err = strconv.Atoi(value)
						break out
					}
				}
			}
		}
	}

	//check what is pulled
	fmt.Println(board_number)

	//get summation of the board_number
	summation := calc_summation(boards, board_number)

	fmt.Println("sum", summation)
	fmt.Println("last number drawn", last_number_drawn)
	fmt.Println(summation * last_number_drawn)

	//======================================end of part 1 but honestly lazy atm to split this, so part two start here :)

	var boards_that_won [100]int
	var index = 0

	for _, value := range numbers {
		for i := 0; i < len(boards); i++ {
		out_board:
			for column := 0; column < 5; column++ {
				for row := 0; row < 5; row++ {
					if boards[i].values[column][row] == value {
						boards[i].pulled[column][row] = true
					}
					// check for columns
					if boards[i].pulled[0][row] == true && boards[i].pulled[1][row] == true && boards[i].pulled[2][row] == true && boards[i].pulled[3][row] == true && boards[i].pulled[4][row] == true {
						board_number = i
						last_number_drawn, err = strconv.Atoi(value)
						for _, won_board := range boards_that_won {
							if won_board == i {
								break out_board
							}
						}
						boards_that_won[index] = i
						fmt.Println("added", i, "to board")
						fmt.Println("last number that got drawn:", value)
						index++
					}
					if boards[i].pulled[column][0] == true && boards[i].pulled[column][1] == true && boards[i].pulled[column][2] == true && boards[i].pulled[column][3] == true && boards[i].pulled[column][4] == true {
						board_number = i
						last_number_drawn, err = strconv.Atoi(value)
						for _, won_board := range boards_that_won {
							if won_board == i {
								break out_board
							}
						}
						boards_that_won[index] = i
						fmt.Println("added", i, "to board")
						fmt.Println("last number that got drawn:", value)
						index++
					}
				}
			}
		}
	}

	fmt.Println(boards_that_won)

	last_board_that_won := boards_that_won[len(boards_that_won)-2]
	fmt.Println(last_board_that_won, last_number_drawn)
	result := calc_summation(boards, last_board_that_won)

	fmt.Println(result * 64)
}

func calc_summation(boards [100]board, board_number int) int {
	summation := 0
	//get summation of the board_number
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if boards[board_number].pulled[i][j] == false {
				number, err := strconv.Atoi(boards[board_number].values[i][j])
				if err != nil {
					fmt.Println(err)
				}
				summation += number
			}
		}
	}
	return summation
}
