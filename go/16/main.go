package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Sue 1: cars: 9, akitas: 3, goldfish: 0
//Sue 2: akitas: 9, children: 3, samoyeds: 9
//Sue 3: trees: 6, cars: 6, children: 4
//Sue 4: trees: 4, vizslas: 4, goldfish: 9
//Sue 5: akitas: 9, vizslas: 7, cars: 5

var sues = map[string]map[string]int{}

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		i := strings.Index(line, ":")
		sue := strings.Fields(line[:i])[1]
		items := strings.Split(line[i+1:], ",")
		dict := map[string]int{}
		for _, j := range items {
			f := strings.Split(j, ":")
			dict[strings.TrimSpace(f[0])] = atoi(f[1])
		}
		sues[sue] = dict
	}

	//children: 3
	//cats: 7
	//samoyeds: 2
	//pomeranians: 3
	//akitas: 0
	//vizslas: 0
	//goldfish: 5
	//trees: 3
	//cars: 2
	//perfumes: 1

	for sue, dict := range sues {
		for item, value := range dict {
			if !check(item, value) {
				delete(sues, sue)
			}
		}
	}
	fmt.Println("and...", sues)

}

//In particular, the cats and trees readings indicates that there are
//greater than that many (due to the unpredictable nuclear decay of
//cat dander and tree pollen), while the pomeranians and goldfish
//readings indicate that there are fewer than that many (due to the
//modial interaction of magnetoreluctance).

func check(s string, val int) bool {
	switch s {
	case "children":
		return val == 3
	case "cats":
		return val > 7
	case "samoyeds":
		return val == 2
	case "pomeranians":
		return val < 3
	case "akitas":
		return val == 0
	case "vizslas":
		return val == 0
	case "goldfish":
		return val < 5
	case "trees":
		return val > 3
	case "cars":
		return val == 2
	case "perfumes":
		return val == 1
	}
	return true
}

func atoi(s string) int {
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
