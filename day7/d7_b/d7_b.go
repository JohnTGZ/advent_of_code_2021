package main

import (
	"aoc/aoc"
	"bufio" //For reading line by line
	"flag"  //For command line parsing
	"fmt"
	"math"
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

	return num_arr
}

func getFuelUsage(num_arr []int, des_pos int) int {
	total_fuel := 0
	for _, val := range num_arr {
		total_fuel += aoc.ConsecutiveSum(val, des_pos)
	}
	return total_fuel
}

func main() {

	input_file := procArg()

	num_arr := getInput(input_file)

	//sort in ascending order
	sort.Ints(num_arr)

	min_fuel_usage := math.MaxInt64
	most_efficient_pos := -1

	//Brute force method
	for i := 1; i < num_arr[len(num_arr)-1]; i++ {
		fuel_used := getFuelUsage(num_arr, i)
		//get the minimum fuel used at each position
		if fuel_used < min_fuel_usage {
			min_fuel_usage = fuel_used
			most_efficient_pos = i
		}
		fmt.Printf("Position %d: Total fuel usage is %d \n", i, fuel_used)
	}

	// for i, val := range num_arr {
	// 	fuel_used := getFuelUsage(num_arr, val)
	// 	//get the minimum fuel used at each position
	// 	if fuel_used < min_fuel_usage {
	// 		min_fuel_usage = fuel_used
	// 		most_efficient_pos = val
	// 	}
	// 	fmt.Printf("Itr %d: Total fuel usage is %d \n", i, fuel_used)
	// }

	fmt.Printf("Minimum fuel %d used at %d \n", min_fuel_usage, most_efficient_pos)

}
