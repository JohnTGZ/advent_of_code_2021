package main

import (
	"aoc/aoc"
	"bufio" //For reading line by line
	"flag"  //For command line parsing
	"fmt"
	"math"
	"os" //for opening filess
	"sort"
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
func getInput(filepath string) ([]rune, [][]string) {

	f, err := os.Open(filepath)
	aoc.CheckErr(err)
	//close file at end of program
	defer f.Close()
	//create scanner object to read line by line
	scanner := bufio.NewScanner(f)

	//rules_arr is an array containing the rules in the following form
	//rules_arr[n][0] = XX ,
	//rules_arr[n][1] = X
	var rules_arr [][]string

	//First line is the polymer template
	scanner.Scan()
	polymer_template := []rune(scanner.Text())

	//get rid of blank line
	scanner.Scan()

	for scanner.Scan() { //scan each line
		text := scanner.Text()

		// fmt.Printf("current line: %s \n", text)
		current_line := strings.Split(text, " -> ")

		rules := make([]string, 4)
		rules[0] = current_line[0]
		rules[1] = current_line[1]
		// r := make([]rune, 3)
		// r = append(r)
		sub_polymer := []rune(current_line[0])
		added_elem := []rune(current_line[1])
		rules[2] = string([]rune{sub_polymer[0], added_elem[0]})
		rules[3] = string([]rune{added_elem[0], sub_polymer[1]})

		rules_arr = append(rules_arr, rules)
	}

	return polymer_template, rules_arr
}

//Get map of element to it's quantity in the polymer template
func getElemsMap(polymer_template []rune) map[rune]int {
	elem_map := make(map[rune]int)

	for _, element := range polymer_template {
		//if element already exist then simply increment the value
		if _, contains_elem := elem_map[element]; contains_elem {
			elem_map[element] += 1
		} else { //else create the key
			elem_map[element] = 1
		}
	}

	return elem_map
}

//Get map of each subpolymer to it's positions within the polymer template
func getSubPolyMap(polymer_template []rune) map[string][]int {
	sub_poly_map := make(map[string][]int)

	for i, _ := range polymer_template[:len(polymer_template)-1] {

		sub_poly := string(polymer_template[i : i+2])
		sub_poly_map[sub_poly] = append(sub_poly_map[sub_poly], i)
	}

	return sub_poly_map
}

//Insert into a selected index of a rune
func insert(arr []rune, idx int, val rune) []rune {
	if len(arr) == idx {
		return append(arr, val)
	}
	//Add an extra element in the middle
	arr = append(arr[:idx+1], arr[idx:]...)
	arr[idx] = val
	return arr

}

func main() {

	input_file := procArg()

	polymer_template, rules_arr := getInput(input_file)

	fmt.Printf("Polymer Template at start: %s \n", string(polymer_template))

	///////////////////////////////////////////
	//Implementation 1
	///////////////////////////////////////////
	//Naive implementation (10 * n * m)
	// for i := 0; i < 10; i++ {
	// 	//Iterate through the rules
	// 	for _, rules := range rules_arr {
	// 		for i, _ := range polymer_template[:len(polymer_template)-1] {
	// 			if string([]rune{polymer_template[i], polymer_template[i+1]}) == rules[0] {
	// 				polymer_template = insert(polymer_template, i, []rune(rules[1])[0])
	// 			}
	// 		}
	// 	}
	// }

	///////////////////////////////////////////
	//Implementation 2: Keep a mapping that can be looked up when adding to polymers
	///////////////////////////////////////////

	//Keep a map of sub-polymers, which map each sub-polymer to an array of positions
	//	e.g. for the polymer NNCNN: subPolyMap["NN"] = [0, 3], subPolyMap["NC"] = [1], subPolyMap["CN"] = [2]

	//Create map of where each sub-polymer
	subPolyMap := getSubPolyMap(polymer_template)

	for sub_poly_string, sub_poly := range subPolyMap {
		fmt.Printf("%s: ", sub_poly_string)
		for _, sub_poly_pos := range sub_poly {
			fmt.Printf("%d, ", sub_poly_pos)
		}
		fmt.Printf("\n")
	}

	// subPolyMap_new := make(map[string][]int)

	//copy over map
	// for k, v := range subPolyMap {
	// 	subPolyMap_new[k] = v
	// }

	//Keep mapping of what pairs there are
	for i := 0; i < 40; i++ {
		element_to_pos_map := make(map[int]rune)

		subPolyMap_new := getSubPolyMap(polymer_template)

		//Iterate through the rules
		for _, rules := range rules_arr {
			//rules[0]: Original polymer
			//rules[1]: Element to be added
			//rules[2]: First 2 elemnts of new subpolymer
			//rules[3]: Last 2 elemnts of new subpolymer

			//If subpolymer matches the current rule
			if position_arr, ok := subPolyMap_new[rules[0]]; ok {
				fmt.Printf("Current rule: %s -> %s \n", rules[0], rules[1])
				//add elements to all sub polymers in position_arr
				fmt.Printf("	Positions: ")
				for _, pos := range position_arr {
					fmt.Printf("%d,", pos)
					//Offset is used to make sure that subsequently inserted positions are correct
					// polymer_template = insert(polymer_template, pos+1, []rune(rules[1])[0])
					element_to_pos_map[pos+1] = []rune(rules[1])[0]

					// //remove existing mapping
					// subPolyMap_new[rules[0]] = aoc.FindAndDelete(subPolyMap_new[rules[0]], pos)

					// //add new mappings
					// subPolyMap_new[rules[2]] = append(subPolyMap_new[rules[2]], pos)
					// subPolyMap_new[rules[3]] = append(subPolyMap_new[rules[3]], pos+1)

				}
				fmt.Printf("\n")
				fmt.Printf("Polymer template: %s,", string(polymer_template))
				fmt.Printf("\n")
			}
		}

		//Sort the inserted positions
		insert_pos_arr := make([]int, 0, len(element_to_pos_map))

		for key := range element_to_pos_map {
			insert_pos_arr = append(insert_pos_arr, key)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(insert_pos_arr)))

		//Insert elements into polymer_template from highest position index to lowest
		for _, key := range insert_pos_arr {

			polymer_template = insert(polymer_template, key, element_to_pos_map[key])
		}

	}

	fmt.Printf("\n")

	fmt.Printf("Polymer Template at end: %s \n", string(polymer_template))

	fmt.Printf("\n")

	//Get mapping of elements to their quantity
	elem_map := getElemsMap(polymer_template)

	var most_common_elem_key, least_common_elem_key rune
	lowest_num := math.MaxInt32
	highest_num := 0
	for element, quantity := range elem_map {
		// fmt.Printf("Elem: %c, Quantity: %d \n", element, quantity)
		if quantity <= lowest_num {
			lowest_num = quantity
			// fmt.Printf("lowest num: %d \n", quantity)
			least_common_elem_key = element
		}
		if quantity >= highest_num {
			highest_num = quantity
			// fmt.Printf("highest num: %d \n", quantity)
			most_common_elem_key = element
		}
	}

	fmt.Printf("\n")

	fmt.Printf("Length of polymer template at end: %d \n", len(polymer_template))
	//Answer: Take quantity of most common element
	//and subtract the quantity of the least common element
	fmt.Printf("No. of %c (%d) - no. of %c (%d) =  %d \n",
		most_common_elem_key, elem_map[most_common_elem_key],
		least_common_elem_key, elem_map[least_common_elem_key],
		elem_map[most_common_elem_key]-elem_map[least_common_elem_key])

}
