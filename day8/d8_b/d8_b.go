package main

import (
	"aoc/aoc"
	"bufio" //For reading line by line
	"flag"  //For command line parsing
	"fmt"
	"os" //for opening filess
	"strconv"
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
func getInput(filepath string) ([]string, []string) {

	f, err := os.Open(filepath)
	aoc.CheckErr(err)

	//close file at end of program
	defer f.Close()

	//create scanner object to read line by line
	scanner := bufio.NewScanner(f)

	var all_patterns_str_sorted []string
	var all_output_str_sorted []string

	for scanner.Scan() {
		text := scanner.Text()

		//split into unique signal patterns and four digit output val
		char_pair := strings.Split(text, " | ")
		patterns_arr := strings.Split(char_pair[0], " ")
		output_arr := strings.Split(char_pair[1], " ")

		var patterns_arr_sorted, output_arr_sorted []string

		for _, str := range patterns_arr {
			patterns_arr_sorted = append(patterns_arr_sorted, aoc.BubbleSortRunes([]rune(str)))
		}
		for _, str := range output_arr {
			output_arr_sorted = append(output_arr_sorted, aoc.BubbleSortRunes([]rune(str)))
		}

		patterns_str_sorted := strings.Join(patterns_arr_sorted, " ")
		output_str_sorted := strings.Join(output_arr_sorted, " ")

		//Combine to string again
		all_patterns_str_sorted = append(all_patterns_str_sorted, patterns_str_sorted)
		all_output_str_sorted = append(all_output_str_sorted, output_str_sorted)
	}

	return all_patterns_str_sorted, all_output_str_sorted
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

//Get difference between 2 strings
//assumes that strings are sorted
func getStringDiff(s1_string string, s2_string string) string {
	var differing_str string

	if len(s2_string) > len(s1_string) {
		s1_string, s2_string = s2_string, s1_string
	}

	//for each rune in s1, compare against the rest of s2
	for _, s1_rune := range s1_string {
		s2_contains_current_rune := false
		for _, s2_rune := range s2_string {
			if s1_rune == s2_rune {
				// fmt.Printf("%c == %c \n", s1_rune, s2_rune)
				s2_contains_current_rune = true
				//Runes are equal
				break
			}
		}
		//If s2 does not contain the rune of s1, addo to carto
		if !s2_contains_current_rune {
			// fmt.Printf("Added difference %s \n", string(s1_rune))
			differing_str += string(s1_rune)
		}
	}

	return differing_str
}

/*
CHecks if string s2 is inside string s1
*/
func stringContains(s1_string string, s2_string string) bool {
	var similar_str string

	// if len(s2_string) > len(s1_string) {
	// 	s1_string, s2_string = s2_string, s1_string
	// }

	for _, s1_rune := range s1_string {
		s1_contains_current_rune := false
		for _, s2_rune := range s2_string {
			if s1_rune == s2_rune {
				s1_contains_current_rune = true
				//Runes are equal
				break
			}
		}
		//If s1 contains the rune of s2, addo to carto
		if s1_contains_current_rune {
			similar_str += string(s1_rune)
		}
	}

	// fmt.Printf("similar_Str %s \n", similar_str)

	if len(similar_str) != len(s2_string) {
		return false
	}

	return true
}

// func main() {
// 	// s1_string := "bc"
// 	// s2_string := "abcd"

// 	// fmt.Printf("Different string: %s \n", getStringDiff(s1_string, s2_string))

// 	s1_string := "abcdeg"
// 	s2_string := "dg"

// 	fmt.Printf("Contains string: %t \n", stringContains(s1_string, s2_string))
// }

func main() {

	input_file := procArg()

	all_patterns_str_sorted, all_output_str_sorted := getInput(input_file)

	//2 segments: 1
	//3 seg.: 7
	//4 seg.: 4
	//5 seg.: 2,3,5
	//6 seg.: 0, 6, 9
	//7 seg.: 8

	// total_instances := 0

	var output_number_slice []string

	for i, arr := range all_patterns_str_sorted {

		patterns_slice := strings.Split(arr, " ")

		fmt.Printf("Iteration %d \n", i)
		//Iterate through each pattern and get the set of letters that represent
		//the following components
		//alpha = "7" - "1"
		//beta = "4" - "1"
		//gamma = "1"
		//delta = "8" - "4" - alpha

		//Get numbers containing the following components
		//2: delta
		//3: gamma
		//5: beta
		//0: gamma, delta
		//6: beta, delta
		//9: beta, gamma

		number_str := make([]string, 10)

		var remaining_numbers_str []string

		for _, pattern := range patterns_slice {

			switch len(pattern) {
			case 2:
				// fmt.Printf("1: %s", pattern)
				number_str[1] = pattern
			case 3:
				// fmt.Printf("7: %s", pattern)
				number_str[7] = pattern
			case 4:
				// fmt.Printf("4: %s", pattern)
				number_str[4] = pattern
			case 7:
				// fmt.Printf("8: %s", pattern)
				number_str[8] = pattern
			case 5:
				// fmt.Printf("2, 3, 5: %s", pattern)
				remaining_numbers_str = append(remaining_numbers_str, pattern)
			case 6:
				// fmt.Printf("0, 6, 9: %s", pattern)
				remaining_numbers_str = append(remaining_numbers_str, pattern)
			}
			// fmt.Printf("\n")
		}

		alpha := getStringDiff(number_str[7], number_str[1])
		beta := getStringDiff(number_str[4], number_str[1])
		gamma := number_str[1]
		delta := getStringDiff(getStringDiff(number_str[8], number_str[4]), alpha)

		// fmt.Printf("Beta: %s \n", beta)
		// fmt.Printf("gamma: %s \n", gamma)
		// fmt.Printf("delta: %s \n", delta)

		for _, pattern := range remaining_numbers_str {
			b := stringContains(pattern, beta)
			g := stringContains(pattern, gamma)
			d := stringContains(pattern, delta)
			// fmt.Printf("Current pattern (%s): b(%t), g(%t), d(%t) \n", pattern, b, d, g)

			if b && g {
				number_str[9] = pattern
			} else if b && d {
				number_str[6] = pattern
			} else if g && d {
				number_str[0] = pattern
			} else if b {
				number_str[5] = pattern
			} else if g {
				number_str[3] = pattern
			} else {
				number_str[2] = pattern
			}
		}
		//print out all numbers
		// aoc.PrintArrStr(number_str)

		//decode the output number

		outputs_slice := strings.Split(all_output_str_sorted[i], " ")

		output_number_str := ""
		for _, num_pattern := range outputs_slice {
			//iterate through each output and determine the number
			for i, actual_number_pattern := range number_str {
				if num_pattern == actual_number_pattern {
					// fmt.Printf("Appended %d using %s\n", i, num_pattern)
					output_number_str += strconv.Itoa(i)
				}
			}
		}
		fmt.Printf("Decoded output: %s\n", output_number_str)
		output_number_slice = append(output_number_slice, output_number_str)
	}
	aoc.PrintArrStr(output_number_slice)

	sum := 0
	for _, num_str := range output_number_slice {
		num, _ := strconv.Atoi(num_str)
		sum += num
	}
	fmt.Printf("Final sum: %d \n", sum)
}
