//Part 2

package main

import (
	"bufio" //For reading line by line
	"flag"  //For command line parsing
	"fmt"   //Converts string into integers
	"math"
	"os" //for opening filess
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
func getInput(filepath string) ([]string, int) {

	f, err := os.Open(filepath)
	check(err)

	//close file at end of program
	defer f.Close()

	//read file line by line
	scanner := bufio.NewScanner(f)

	var input_arr []string

	input_sz := 0

	for scanner.Scan() {
		full_cmd := scanner.Text()
		input_arr = append(input_arr, full_cmd)

		input_sz++
	}

	return input_arr, input_sz
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

	//get size of binary number
	bin_sz := len(input_arr[0])

	o2_bin := ""
	co2_bin := ""

	//copy input
	input_arr2 := input_arr

	//iterate through each binary position
	for i := 0; i < bin_sz; i++ {

		var fil_arr_0 []string
		var fil_arr_1 []string
		ones_sum := 0
		input_sz := 0

		//iterate through each remaining entry
		for _, arr := range input_arr {
			bin_num_arr := strings.Split(arr, "")

			if bin_num_arr[i] == "1" {
				fil_arr_1 = append(fil_arr_1, arr)
				// fmt.Printf("	printing fil_arr_1: ")
				// printArr(fil_arr_1)
				ones_sum++
			} else {
				fil_arr_0 = append(fil_arr_0, arr)
				// fmt.Printf("	printing fil_arr_0: ")
				// printArr(fil_arr_0)
			}
			input_sz++
		}
		//reduce input array to only the entries with
		//the most common binary number in the current position
		if ones_sum >= int(math.Ceil((float64(input_sz) / 2.0))) {
			//if 1 is more common
			input_arr = fil_arr_1
		} else {
			//if 0 is more common
			input_arr = fil_arr_0
		}

		var fil_arr_0_2 []string
		var fil_arr_1_2 []string
		ones_sum_2 := 0
		input_sz_2 := 0

		//iterate through each remaining entry
		for _, arr := range input_arr2 {
			bin_num_arr := strings.Split(arr, "")

			if bin_num_arr[i] == "1" {
				fil_arr_1_2 = append(fil_arr_1_2, arr)
				// fmt.Printf("	printing fil_arr_1_2: ")
				// printArr(fil_arr_1_2)
				ones_sum_2++
			} else {
				fil_arr_0_2 = append(fil_arr_0_2, arr)
				// fmt.Printf("	printing fil_arr_0_2: ")
				// printArr(fil_arr_0_2)
			}
			input_sz_2++
		}
		//reduce input array to only the entries with
		//the least common binary number in the current position
		if ones_sum_2 >= int(math.Ceil((float64(input_sz_2) / 2.0))) {
			input_arr2 = fil_arr_0_2
		} else {
			input_arr2 = fil_arr_1_2
		}

		if len(input_arr) == 1 {
			o2_bin = input_arr[0]
		}
		if len(input_arr2) == 1 {
			co2_bin = input_arr2[0]
		}

	}

	o2_dec, _ := strconv.ParseInt(o2_bin, 2, 32)
	co2_dec, _ := strconv.ParseInt(co2_bin, 2, 32)

	fmt.Printf("o2_bin: %s, dec: %d \n", o2_bin, o2_dec)
	fmt.Printf("co2_bin: %s, dec: %d \n", co2_bin, co2_dec)

	fmt.Printf("answer: %d\n", o2_dec*co2_dec)

}
