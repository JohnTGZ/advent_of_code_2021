package main

import (
	"bufio" //For reading line by line
	"flag"  //For command line parsing
	"fmt"   //Converts string into integers
	"os"    //for opening filess
	"strconv"
	"strings" //Converts string into integers
)

/* Process command line arguemnts to read from either
test input or normal input file
*/
func procArg() string {
	testingPtr := flag.Bool("t", false, "Enable testing")

	flag.Parse()

	input_file := "../input.txt"
	if *testingPtr {
		input_file = "../test_input.txt"
	}
	return input_file
}

/* Checks error and PANIK!!
 */
func check(e error) {
	if e != nil {
		panic(e)
	}
}

/* Read the input into an iterable array
 */
func getInput(filepath string) ([]string, int, int) {

	f, err := os.Open(filepath)
	check(err)

	//close file at end of program
	defer f.Close()

	//create scanner object to read line by line
	scanner := bufio.NewScanner(f)

	var input_arr []string

	num_lines := 0

	//read file line by line
	for scanner.Scan() {
		text := scanner.Text()

		input_arr = append(input_arr, text)
		num_lines++
	}

	num_boards := (num_lines - 1) / 6

	return input_arr, num_lines, num_boards
}

/*
Function to iterate through arrays
and print each entry
*/
func printArrInt(str_arr []int) {

	for _, arr := range str_arr {
		fmt.Printf("%d, ", arr)
	}
	fmt.Printf("\n")

}

func printArrStr(str_arr []string) {

	for _, arr := range str_arr {
		fmt.Printf("%s, ", arr)
	}
	fmt.Printf("\n")

}

type board_interface interface {
	fillNum()
	printBoard()
}

type board struct {
	idx        int
	val_arr    []int
	width      int
	height     int
	filled_pos []int //Positions that have been filled up
}

func (b board) printBoard() {
	fmt.Printf("===Printing board %d===", b.idx)

	for i, val := range b.val_arr {
		if i%b.width == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%d ", val)
	}
	fmt.Printf("\n")
}

/*
Fill up the board with the number
*/
func (b board) fillNum(drawn_num int) {
	//modify the val_arr[]int
	//add position of value to filled_pos

	for i, val := range b.val_arr {
		if val == drawn_num {
			fmt.Printf("i(%d), drawn(%d) \n", val, drawn_num)
			(&b).filled_pos = append((&b).filled_pos, i)
			(&b).val_arr[i] = 666
		}

	}

}

/*
Fill up the board with the number
*/
// func (b board) checkBingo(drawn_num int) {
// 	//modify the val_arr[]int
// 	//add position of value to filled_pos

// 	for i, val := range b.val_arr {
// 		if val == drawn_num {
// 			// fmt.Printf("val(%d), drawn(%d) \n", val, drawn_num)
// 			(&b).filled_pos = append(b.filled_pos, i)
// 			(&b).val_arr[i] = 666
// 		}

// 	}

// }

/*
Search through mapping and fill up the relevant boards with the drawn number
*/
func drawNum(boards []board, board_map map[int][]int, drawn_num int) {
	for _, board_idx := range board_map[drawn_num] {
		boards[board_idx].fillNum(drawn_num)
	}
}

//todo: convert array of str into array of int

func main() {

	input_file := procArg()

	input_arr, num_lines, num_boards := getInput(input_file)

	fmt.Printf("No. of boards: %d \n", num_boards)

	//constant variables
	board_width := 5
	board_height := 5

	//read first line and split into an array
	draw_num_arr := strings.Split(input_arr[0], ",")
	printArrStr(draw_num_arr)

	//line 0: numbers being drawn
	// 	...
	//line 1, 7, 13, 19: Empty line
	// 	General formulat: 1+index*6
	//line 2-6, 8-12, 14-18: board data
	//	General formula: 2+index*6 -> 6*(index+1)

	//initialize data structures
	boards := make([]board, num_boards)
	board_map := make(map[int][]int) //maps numbers to the boards they belong to

	current_board_idx := 0
	var board_input []int

	//iterate through all lines
	for i := 2; i < num_lines+1; i++ {
		if i == 7+6*(current_board_idx) {
			//Line: empty line

			//Initialize board data
			boards[current_board_idx] = board{idx: current_board_idx, val_arr: board_input, width: board_width, height: board_height}

			//add to map
			for _, input := range board_input {
				board_map[input] = append(board_map[input], current_board_idx)
			}

			board_input = nil
			current_board_idx++

		} else {
			//Line: board data

			for _, val := range strings.Fields(input_arr[i]) {
				val_int, err := strconv.Atoi(val)
				check(err)
				board_input = append(board_input, val_int)
			}
		}

	}

	for _, drawn_num := range draw_num_arr {
		drawn_num_int, _ := strconv.Atoi(drawn_num)
		drawNum(boards, board_map, drawn_num_int)
	}

	//print all boards
	fmt.Printf("Printing boards... \n ")
	for _, current_board := range boards {
		current_board.printBoard()
	}
	printArrInt(boards[0].filled_pos)

	//print all mappings
	// fmt.Printf("Printing mappings... \n ")
	// for key, arr := range board_map {
	// 	fmt.Printf("Number %d: \n", key)
	// 	for _, val := range arr {
	// 		fmt.Printf("%d, ", val)
	// 	}
	// 	fmt.Printf("\n ")
	// }

}
