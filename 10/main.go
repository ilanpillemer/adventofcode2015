package main

import "fmt"

//1 becomes 11 (1 copy of digit 1).
//11 becomes 21 (2 copies of digit 1).
//21 becomes 1211 (one 2 followed by one 1).
//1211 becomes 111221 (one 1, one 2, and two 1s).
//111221 becomes 312211 (three 1s, two 2s, and one 1).

func main() {
	input := "1321131112"
	prevlen := 1.0
	for i := 0; i < 50; i++ {
		input = transform(input)
		fmt.Println(len(input), float64(len(input))/prevlen)
		prevlen = float64(len(input))
	}
	fmt.Println("finished")
	fmt.Println(len(input), float64(len(input))/prevlen)
}

func transform(input string) string {
	prev := input[0]
	count := 0
	out := ""
	for i := 1; i < len(input); i++ {
		if input[i] == prev {
			count++
		}

		if input[i] != prev {
			count++
			out = fmt.Sprintf("%s%d%c", out, count, prev)
			count = 0
		}

		prev = input[i]
	}

	// last one
	count++
	out = fmt.Sprintf("%s%d%c", out, count, prev)
	return out
}
