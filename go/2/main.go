package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	total := 0
	length := 0
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		vals := strings.Split(line, "x")
		l, w, h := atoi(vals[0]), atoi(vals[1]), atoi(vals[2])
		paper := paper(l, w, h)
		total += paper
		length += ribbon(l, w, h)
	}

	fmt.Println("wrapping paper", total)
	fmt.Println("ribbon", length)
}

func area(l, w, h int) int {
	return 2*l*w + 2*w*h + 2*h*l
}

func min(l, w, h int) int {
	a := l * w
	b := w * h
	c := l * h

	temp := a
	if b < a {
		temp = b
	}

	if temp < c {
		return temp
	}

	return c
}

func smallesttwo(a, b, c int) (int, int) {

	if a >= b && a >= c {
		return b, c
	}

	if b >= a && b >= c {
		return a, c
	}

	return b, a

}

func paper(l, w, h int) int {
	return area(l, w, h) + min(l, w, h)
}

func ribbon(l, w, h int) int {
	a, b := smallesttwo(l, w, h)
	return (a + a + b + b) + (l * w * h)

}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
