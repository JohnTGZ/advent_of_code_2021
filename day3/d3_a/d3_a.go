//Part 2

package main

import (
	"bufio" //For reading line by line
	"flag"  //For command line parsing
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

	//get size of binary number
	bin_sz := len(input_arr[0])

	var ones_sum = make([]int, bin_sz)

	for i := 0; i < input_sz; i++ {
		bin_num_arr := strings.Split(input_arr[i], "")

		for i, bin_ := range bin_num_arr {
			if bin_ == "1" {
				ones_sum[i]++
			}
		}
	}

	gamma_bin := ""
	eps_bin := ""

	for _, num := range ones_sum {
		if num > input_sz/2 {
			gamma_bin += "1"
			eps_bin += "0"
		} else {
			gamma_bin += "0"
			eps_bin += "1"
		}
	}
	gamma_dec, _ := strconv.ParseInt(gamma_bin, 2, 32)
	eps_dec, _ := strconv.ParseInt(eps_bin, 2, 32)

	fmt.Printf("gamma_bin: %s, dec: %d \n", gamma_bin, gamma_dec)
	fmt.Printf("eps_bin: %s, dec: %d \n", eps_bin, eps_dec)

	fmt.Printf("answer: %d\n", gamma_dec*eps_dec)

}
