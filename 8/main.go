package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	code, mem, encoded, raw := 0, 0, 0, 0

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		c1, c2 := counts(line)
		code += c1
		mem += c2

		c3, c4 := counts2(line)
		encoded += c3
		raw += c4
	}

	fmt.Println("result part1: ", code-mem)
	fmt.Println("result part2: ", encoded-raw)
}

func counts(quoted string) (int, int) {
	unquoted, err := strconv.Unquote(quoted)
	if err != nil {
		panic(err)
	}
	return len(quoted), len(unquoted)
}

func counts2(unquoted string) (int, int) {
	quoted := strconv.Quote(unquoted)
	return len(quoted), len(unquoted)
}
