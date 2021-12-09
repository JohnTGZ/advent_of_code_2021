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
	// getPos() int16
	// fillNum() int16
	printBoard()
	addBoard()
}

type board struct {
	idx     int
	val_arr []int
	width   int
	height  int
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

//todo: convert array of str into array of int

func main() {

	input_file := procArg()

	input_arr, num_lines, num_boards := getInput(input_file)

	fmt.Printf("No. of boards: %d \n", num_boards)

	//constant variables
	board_width := 5
	board_height := 5

	//read first line and split into an array
	// draw_num_arr := strings.Split(input_arr[0], ",")
	// printArrStr(draw_num_arr)

	//line 0: numbers being drawn
	// 	...
	//line 1, 7, 13, 19: Empty line
	// 	General formulat: 1+index*6
	//line 2-6, 8-12, 14-18: board data
	//	General formula: 2+index*6 -> 6*(index+1)

	boards := make([]board, num_boards)

	current_board_idx := 0
	var board_input []int

	//iterate through all lines
	for i := 2; i < num_lines+1; i++ {
		if i == 7+6*(current_board_idx) {
			//empty line
			boards[current_board_idx] = board{idx: current_board_idx, val_arr: board_input, width: board_width, height: board_height}
			board_input = nil
			current_board_idx++

		} else {
			//get board data
			for _, val := range strings.Fields(input_arr[i]) {
				val_int, err := strconv.Atoi(val)
				check(err)
				board_input = append(board_input, val_int)
			}
		}

	}

	//print all boards
	for _, current_board := range boards {
		current_board.printBoard()
	}

}
