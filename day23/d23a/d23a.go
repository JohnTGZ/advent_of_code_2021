package main

import (
	//For reading line by line
	"aoc/aoc"
	"bufio"
	"flag" //For command line parsing
	"fmt"
	"os"
	"strings"
)

//constant variables
const (
	A_COST              = 1
	B_COST              = 10
	C_COST              = 100
	D_COST              = 1000
	PROHIBITED_POS      = []int{2, 4, 6, 8}           //Prohibited position
	ALLOWED_HALLWAY_POS = []int{0, 1, 3, 5, 7, 9, 10} //Allowed positions in hallway

	A_ALLOWED_POS = []int{0, 1, 3, 5, 7, 9, 10, 11, 15} //Allowed positions to move to
	B_ALLOWED_POS = []int{0, 1, 3, 5, 7, 9, 10, 12, 16} //Allowed positions to move to
	C_ALLOWED_POS = []int{0, 1, 3, 5, 7, 9, 10, 13, 17} //Allowed positions to move to
	D_ALLOWED_POS = []int{0, 1, 3, 5, 7, 9, 10, 14, 18} //Allowed positions to move to
)

var _expected_test_answer int = 0
var is_testing_ptr *bool

type Amphipod struct {
	species string //Type of amphipod: "A", "B", "C" or "D"

	total_cost int //Total accumulated movement cost
	move_cost  int //Cost of movement

	current_pos   int   //Initial position of amphipod
	check_pos_arr []int //Array to check if desired pose to move to is possible
}

//Constructor for amphipod
func NewAmphipod(species string, initial_pos int) Amphipod {
	amphipod := Amphipod{}
	amphipod.species = species
	amphipod.current_pos = initial_pos
	amphipod.total_cost = 0

	switch species {
	case "A":
		amphipod.move_cost = A_COST
		amphipod.check_pos_arr = A_ALLOWED_POS
	case "B":
		amphipod.move_cost = B_COST
		amphipod.check_pos_arr = B_ALLOWED_POS
	case "C":
		amphipod.move_cost = C_COST
		amphipod.check_pos_arr = C_ALLOWED_POS
	case "D":
		amphipod.move_cost = D_COST
		amphipod.check_pos_arr = D_ALLOWED_POS
	}
	return amphipod
}

//Move Amphipod to desired position
func (amp Amphipod) move(des_pos int) bool {

	for _, check_pos := range check_pos_arr {
		if des_pos == check_pos {
			//Successful move
			amp.current_pos = des_pos
			amp.total_cost += amp.move_cost
			return true
		}
	}
	return false
}

/* Process command line arguemnts to read from either
test input or normal input file
*/
func procArg() string {
	is_testing_ptr = flag.Bool("t", false, "Enable testing")

	flag.Parse()

	input_file := "../input.txt"
	if *is_testing_ptr {
		input_file = "../test_input.txt"
		_expected_test_answer = 12521
	}

	return input_file
}

/* Read the input into an iterable array
 */
func getInput(filepath string) ([]string, []string) {

	f, err := os.Open(filepath)
	aoc.CheckErr(err)
	//close file at end of program
	defer f.Close()
	//create scanner object to read line by line
	scanner := bufio.NewScanner(f)

	line_num := 1

	var mid_arr, base_arr []string
	for scanner.Scan() { //scan each line
		text := scanner.Text()

		current_line_arr := strings.Split(text, "")
		// fmt.Printf("%s \n", current_line_arr)

		if line_num == 3 {
			for _, letter := range current_line_arr {
				if letter != "#" && letter != " " {
					mid_arr = append(mid_arr, letter)
				}

			}
		}
		if line_num == 4 {
			for _, letter := range current_line_arr {
				if letter != "#" && letter != " " {
					base_arr = append(base_arr, letter)
				}
			}
		}

		line_num++
	}

	return base_arr, mid_arr
}

/* Sets up the map from base and mid arrays
 */
func setUpMap(base_arr_str []string, mid_arr_str []string) []int {
	map_int_arr := make([]int, 19)

	var base_arr_int [4]int
	var mid_arr_int [4]int
	for i, elem := range base_arr_str {
		var elem_int int
		switch elem {
		case ".":
			elem_int = 0
		case "A":
			elem_int = 1
		case "B":
			elem_int = 2
		case "C":
			elem_int = 3
		case "D":
			elem_int = 4
		}
		base_arr_int[i] = elem_int
	}

	for i, elem := range mid_arr_str {
		var elem_int int
		switch elem {
		case ".":
			elem_int = 0
		case "A":
			elem_int = 1
		case "B":
			elem_int = 2
		case "C":
			elem_int = 3
		case "D":
			elem_int = 4
		}
		mid_arr_int[i] = elem_int
	}

	for i, elem := range mid_arr_int {
		map_int_arr[i+11] = elem
	}

	for i, elem := range base_arr_int {
		map_int_arr[i+15] = elem
	}

	return map_int_arr
}

