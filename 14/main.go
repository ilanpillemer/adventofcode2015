package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//   0    1   2   3  4    5   6   7       8   9    10    11  12  13  14
// Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds
func main() {
part1()
}

func part1() {

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		f := strings.Fields(line)
		r := rate{
			km:   atoi(f[3]),
			s:    atoi(f[6]),
			rest: atoi(f[13]),
		}
		fmt.Println(distanceAt(2503, r))
	}

}

type rate struct {
	km   int
	s    int
	rest int
}

func distanceAt(seconds int, r rate) int {
	runs, rem := seconds/(r.rest+r.s), seconds%(r.rest+r.s)
	d := runs * r.km * r.s
	left := min(rem, r.s) * r.km
	return d + left
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
