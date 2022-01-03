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
func getInput(filepath string) ([]string, []string) {

	f, err := os.Open(filepath)
	aoc.CheckErr(err)

	//close file at end of program
	defer f.Close()

	//create scanner object to read line by line
	scanner := bufio.NewScanner(f)

	var unique_signal_patterns []string
	var digital_output_val []string

	for scanner.Scan() {
		text := scanner.Text()

		//split into unique signal patterns and four digit output val
		char_pair := strings.Split(text, " | ")

		unique_signal_patterns = append(unique_signal_patterns, char_pair[0])
		digital_output_val = append(digital_output_val, char_pair[1])
	}

	return unique_signal_patterns, digital_output_val
}

func main() {

	input_file := procArg()

	_, digital_output_val := getInput(input_file)

	//2 segments: 1
	//3 seg.: 7
	//4 seg.: 4
	//5 seg.: 2,3,5
	//6 seg.: 0, 6, 9
	//7 seg.: 8

	total_instances := 0

	for _, arr := range digital_output_val {
		pattern_arr := strings.Split(arr, " ")

		fmt.Printf("____________ \n")
		for _, pattern := range pattern_arr {
			switch len(pattern) {
			case 2:
				fmt.Printf("1: %s", pattern)
				total_instances += 1
			case 3:
				fmt.Printf("7: %s", pattern)
				total_instances += 1
			case 4:
				fmt.Printf("4: %s", pattern)
				total_instances += 1
			case 5:
				fmt.Printf("2, 3, 5: %s", pattern)
			case 6:
				fmt.Printf("0, 6, 9: %s", pattern)
			case 7:
				fmt.Printf("8: %s", pattern)
				total_instances += 1
			default:
				fmt.Printf("Invalid! Check input! %s", pattern)
			}
			fmt.Printf("\n")
		}
	}

	fmt.Printf("Printing total instances of 1, 4, 7 and 8 in output values: %d \n", total_instances)
	// aoc.PrintArrStr(unique_signal_patterns)

	// fmt.Printf("Printing digital_output_val... \n")
	// aoc.PrintArrStr(digital_output_val)

}
