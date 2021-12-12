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
func printArrInt(int_arr []int) {
	for _, arr := range int_arr {
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

func intInSlice(des_val int, int_arr []int) bool {
	for _, val := range int_arr {
		if val == des_val {
			return true
		}
	}
	return false
}

type board_interface interface {
	fillNum()
	printBoard()
}

type Board struct {
	idx     int
	val_arr []int //currently available values
	width   int
	height  int

	rows [5][]int
	cols [5][]int
}

func (b *Board) init(idx int, val_arr []int, width int, height int) {
	b.idx = idx
	b.val_arr = val_arr
	b.width = width
	b.height = height
}

/*
Pretty print the board out in 2d presentation
*/
func (b *Board) printBoard() {
	fmt.Printf("===Printing Board %d===", b.idx)

	for i, val := range b.val_arr {
		if i%b.width == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%d ", val)
	}
	fmt.Printf("\n")
}

/*
Add to b.rows and b.cols which are used to check bingo
*/
func (b *Board) addToRowColArr(position int) {
	row_idx := int(position / b.width)
	col_idx := int(position % b.height)
	// fmt.Printf("Col(%d), Horz(%d): drawn(%d) \n", col_idx, row_idx, position)

	b.rows[row_idx] = append(b.rows[row_idx], position)
	b.cols[col_idx] = append(b.cols[col_idx], position)
}

/*
Fill up the Board with the number
*/
func (b *Board) fillNum(drawn_num int) {
	//modify the val_arr[]int
	for i, val := range b.val_arr {
		if val == drawn_num {
			b.addToRowColArr(i)
			b.val_arr[i] = -1
		}
	}
}

/*
Check if board has bingoed
*/
func (b *Board) checkBingo() bool {
	//check horizontal row
	for _, horz_arr := range b.rows {
		if len(horz_arr) >= b.width {
			return true
		}
	}
	//check vertical row
	for _, col_arr := range b.cols {
		if len(col_arr) >= b.height {
			return true
		}
	}
	return false
}

/*
Search through mapping and fill up the relevant boards with the drawn number
*/
func drawNum(boards []Board, board_map map[int][]int, drawn_num int) {
	//for each board with the drawn number, fill it in
	for _, board_idx := range board_map[drawn_num] {
		boards[board_idx].fillNum(drawn_num)
	}
}

/*
Checks every board for a bingo, except for those that have already won
returns an array of bingo boards
*/
func checkBingos(boards []Board, boards_won []int) (bool, []int) {
	var bingo_boards []int
	for idx, board := range boards {
		if !intInSlice(idx, boards_won) {
			if board.checkBingo() {
				bingo_boards = append(bingo_boards, idx)
			}
		}
	}

	if len(bingo_boards) > 0 {
		return true, bingo_boards
	} else {
		return false, bingo_boards
	}
}

func main() {

	input_file := procArg()

	input_arr, num_lines, num_boards := getInput(input_file)

	final_drawn_num_int := 0
	last_remaining_num := 0

	var boards_won []int

	fmt.Printf("No. of boards: %d \n", num_boards)

	//constant variables
	board_width := 5
	board_height := 5

	//read first line and split into an array
	draw_num_arr := strings.Split(input_arr[0], ",")
	// printArrStr(draw_num_arr)

	//line 0: numbers being drawn
	// 	...
	//line 1, 7, 13, 19: Empty line
	// 	General formulat: 1+index*6
	//line 2-6, 8-12, 14-18: Board data
	//	General formula: 2+index*6 -> 6*(index+1)

	//initialize data structures
	boards := make([]Board, num_boards)
	board_map := make(map[int][]int) //maps numbers to the boards they belong to

	current_board_idx := 0
	var board_input []int

	//iterate through all lines
	for i := 2; i < num_lines+1; i++ {
		if i == 7+6*(current_board_idx) {
			//Line: empty line

			//Initialize Board data
			boards[current_board_idx].init(current_board_idx, board_input, board_width, board_height)

			//add to map
			for _, input := range board_input {
				board_map[input] = append(board_map[input], current_board_idx)
			}

			board_input = nil
			current_board_idx++

		} else {
			//Line: Board data

			for _, val := range strings.Fields(input_arr[i]) {
				val_int, err := strconv.Atoi(val)
				check(err)
				board_input = append(board_input, val_int)
			}
		}

	}

	//draw the numbers and fill in the boards
	for _, drawn_num := range draw_num_arr {
		drawn_num_int, _ := strconv.Atoi(drawn_num)
		drawNum(boards, board_map, drawn_num_int)

		bingo, win_boards := checkBingos(boards, boards_won)

		if bingo {
			for _, win_idx := range win_boards {
				boards_won = append(boards_won, win_idx)

				current_remaining_num := 0
				for _, val := range boards[win_idx].val_arr {
					if val != -1 {
						current_remaining_num += val
					}
				}
				last_remaining_num = current_remaining_num
				final_drawn_num_int = drawn_num_int
				fmt.Printf("Last board %d bingoed with drawn num %d \n", boards_won[len(boards_won)-1], final_drawn_num_int)
				fmt.Printf("Sum of remaining num: %d \n", last_remaining_num)
			}

		}
	}
	fmt.Printf("Sum of remaining num: %d \n", last_remaining_num*final_drawn_num_int)

	printArrInt(boards_won)

	//print boards
	// fmt.Printf("Printing boards... \n ")
	// for _, current_board := range boards {
	// 	current_board.printBoard()
	// }

}
