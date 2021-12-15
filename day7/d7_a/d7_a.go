package main

import (
	"aoc/aoc"
	"bufio" //For reading line by line
	"flag"  //For command line parsing
	"fmt"
	"os" //for opening filess
	"sort"
	"strings"
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

/* Read the input into an iterable array
 */
func getInput(filepath string) []int {

	f, err := os.Open(filepath)
	aoc.CheckErr(err)

	//close file at end of program
	defer f.Close()

	//create scanner object to read line by line
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	text := scanner.Text()
	//split into array
	num_arr_str := strings.Split(text, ",")

	var num_arr []int
	num_arr = aoc.ConvArrStrToInt(num_arr_str)

	//sort in ascending order
	sort.Ints(num_arr)

	return num_arr
}

func getFuelUsage(num_arr []int, des_pos int) int {
	total_fuel := 0
	for _, val := range num_arr {
		if val <= des_pos {
			total_fuel += des_pos - val
		} else {
			total_fuel += val - des_pos
		}
	}
	return total_fuel
}

func main() {

	input_file := procArg()

	num_arr := getInput(input_file)

	median := aoc.GetMedian(num_arr)

	fmt.Printf("Median: %d \n", median)

	fmt.Printf("Total fuel usage: %d \n", getFuelUsage(num_arr, median))

}
