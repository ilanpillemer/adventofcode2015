package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanRunes)
	total := 0
	inNumber := false
	current := ""
	for in.Scan() {
		c, _ := utf8.DecodeRuneInString(in.Text())
		switch {
		case c == '-':
			current = "-"
			inNumber = true
		case unicode.IsNumber(c):
			current += string(c)
			inNumber = true
		case inNumber:
			i := atoi(current)
			total += i
			current = ""
			inNumber = false
		}
	}
	fmt.Println("total", total)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
