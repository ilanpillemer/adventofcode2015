package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type node string
type edge struct {
	from node
	to   node
	d    int
}

var nodes = map[node]bool{}
var edges = []edge{}

//0       1  2       3  4
//Faerun to Tristram = 65

func main() {

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		fields := strings.Fields(line)
		from, to, dstr := fields[0], fields[2], fields[4]
		d, err := strconv.Atoi(dstr)
		if err != nil {
			panic(err)
		}
		nodes[node(from)] = true
		nodes[node(to)] = true
		edges = append(edges, edge{from: node(from), to: node(to), d: d})
		edges = append(edges, edge{from: node(to), to: node(from), d: d})
	}
	total := len(nodes)
	for _, e := range edges {
		seen := map[node]bool{e.from: true}
		walk(total-1, []edge{e}, edges, seen)
	}
	fmt.Println("solution", solution, score(solution))
	fmt.Println("worst", badsolution, score(badsolution))
}

var solution = []edge{}
var best = math.MaxInt64
var worst = -1
var badsolution = []edge{}

func walk(i int, path []edge, edges []edge, seen map[node]bool) {
	if len(path) == i {
		if score(path) < best {
			solution = path
			best = score(path)
		}
		if score(path) > worst {
			badsolution = path
			worst = score(path)
		}
		return
	}

	e := path[len(path)-1]
	for _, e2 := range edges {
		if e.to == e2.from && !seen[e2.from] && !seen[e2.to] {
			seen[e2.from] = true
			walk(i, append(path, e2), edges, seen)
			delete(seen, e2.from)
		}
	}
}

func score(path []edge) int {
	total := 0
	for _, e := range path {
		total += e.d
	}
	return total
}
