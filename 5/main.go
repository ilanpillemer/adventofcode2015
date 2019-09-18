package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1count := 0
	part2count := 0
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		if isNice2(line) {
			part2count++
		}

		if isNice(line) {
		part1count++
		}
	}
	fmt.Println("part1", part1count)
	fmt.Println("part2", part2count)

}

func isNice(str string) bool {
	vowels := 0
	double := false
	for i, v := range str {
		switch v {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		}
		if i != 0 {
			if str[i] == str[i-1] {
				double = true
			}

			icky := str[i-1 : i+1]
			switch icky {
			case "ab", "cd", "pq", "xy":
				return false
			}
		}
	}
	if vowels >= 3 && double {
		return true
	}
	return false
}

func isNice2(str string) bool {

	pairs := pairs(str)
	hasDupe := dupe(pairs)
	triples := triples(str)
	isCirci := circ(triples)
	return hasDupe && isCirci
}

func pairs(str string) map[string]string {
	p := map[string]string{}
	for i := 0; i < len(str); i++ {
		if i != 0 {
			p[str[i-1:i+1]] = str[:i-1] + "__" + str[i+1:]
		}
	}
	return p
}

func triples(str string) map[string]bool {
	t := map[string]bool{}
	for i := 0; i < len(str); i++ {
		if i > 1 {
			t[str[i-2:i+1]] = true
		}
	}
	return t
}

func circ(triples map[string]bool) bool {
	for k, _ := range triples {
		if k[0] == k[2] {
			return true
		}
	}
	return false
}

func dupe(pairs map[string]string) bool {
	for k, v := range pairs {
		if strings.Contains(v, k) {
			return true
		}
	}
	return false
}
