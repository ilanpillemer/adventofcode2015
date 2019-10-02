package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	nums := []int{}

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		nums = append(nums, atoi(line))
	}
	fmt.Println(sum(nums) / 3)
	fmt.Println(sum(nums) / 4)
	//	search(sum(nums)/3, nums, []int{}) // part1
	search(sum(nums)/4, nums, []int{}) // part2
	fmt.Println(math.MaxInt64)
	fmt.Println(min)
	fmt.Println(qe)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func sum(ints []int) (s int) {
	for _, v := range ints {
		s += v
	}
	return s
}

func product(ints []int) int {
	p := 1
	for _, v := range ints {
		p = p * v
	}
	if p < 0 { //overflow
		return math.MaxInt64
	}
	return p
}

var min = math.MaxInt64
var qe = math.MaxInt64

func search(goal int, nums []int, ints []int) {

	if sum(ints) == goal {
		if len(ints) <= min {
			if len(ints) < min {
				qe = product(ints)
				fmt.Println(ints)
				fmt.Println(sum(ints))
				fmt.Println(product(ints))
			}
			if product(ints) < qe {
				qe = product(ints)
				fmt.Println(ints)
				fmt.Println(sum(ints))
				fmt.Println(product(ints))
			}
			min = len(ints)

		}

		return
	}
	if len(nums) == 0 {
		return
	}
	next := nums[0]
	nums = nums[1:]
	search(goal, nums, append(ints, next))
	search(goal, nums, ints)

}
