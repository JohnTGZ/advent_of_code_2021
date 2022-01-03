package main

import (
	"bufio" //For reading line by line
	"flag"  //For command line parsing
	"fmt"   //Converts string into integers
	"math"
	"os" //for opening filess
	"strconv"
	"strings"
	//Converts string into integers
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

/* Checks error and PANIK!!
 */
func check(e error) {
	if e != nil {
		panic(e)
	}
}

/* Read the input into an iterable array
 */
func getInput(filepath string) ([][]int, int, int) {

	f, err := os.Open(filepath)
	check(err)

	//close file at end of program
	defer f.Close()

	//create scanner object to read line by line
	scanner := bufio.NewScanner(f)

	var coordinate_arr [][]int

	max_width := 0
	max_height := 0

	//read file line by line
	for scanner.Scan() {
		text := scanner.Text()

		coordinates := strings.Split(text, " -> ")
		xy1 := strings.Split(coordinates[0], ",")
		xy2 := strings.Split(coordinates[1], ",")

		x1, _ := strconv.Atoi(xy1[0])
		y1, _ := strconv.Atoi(xy1[1])
		x2, _ := strconv.Atoi(xy2[0])
		y2, _ := strconv.Atoi(xy2[1])

		//Get max width and height
		if x1 > max_width {
			max_width = x1
		} else if x2 > max_width {
			max_width = x2
		}
		if y1 > max_height {
			max_height = y1
		} else if y2 > max_height {
			max_height = y2
		}

		//calculate difference
		dx, dy := getDirectionVect(x1, y1, x2, y2)

		xy_pos := []int{x1, y1, x2, y2, dx, dy}

		if Abs(dx) > 0 && Abs(dy) > 0 {
			fmt.Printf("Not diagonal: %d,%d -> %d, %d \n", x1, y1, x2, y2)
		} else {
			coordinate_arr = append(coordinate_arr, xy_pos)
		}
	}

	return coordinate_arr, max_width + 1, max_height + 1
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

/*
Convert from 2 Dimensional (x,y) to 1D index
*/
func get1DIdx(x int, y int, width int) int {
	return y*width + x
}

/*
Convert from 1D Idx to  2 Dimensional (x,y) index
*/
func get2DIdx(idx int, width int) (int, int) {

	y := int(math.Floor(float64(idx / width)))
	x := idx % width

	return x, y
}

/*
Get direction vector
*/
func getDirectionVect(x1 int, y1 int, x2 int, y2 int) (int, int) {
	dx := x2 - x1
	dy := y2 - y1

	return dx, dy
}

/*
Fills up the vent arr given the coordinate positions
*/
func lineVentUp(vents_arr []int, coords []int, max_width int) {
	//get direction vector
	dx, dy := coords[4], coords[5]
	x, y := coords[0], coords[1]

	idx := get1DIdx(x, y, max_width)
	vents_arr[idx]++

	for {

		//stepping through
		if dy == 0 {
			if dx > 0 {
				x++

			} else {
				x--
			}

		} else {
			if dy > 0 {
				y++
			} else {
				y--
			}
		}

		//fill up array
		idx := get1DIdx(x, y, max_width)
		vents_arr[idx]++

		if x == coords[2] && y == coords[3] {
			break
		}
	}

}

func printVentMap(vents_arr []int, max_width int) {
	fmt.Printf("===Printing vent===")

	for i, val := range vents_arr {
		if i%max_width == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%d ", val)
	}
	fmt.Printf("\n")
}

func getTotalDangerousPoints(vents_arr []int, threshold int) int {
	sum := 0
	for _, val := range vents_arr {
		if val >= threshold {
			sum++
		}
	}

	return sum
}

func main() {

	input_file := procArg()

	coordinate_arr, max_width, max_height := getInput(input_file)

	//create 1d array of vents
	vents_arr := make([]int, max_width*max_height)

	//iterate through each coord (x1, y1, x2, y2, dx, dy)
	for _, coord := range coordinate_arr {
		lineVentUp(vents_arr, coord, max_width)
	}

	printVentMap(vents_arr, max_width)

	fmt.Printf("Total dangerous points: %d \n", getTotalDangerousPoints(vents_arr, 2))

}
