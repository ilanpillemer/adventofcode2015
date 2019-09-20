package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct{ x, y int }

var grid = map[point]bool{}
var bright = map[point]int{}

//turn on 489,959 through 759,964
func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		//fmt.Println(line)
		if strings.HasPrefix(line, "turn on") {
			line = strings.Replace(line, "turn on ", "", -1)
			line = strings.Replace(line, " through ", ",", -1)
			nums := strings.Split(line, ",")
			x1, y1, x2, y2 := atoi(nums[0]), atoi(nums[1]), atoi(nums[2]), atoi(nums[3])
			turnon(point{x1, y1}, point{x2, y2})
		}
		if strings.HasPrefix(line, "turn off") {
			line = strings.Replace(line, "turn off ", "", -1)
			line = strings.Replace(line, " through ", ",", -1)
			nums := strings.Split(line, ",")
			x1, y1, x2, y2 := atoi(nums[0]), atoi(nums[1]), atoi(nums[2]), atoi(nums[3])
			turnoff(point{x1, y1}, point{x2, y2})
		}
		if strings.HasPrefix(line, "toggle") {
			line = strings.Replace(line, "toggle ", "", -1)
			line = strings.Replace(line, " through ", ",", -1)
			nums := strings.Split(line, ",")
			x1, y1, x2, y2 := atoi(nums[0]), atoi(nums[1]), atoi(nums[2]), atoi(nums[3])
			toggle(point{x1, y1}, point{x2, y2})
		}
	}

	count := 0
	for _, v := range grid {
		if v {
			count++
		}
	}
	brightness := 0
	for _, v := range bright {
		brightness += v
	}

	fmt.Println("lights on", count)
	fmt.Println("brightness", brightness)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func turnon(p1, p2 point) {
	if p1.x > p2.x {
		panic("x1>x2")
	}
	if p1.y > p2.y {
		panic("y1>y2")
	}

	for x := p1.x; x < p2.x+1; x++ {
		for y := p1.y; y < p2.y+1; y++ {
			bright[point{x, y}] = bright[point{x, y}] + 1
			grid[point{x, y}] = true
		}
	}
}

func turnoff(p1, p2 point) {
	if p1.x > p2.x {
		panic("x1>x2")
	}
	if p1.y > p2.y {
		panic("y1>y2")
	}

	for x := p1.x; x < p2.x+1; x++ {
		for y := p1.y; y < p2.y+1; y++ {
			bright[point{x, y}] = bright[point{x, y}] - 1
			if bright[point{x, y}] < 0 {
				bright[point{x, y}] = 0
			}
			grid[point{x, y}] = false
		}
	}
}

func toggle(p1, p2 point) {
	if p1.x > p2.x {
		panic("x1>x2")
	}
	if p1.y > p2.y {
		panic("y1>y2")
	}

	for x := p1.x; x < p2.x+1; x++ {
		for y := p1.y; y < p2.y+1; y++ {
			bright[point{x, y}] = bright[point{x, y}] + 2
			grid[point{x, y}] = !grid[point{x, y}]
		}
	}
}
