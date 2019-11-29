package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := 0
	position := 0
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanRunes)
	for in.Scan() {
		c := in.Text()
		switch c {
		case "(":
			count++
		case ")":
			count--
		}
		position ++
		if count == -1 {
fmt.Println("Position ",position)
		}
	}
	fmt.Println(count)
}
