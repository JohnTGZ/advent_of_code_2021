package aoc

/*
Package containing helper functions used for AOC 2021
*/

import (
	"fmt"
	"strconv"
)

/*
Iterate through an array of strings
and print each entry
*/
func PrintArrInt(int_arr []int) {
	fmt.Printf("\n")
	for _, val := range int_arr {
		fmt.Printf("%d, ", val)
	}
	fmt.Printf("\n")
}

/*
Iterate through an array of strings
and print each entry
*/
func PrintArrStr(str_arr []string) {
	fmt.Printf("\n")
	for _, val := range str_arr {
		fmt.Printf("%s, ", val)
	}
	fmt.Printf("\n")
}

/*
Iterate through an array of strings
and print each entry
*/
func PrintIntToIntMap(int_map map[int]int) {
	fmt.Printf("\n")
	for key, val := range int_map {
		fmt.Printf("key: %d, val: %d, \n", key, val)
	}
	fmt.Printf("\n")
}

/*
Convert from array of strings to array of ints
*/
func ConvArrStrToInt(str_arr []string) []int {
	var int_arr []int
	for _, val := range str_arr {
		val_int, err := strconv.Atoi(val)
		int_arr = append(int_arr, val_int)
		CheckErr(err)
	}
	return int_arr
}

/*
Checks if integer resides in given slice
*/
func IntInSlice(des_val int, int_arr []int) bool {
	for _, val := range int_arr {
		if val == des_val {
			return true
		}
	}
	return false
}

/*
Get the sum of all elements in an array
*/
func GetSum(int_arr []int) int {
	sum := 0
	for _, val := range int_arr {
		sum += val
	}
	return sum
}

/*
Returns an absolute integer
*/
func Abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

/* Checks error and PANIK!!
 */
func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}
