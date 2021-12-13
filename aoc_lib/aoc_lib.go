package aoc_lib

/*
Package containing helper functions used for AOC 2021
*/

import (
	"fmt"
)

/* Checks error and PANIK!!
 */
func check(e error) {
	if e != nil {
		panic(e)
	}
}

/*
Function to iterate through arrays
and print each entry
*/
func printArrInt(int_arr []int) {
	for _, arr := range int_arr {
		fmt.Printf("%d, ", arr)
	}
	fmt.Printf("\n")
}

func printArrStr(str_arr []string) {
	for _, arr := range str_arr {
		fmt.Printf("%s, ", arr)
	}
	fmt.Printf("\n")
}

func intInSlice(des_val int, int_arr []int) bool {
	for _, val := range int_arr {
		if val == des_val {
			return true
		}
	}
	return false
}

func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
