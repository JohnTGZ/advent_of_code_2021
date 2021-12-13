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
func stepThrough(fish_map []int) {

	num_fishes_gonna_reborn := fish_map[0]
	//move number of fishes one index down
	for i := 0; i < 8; i++ {
		// fmt.Printf("iteration: %d(%d) <- %d(%d) \n", i, fish_map[i], i+1, fish_map[i+1])
		fish_map[i] = fish_map[i+1]
	}

	//fish are reborn
	fish_map[6] += num_fishes_gonna_reborn
	//new fish are pooped out
	fish_map[8] = num_fishes_gonna_reborn
}

func main() {

	input_file := procArg()

	num_arr := getInput(input_file)

	//create mapping of remaining linespan to number of fishes
	lifespan_to_fishes_map := make([]int, 9)

	//iterate through the range of lifespan values
	for _, val := range num_arr {
		lifespan_to_fishes_map[val] += 1
	}

	//step through 1 day at a time
	for i := 0; i < 256; i++ {
		// fmt.Printf("Iteration %d \n", i)
		stepThrough(lifespan_to_fishes_map)
		// aoc.PrintArrInt(lifespan_to_fishes_map)

	}

	fmt.Printf("Sum of lanternfishes: %d \n", aoc.GetSum(lifespan_to_fishes_map))

}
