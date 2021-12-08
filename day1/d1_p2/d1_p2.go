//Part 2

package main

import (
	//for printing
	"fmt"
	"strconv"

	// "io"
	"bufio" //For reading line by line
	"os"    //for opening filess
	//Converts string into integers
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInput(filepath string) ([]int, int) {

	f, err := os.Open(filepath)
	check(err)

	//close file at end of program
	defer f.Close()

	//read file line by line
	scanner := bufio.NewScanner(f)

	var input_arr []int

	input_sz := 0
	for scanner.Scan() {
		next_number, err := strconv.Atoi(scanner.Text())
		check(err)

		input_arr = append(input_arr, next_number)

		input_sz++
	}

	return input_arr, input_sz
}

func main() {

	deeper_counter := 0

	input_arr, input_sz := getInput("../input.txt")

	for i := 0; i < input_sz-3; i++ {
		prev_sum := 0
		next_sum := 0

		//current sum
		for j := i; j < i+3; j++ {
			prev_sum += input_arr[j]
		}
		//next sum
		for j := i + 1; j < i+4; j++ {
			next_sum += input_arr[j]
		}

		if next_sum > prev_sum {
			fmt.Printf("%d > %d \n", next_sum, prev_sum)
			deeper_counter++
		} else {
			fmt.Printf("%d < %d \n", next_sum, prev_sum)
		}
	}

	fmt.Printf("Deeper counter: %d \n", deeper_counter)
}