/* Prints current state of the map to indicate which positions are correct
 */
func printMapOrg(map_int_arr []bool) {

	//First line
	fmt.Printf("#############\n")

	//Second line
	fmt.Printf("#")
	fmt.Printf("...........")
	fmt.Printf("#\n")

	//Third line
	fmt.Printf("###")
	for _, letter := range map_int_arr[11:15] {
		fmt.Printf("%t#", letter)
	}
	fmt.Printf("##\n")

	//Fourth line
	fmt.Printf("  #")
	for _, letter := range map_int_arr[15:19] {
		fmt.Printf("%t#", letter)
	}
	fmt.Printf("\n")

	//Fifth line
	fmt.Printf("  #########\n")

}

/* Prints current state of the map
 */
func printMap(iteration int, map_int_arr []int) {
	fmt.Printf("\n")

	fmt.Printf("Iteration %d \n", iteration)

	var map_str_arr [19]string //String representation of map
	for i, elem := range map_int_arr {
		var elem_str string
		switch elem {
		case 0:
			elem_str = "."
		case 1:
			elem_str = "A"
		case 2:
			elem_str = "B"
		case 3:
			elem_str = "C"
		case 4:
			elem_str = "D"
		}
		map_str_arr[i] = elem_str
	}

	//First line
	fmt.Printf("#############\n")

	//Second line
	fmt.Printf("#")
	for _, letter := range map_str_arr[:11] {
		fmt.Printf("%s", letter)
	}
	fmt.Printf("#\n")

	//Third line
	fmt.Printf("###")
	for _, letter := range map_str_arr[11:15] {
		fmt.Printf("%s#", letter)
	}
	fmt.Printf("##\n")

	//Fourth line
	fmt.Printf("  #")
	for _, letter := range map_str_arr[15:19] {
		fmt.Printf("%s#", letter)
	}
	fmt.Printf("\n")

	//Fifth line
	fmt.Printf("  #########\n")

}

/* Checks the base to see which positions are currently out of order
 */
func checkMap(map_int_arr []int, map_org_int []bool) bool {
	organization_complete := true
	//Check Base Rooms
	for i, room := range map_int_arr[15:19] {
		fmt.Printf("Room(%d) == i+1(%d) \n ", room, i+1)
		if room == i+1 { //If room matches amphipod
			map_org_int[i+15] = true
		}
	}
	//Check Mid Rooms
	for i, room := range map_int_arr[11:15] {
		fmt.Printf("Room(%d) == i+1(%d) \n", room, i+1)
		if room == i+1 { //If room matches amphipod
			map_org_int[i+11] = true
		}
	}

	for i := 0; i < 4; i++ {
		if !map_org_int[i+11] || !map_org_int[i+15] {
			organization_complete = false
		}
	}
	return organization_complete
}

/*
Iterate through the solution
*/
func iterate(map_int_arr []int)

func main() {

	input_file := procArg()
	fmt.Printf("Input File Name: %s \n", input_file)
	base_arr, mid_arr := getInput(input_file)

	map_org_int := make([]bool, 19) //map_org_int indicates which positions are already organized correctly

	map_int_arr := setUpMap(base_arr, mid_arr) //map_int_arr indicates which amphipods are in which position
	organization_complete := checkMap(map_int_arr, map_org_int)

	printMap(0, map_int_arr)
	// printMapOrg(map_org_int)

	i := 0

	for !organization_complete {
		iterate(map_int_arr)
		printMap(i, map_int_arr)
		organization_complete = checkMap(map_int_arr, map_org_int)
		i++
	}

	final_answer := 0

	if *is_testing_ptr {
		if final_answer != _expected_test_answer {
			fmt.Printf("FAILURE: Final answer %d did not match expected answer %d \n", final_answer, _expected_test_answer)
		} else {
			fmt.Printf("SUCCESS: Final answer %d matched expected answer %d! \n", final_answer, _expected_test_answer)
		}
	}

}
