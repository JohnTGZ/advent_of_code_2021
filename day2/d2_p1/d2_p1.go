//Part 2

package main

import (
	"bufio" //For reading line by line
	"flag"

	//For command line parsing
	"fmt"
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

	fmt.Printf("%t \n", *testingPtr)

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

func main() {
	input_file := procArg()

	input_arr, input_sz := getInput(input_file)

	forward := 0
	depth := 0

	for i := 0; i < input_sz; i++ {
		//split the full command into cmd and value
		cmd_split := strings.Split(input_arr[i], " ")

		cmd := cmd_split[0]
		value, err := strconv.Atoi(cmd_split[1])
		check(err)

		//Increment command sum
		switch cmd {
		case "forward":
			forward += value
		case "up":
			depth -= value
		case "down":
			depth += value
		}
	}

	fmt.Printf("depth: %d, forward: %d, answer: %d \n", depth, forward, depth*forward)
}
