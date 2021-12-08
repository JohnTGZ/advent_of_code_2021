
package main

import (
	"fmt" //for printing
	// "io"
	"bufio" //For reading line by line
	"os" //for opening filess
	"strconv" //Converts string into integers
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main(){
	// input, err := os.ReadFile("./test_input.txt")
	// check(err)
	// fmt.Print(string(input))

	f, err := os.Open("../input.txt")
	check(err)

	//close file at end of program
	defer f.Close()

	//read file line by line
	scanner := bufio.NewScanner(f)

	var prev int 
	var next int 
	var counter int = 0

	//scan the first integer
	scanner.Scan()
	prev, err = strconv.Atoi(scanner.Text())
	check(err)

	for scanner.Scan(){
		next, err = strconv.Atoi(scanner.Text())
		check(err)
		if (next > prev){
			counter++
		}
		prev = next
	}
	fmt.Printf("counter: %d \n", counter)

}
