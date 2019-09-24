package main

import (
	"bufio"
	"fmt"
	"os"
)

type pos struct{ x, y int }

var grid = map[pos]bool{}

func main() {
	in := bufio.NewScanner(os.Stdin)
	y := 0
	for in.Scan() {
		line := in.Text()

		for x, v := range line {
			switch v {
			case '#':
				grid[pos{x, y}] = true
			case '.':
				grid[pos{x, y}] = false
			}
		}
		y++
	}

	grid[pos{0, 0}] = true
	grid[pos{99, 99}] = true
	grid[pos{0, 99}] = true
	grid[pos{99, 0}] = true

	for i := 0; i < 100; i++ {
		grid = iterate()
		//display()
	}

	display()
	fmt.Println("Number On", numOn())
}

func display() {
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			switch grid[pos{x, y}] {
			case true:
				fmt.Print("#")
			case false:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func iterate() map[pos]bool {
	var next = map[pos]bool{}
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			next[pos{x, y}] = value(x, y)
		}
	}
	next[pos{0, 0}] = true
	next[pos{99, 99}] = true
	next[pos{0, 99}] = true
	next[pos{99, 0}] = true
	return next
}

//A light which is on stays on when 2 or 3 neighbors are on, and turns off otherwise.
//A light which is off turns on if exactly 3 neighbors are on, and stays off otherwise.
func value(x, y int) bool {
	isOn := grid[pos{x, y}]
	neigh := onNeigh(x, y)
	if isOn {
		if neigh == 2 || neigh == 3 {
			return true
		}
		return false
	}
	if neigh == 3 {
		return true
	}
	return false
}

func onNeigh(x, y int) int {
	count := 0
	if grid[pos{x - 1, y + 1}] {
		count++
	}
	if grid[pos{x - 1, y}] {
		count++
	}
	if grid[pos{x - 1, y - 1}] {
		count++
	}
	if grid[pos{x, y + 1}] {
		count++
	}

	if grid[pos{x, y - 1}] {
		count++
	}
	if grid[pos{x + 1, y + 1}] {
		count++
	}
	if grid[pos{x + 1, y}] {
		count++
	}
	if grid[pos{x + 1, y - 1}] {
		count++
	}
	return count
}

func numOn() int {
	c := 0
	for _, v := range grid {
		if v {
			c++
		}
	}
	return c
}
