/*
Package containing helper functions used for AOC 2021
*/
package aoc

import (
	"fmt"
	"math"
	"strconv"
)

/*
Iterate through an array of strings
and print each entry
*/
func PrintArrInt(int_arr []int) {
	fmt.Printf("\n")
	fmt.Printf("Printing int array: \n")
	fmt.Printf("[")
	for _, val := range int_arr[:len(int_arr)-1] {
		fmt.Printf("%d, ", val)
	}
	fmt.Printf("%d]", int_arr[len(int_arr)-1])
	fmt.Printf("\n")
}

/*
Iterate through an array of strings
and print each entry
*/
func PrintArrStr(str_arr []string) {
	fmt.Printf("\n")
	fmt.Printf("Printing string array: \n")
	fmt.Printf("[")
	for _, val := range str_arr[:len(str_arr)-1] {
		fmt.Printf("%s, ", val)
	}
	fmt.Printf("%s]", str_arr[len(str_arr)-1])
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

func GetMedian(int_arr []int) int {
	arr_size := len(int_arr)
	var median int

	if arr_size <= 0 {
		fmt.Printf("Empty array")
		return -1
	}

	if arr_size%2 == 0 {
		median_1 := int_arr[arr_size/2-1]
		median_2 := int_arr[arr_size/2]
		median = (median_1 + median_2) / 2
	} else {
		median = int_arr[int(math.Ceil(float64(arr_size)/2.0))]
	}

	return median
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

/*
*Gets the consecutive sum of the difference between a start and end integer
 */
func ConsecutiveSum(start_val int, end_val int) int {
	consecutive_sum := 0
	difference := Abs(end_val - start_val)

	for i := 1; i < difference+1; i++ {
		consecutive_sum += i
	}

	return consecutive_sum
}

func BubbleSortRunes(runes []rune) string {
	for i := 0; i < len(runes); i++ {
		for j := 0; j < len(runes)-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}

/*
Convert from 2 Dimensional (x,y) to 1D index
*/
func XYToIdx(x int, y int, width int) int {
	return y*width + x
}

/*
Convert from 1D Idx to  2 Dimensional (x,y) index
*/
func IdxToXY(idx int, width int) (int, int) {

	y := int(math.Floor(float64(idx / width)))
	x := idx % width

	return x, y
}
