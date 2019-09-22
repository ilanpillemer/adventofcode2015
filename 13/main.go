package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//0     1     2    3 4          5     6   7       8   9 10
//Alice would lose 2 happiness units by sitting next to Bob.

var nodes = map[string]bool{}
var edges = []edge{}
var max = 0

type edge struct {
	from string
	to   string
	val  int
}

func main() {

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		fields := strings.Fields(line)
		nodes[fields[0]] = true
		e := edge{
			from: fields[0],
			to:   strings.TrimSuffix(fields[10], "."),
			val:  getVal(fields[2], fields[3]),
		}
		edges = append(edges, e)
	}

	for n := range nodes {
		e := edge{
			from: "me",
			to:   n,
		}
		edges = append(edges, e)
		e2 := edge{
			to:   "me",
			from: n,
		}
		edges = append(edges, e2)
	}
	nodes["me"] = true
	search(nodes, map[string]bool{}, []string{}, len(nodes))
	fmt.Println("Max Happiness", max)
}

func getVal(t string, a string) int {
	if t == "lose" {
		return -(atoi(a))
	}
	if t == "gain" {
		return atoi(a)
	}

	panic("unexpected: " + t + " " + a)

}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func search(n map[string]bool, seen map[string]bool, arrange []string, guests int) {
	if len(arrange) == guests {
		s := score(arrange)
		if s > max {
			max = s
		}
		return
	}
	for k, _ := range n {
		if seen[k] {
			continue
		}
		seen[k] = true
		search(n, seen, append(arrange, k), guests)
		delete(seen, k)
	}
}

func score(arrange []string) int {
	total := 0
	arrange = append(arrange, arrange[0])
	for i := 0; i < len(arrange)-1; i++ {
		total += happiness(arrange[i], arrange[i+1])
	}
	return total
}

func happiness(g1 string, g2 string) int {
	h := 0
	for _, e := range edges {
		if e.to == g1 && e.from == g2 {
			h += e.val
		}
		if e.from == g1 && e.to == g2 {
			h += e.val
		}
	}
	return h
}
