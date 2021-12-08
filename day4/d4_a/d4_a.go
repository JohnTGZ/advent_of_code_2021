package main

import (
	"bufio"   //For reading line by line
	"flag"    //For command line parsing
	"fmt"     //Converts string into integers
	"os"      //for opening filess
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
func getInput(filepath string) ([]string, int) {

	f, err := os.Open(filepath)
	check(err)

	//close file at end of program
	defer f.Close()

	//create scanner object to read line by line
	scanner := bufio.NewScanner(f)

	var input_arr []string

	num_boards := 0

	//read file line by line
	for scanner.Scan() {
		text := scanner.Text()
		//check if it is empty
		if len(text) == 0 {
			fmt.Printf("Empty line \n")
		}
		input_arr = append(input_arr, text)
	}

	return input_arr, num_boards
}

/*
Function to iterate through arrays
and print each entry
*/
func printArr(str_arr []string) {

	for _, arr := range str_arr {
		fmt.Printf("%s, ", arr)
	}
	fmt.Printf("\n")

}

func main() {

	input_file := procArg()

	input_arr, _ := getInput(input_file)
	// board_sz := 5

	//read first line and split into an array
	draw_arr := strings.Split(input_arr[0], ",")
	printArr(draw_arr)

	//create array of boards
	// var boards [num_boards][board_sz * board_sz]int

	// //Iterate through remaining lines and save board info
	// for i, line := range input_arr {

	// }

	// o2_dec, _ := strconv.ParseInt(o2_bin, 2, 32)
	// co2_dec, _ := strconv.ParseInt(co2_bin, 2, 32)

	// fmt.Printf("o2_bin: %s, dec: %d \n", o2_bin, o2_dec)
	// fmt.Printf("co2_bin: %s, dec: %d \n", co2_bin, co2_dec)

	// fmt.Printf("answer: %d\n", o2_dec*co2_dec)

}
