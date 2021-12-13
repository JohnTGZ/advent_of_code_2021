package main

import (
	"aoc/aoc"
	"bufio" //For reading line by line
	"flag"  //For command line parsing
	"fmt"
	"os" //for opening filess
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

	return num_arr
}

/*
Step through the lantern fish reproduction
1 day at a time
*/
func stepThrough(num_arr []int) []int {
	next_day_num_arr := num_arr

	// fmt.Printf("Size of arr: %d \n", len(next_day_num_arr))
	for i, _ := range num_arr {
		next_day_num_arr[i] -= 1

		//Time to poop babies
		if next_day_num_arr[i] < 0 {
			next_day_num_arr[i] = 6
			//add new baby
			next_day_num_arr = append(next_day_num_arr, 8)
			// fmt.Printf("Size after append: %d \n", len(next_day_num_arr))
		}

	}

	return next_day_num_arr
}

func main() {

	input_file := procArg()

	num_arr := getInput(input_file)

	//step through 1 day at a time
	for i := 0; i < 80; i++ {
		// fmt.Printf("Iteration %d", i)
		num_arr = stepThrough(num_arr)
		// aoc.PrintArrInt(num_arr)
	}

	// fmt.Printf("Sum of lanternfishes: %d \n", aoc.GetSum(num_arr))
	fmt.Printf("Sum of lanternfishes: %d \n", len(num_arr))

}
