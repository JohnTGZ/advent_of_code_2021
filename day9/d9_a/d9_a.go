package main

import (
	"aoc/aoc"
	"bufio" //For reading line by line
	"flag"  //For command line parsing
	"fmt"
	"image"
	"image/color"
	"image/png"
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
func getInput(filepath string) ([]int, int) {

	f, err := os.Open(filepath)
	aoc.CheckErr(err)
	//close file at end of program
	defer f.Close()
	//create scanner object to read line by line
	scanner := bufio.NewScanner(f)

	var height_map []int
	var map_width int

	init_funcs := true //flag to run some functions only once, like getting the width of the heightmap

	for scanner.Scan() { //scan each line
		text := scanner.Text()
		current_line_height_map := strings.Split(text, "")

		//Functions only run once:
		//get width of map
		if init_funcs {
			map_width = len(current_line_height_map)
			init_funcs = false
		}

		height_map = append(height_map, aoc.ConvArrStrToInt(current_line_height_map)...)
	}

	return height_map, map_width
}

//constant variables
var (
	HIGH_COLOR = [3]float64{255, 0, 0} //red
	LOW_COLOR  = [3]float64{0, 0, 255} //blue
)

func linearGradient(x float64, max_val int) (uint8, uint8, uint8) {
	d := x / float64(max_val)
	r := LOW_COLOR[0] + d*(HIGH_COLOR[0]-LOW_COLOR[0])
	g := LOW_COLOR[1] + d*(HIGH_COLOR[1]-LOW_COLOR[1])
	b := LOW_COLOR[2] + d*(HIGH_COLOR[2]-LOW_COLOR[2])
	return uint8(r), uint8(g), uint8(b)
}

//Generate a visual representation of the height map
func generateHeightMap(height_map []int, map_width int) {
	map_height := (len(height_map) + 1) / map_width

	//create image data struct
	img_topLeft := image.Point{0, 0}
	img_btmRight := image.Point{map_width, map_height}

	img := image.NewRGBA(image.Rectangle{img_topLeft, img_btmRight})

	for i, height_val := range height_map {
		r, g, b := linearGradient(float64(height_val), 9)
		color_gradient := color.RGBA{r, g, b, 0xff}
		x, y := aoc.IdxToXY(i, map_width)
		img.Set(x, y, color_gradient)
	}

	f_int, _ := os.Create("height_map.png")
	png.Encode(f_int, img)
}

func findLowPoints(height_map []int, map_width int) []int {
	map_height := (len(height_map) + 1) / map_width

	//slice of indexes with low points
	var low_points []int

	//Top Left corner: R, B
	//Top Right corner: L, B
	//Btm Left corner: R, U
	//Btm Right corner: L, U

	//Top Left Corner: R, B
	i := 0
	if height_map[i] < height_map[i+1] { //RIGHT
		if height_map[i] < height_map[aoc.XYToIdx(0, 1, map_width)] { //BOTTOM
			low_points = append(low_points, i)
		}
	}

	i = map_width - 1
	//Top Right Corner:  L, B
	if height_map[i] < height_map[i-1] { //LEFT
		if height_map[i] < height_map[aoc.XYToIdx(map_width-1, 1, map_width)] { //BOTTOM
			low_points = append(low_points, i)
		}
	}
	//Bottom Left Corner: R, U
	i = len(height_map) - 1 - (map_width - 1)
	if height_map[i] < height_map[i+1] { //RIGHT
		if height_map[i] < height_map[aoc.XYToIdx(0, map_height-2, map_width)] { //TOP
			low_points = append(low_points, i)
		}
	}
	//Bottom Right Corner: L, U
	i = len(height_map) - 1
	if height_map[i] < height_map[i-1] { //LEFT
		if height_map[i] < height_map[aoc.XYToIdx(map_width-1, map_height-2, map_width)] { //TOP
			low_points = append(low_points, i)
		}
	}
	//Iterate through elements using 1d index
	for i, height_val := range height_map {

		//SKIP CORNERS
		if i == 0 || i == map_width-1 || i == len(height_map)-1-(map_width-1) || i == len(height_map)-1 {
			continue
		}

		//Top Edge: L, R, B
		//Left Edge: R, U, B
		//Right Edge: L, U, B
		//Bottom Edge: L, R, U

		is_low_point := true

		L := false //LEFT
		R := false //RIGHT
		U := false //UP
		B := false //BTM

		if i/map_width < 1 { //handle top edge: L, R, B
			L, R, B = true, true, true
		} else if (i+1)%map_width == 0 { //handle right edge: L, U, B
			L, U, B = true, true, true
		} else if i%map_width == 0 { //handle left edge:  R, U, B
			R, U, B = true, true, true
		} else if i/map_width == (map_height - 1) { //handle bottom layer:  L, R, U
			L, R, U = true, true, true
		} else { //handle everything in between
			L, R, U, B = true, true, true, true
		}

		if L {
			if !(height_val < height_map[i-1]) {
				is_low_point = false
			}
		}
		if R {
			if !(height_val < height_map[i+1]) {
				is_low_point = false
			}
		}
		if U {
			x, y := aoc.IdxToXY(i, map_width)
			if !(height_val < height_map[aoc.XYToIdx(x, y-1, map_width)]) {
				is_low_point = false
			}
		}
		if B {
			x, y := aoc.IdxToXY(i, map_width)
			if !(height_val < height_map[aoc.XYToIdx(x, y+1, map_width)]) {
				is_low_point = false
			}
		}

		if is_low_point {
			low_points = append(low_points, i)
		}

	}

	return low_points
}

func getRiskLevel(height_map []int, low_points []int) int {

	sum := 0
	for _, point_idx := range low_points {
		sum += 1 + height_map[point_idx]
	}
	return sum
}

func main() {

	input_file := procArg()

	height_map, map_width := getInput(input_file)

	generateHeightMap(height_map, map_width)

	low_points := findLowPoints(height_map, map_width)

	fmt.Printf("Risk level: %d", getRiskLevel(height_map, low_points))

	// aoc.PrintArrInt(low_points)

	// fmt.Printf("Printing digital_output_val... \n")
	// aoc.PrintArrStr(digital_output_val)

}
