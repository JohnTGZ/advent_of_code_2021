package main

import (
	//For reading line by line
	"flag" //For command line parsing
	"fmt"
	"math"
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

//constant variables
const (
	MAX_RANGE  = 1000
	DEG_TO_RAD = math.Pi / 180
	RAD_TO_DEG = 180 / math.Pi
)

var (
	fixed_ori_val = []float64{0, 90 * DEG_TO_RAD, 180 * DEG_TO_RAD, 270 * DEG_TO_RAD}
)

type scanner struct {
	rot_mat            [3][3]int //rotation matrix
	beacon_coordinates [][]int   //beacon coordinates
}

/* Read the input into an iterable array
 */
// func getInput(filepath string) ([]int, int) {

// 	f, err := os.Open(filepath)
// 	aoc.CheckErr(err)
// 	//close file at end of program
// 	defer f.Close()
// 	//create scanner object to read line by line
// 	scanner := bufio.NewScanner(f)

// 	for scanner.Scan() { //scan each line
// 		text := scanner.Text()

// 		if strings.Contains(text, "scanner") { //Start of new scanner
// 			continue
// 		}
// 		// current_line_height_map := strings.Split(text, ",")
// 	}

// 	return height_map, map_width
// }

// func main() {

// 	input_file := procArg()

// 	// height_map, map_width := getInput(input_file)

// 	// aoc.PrintArrInt(low_points)

// 	// fmt.Printf("Printing digital_output_val... \n")
// 	// aoc.PrintArrStr(digital_output_val)

// }

func main() {
	rot_mat := [3][3]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}

	rot_mat = [3][3]int{
		{1, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}
	//About X axis
	fmt.Printf("About X axis \n")
	for _, val := range fixed_ori_val {
		rot_mat[1][1] = int(math.Cos(val))
		rot_mat[1][2] = int(-math.Sin(val))
		rot_mat[2][1] = int(math.Sin(val))
		rot_mat[2][2] = int(math.Cos(val))

		fmt.Printf("Rotation: %f: \n", val)
		fmt.Print("  ", rot_mat, "\n")
	}

	rot_mat = [3][3]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	fmt.Printf("About Y axis \n")
	//About Y axis
	for _, val := range fixed_ori_val {
		rot_mat[0][0] = int(math.Cos(val))
		rot_mat[0][2] = int(math.Sin(val))
		rot_mat[2][0] = int(-math.Sin(val))
		rot_mat[2][2] = int(math.Cos(val))

		fmt.Printf("Rotation: %f: \n", val)
		fmt.Print("  ", rot_mat, "\n")
	}

	rot_mat = [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 1},
	}
	fmt.Printf("About Z axis \n")
	//About Z axis
	for _, val := range fixed_ori_val {
		rot_mat[0][0] = int(math.Cos(val))
		rot_mat[0][1] = int(-math.Sin(val))
		rot_mat[1][0] = int(math.Sin(val))
		rot_mat[1][1] = int(math.Cos(val))

		fmt.Printf("Rotation: %f: \n", val)
		fmt.Print("  ", rot_mat, "\n")
	}

}
