package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var options = map[string]bool{}

const mol = "CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr"

var edges = map[string][]string{}

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		fields := strings.Split(line, "=>")
		from, to := strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1])

		e, ok := edges[from]
		if !ok {
			e = []string{}
		}
		e = append(e, to)
		edges[from] = e
	}

	randomSearch()

}

//just ran this a few times and took the smallest
func randomSearch() {
	steps := 0
	target := mol
	count := 0
	for target != "e" {
		count++
		state := target
		for to, e := range edges {
			for _, from := range e {
				re := regexp.MustCompile(from)
				loc := re.FindStringIndex(target)
				if loc != nil {
					target = target[:loc[0]] + to + target[loc[1]:]
					steps++
				}
			}
		}
		if target == state {
			target = mol
		}

	}
	fmt.Println("steps", steps)
}

func calibrate() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		fields := strings.Split(line, "=>")
		from, to := strings.TrimSpace(fields[0]), strings.TrimSpace(fields[1])
		re := regexp.MustCompile(from)
		locs := re.FindAllStringIndex(mol, -1)
		for _, loc := range locs {
			option := mol[:loc[0]] + to + mol[loc[1]:]
			options[option] = true
		}
	}
	fmt.Println(len(options))
}
