package main

import (
	"fmt"
	"os"
)

type pos struct{ x, y int }

var grid = map[pos]int{}
var lookup = map[int]pos{}
var codes = map[pos]int{}

//To continue, please consult the code grid in the manual.  Enter the code at row 2947, column 3029.

func main() {
	x := 0
	y := 0
	maxy := 0
	maxx := 0
	codes[pos{0, 0}] = 20151125
	for i := 1; i < 100000000; i++ {
		grid[pos{x, y}] = i
		lookup[i] = pos{x, y}
		if !(x == 0 && y == 0) {
			previous := lookup[i-1]
			previousCode := codes[previous]
			newCode := (previousCode * 252533) % 33554393
			codes[pos{x, y}] = newCode
			if x == 3029-1 && y == 2947-1 {
				fmt.Println("Got you!", newCode)
				os.Exit(0)
			}
		}
		if y == 0 {
			x = 0
			maxy++
			y = maxy
			continue
		}
		y--
		x++
		if x > maxx {
			maxx = x
		}
	}
}

func display() {
	for y := 0; y < 7; y++ {
		for x := 0; x < 7; x++ {
			fmt.Print(grid[pos{x, y}], " ")
		}
		fmt.Println()
	}
}

func displayCodes() {
	for y := 0; y < 7; y++ {
		for x := 0; x < 7; x++ {
			fmt.Print(codes[pos{x, y}], " ")
		}
		fmt.Println()
	}
}
