package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sizes = []int{}
var ways = map[int][][]int{}

var total = 0

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		s := strings.TrimSpace(in.Text())
		if s != "" {
			sizes = append(sizes, atoi(s))
		}
	}
	fmt.Println(sizes)
	search([]int{}, 0)
	fmt.Println(total)
	for k, v := range ways {
		fmt.Println(k, len(v))
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func search(option []int, i int) {
	if sum(option) == 150 {
		total++
		opt, ok := ways[len(option)]
		if !ok {
			opt = [][]int{}
		}
		opt = append(opt, option)
		ways[len(option)] = opt
		return
	}
	if i == len(sizes) {
		return
	}
	option = append(option, sizes[i])
	search(option, i+1)
	option = option[:len(option)-1]
	search(option, i+1)
}

func sum(o []int) int {
	total := 0
	for _, v := range o {
		total += v
	}
	return total
}
