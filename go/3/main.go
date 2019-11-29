package main

import (
	"bufio"
	"fmt"
	"os"
)

var loc = map[point]bool{}

type point struct {
	x, y int
}

func main() {
	x := 0
	y := 0
	x1 := 0
	y1 := 0

	loc[point{0, 0}] = true
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanRunes)
	alternate := true
	for in.Scan() {
		c := in.Text()
		switch c {
		case ">":
			if alternate {
				x++
			} else {
				x1++
			}
		case "<":
			if alternate {
				x--
			} else {
				x1--
			}
		case "v":
			if alternate {
				y++
			} else {
				y1++
			}
		case "^":
			if alternate {
				y--
			} else {
				y1--
			}
		default:
			panic(c)
		}
		loc[point{x, y}] = true
		loc[point{x1, y1}] = true
		alternate = !alternate
	}
	fmt.Println(len(loc))
}

func main1() {
	x := 0
	y := 0

	loc[point{0, 0}] = true
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanRunes)
	for in.Scan() {
		c := in.Text()
		switch c {
		case ">":
			x++
		case "<":
			x--
		case "v":
			y++
		case "^":
			y--
		default:
			panic(c)
		}
		loc[point{x, y}] = true
	}
	fmt.Println(len(loc))
}
