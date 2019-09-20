package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var wires = map[string]uint16{}

type Op int

const (
	NONE Op = iota
	AND
	LSHIFT
	RSHIFT
	OR
	NOT
)

type statement struct {
	gate     string
	wire     string
	executed bool
}

var prog = []statement{}

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		instr := strings.Split(line, "->")
		gate, wire := strings.TrimSpace(instr[0]), strings.TrimSpace(instr[1])
		prog = append(prog, statement{wire: wire, gate: gate})
	}

	allexecuted := false
	for !allexecuted {
		count := 0
		for i, s := range prog {
			if !s.executed {
				count++
			}
			switch {
			case strings.Contains(s.gate, "AND"):
				instruction(&s, s.wire, AND, strings.Split(s.gate, "AND"))
			case strings.Contains(s.gate, "LSHIFT"):
				instruction(&s, s.wire, LSHIFT, strings.Split(s.gate, "LSHIFT"))
			case strings.Contains(s.gate, "RSHIFT"):
				instruction(&s, s.wire, RSHIFT, strings.Split(s.gate, "RSHIFT"))
			case strings.Contains(s.gate, "OR"):
				instruction(&s, s.wire, OR, strings.Split(s.gate, "OR"))
			case strings.Contains(s.gate, "NOT"):
				instruction(&s, s.wire, NOT, strings.Split(s.gate, "NOT"))
			default:
				s.gate = strings.TrimSpace(s.gate)
				if alive(s.gate) {
					signal(s.wire, eval(s.gate))
					s.executed = true
				}
			}
			prog[i] = s
		}
		if count == 0 {
			allexecuted = true
		}

	}
	fmt.Println("a", wires["a"])

}

func instruction(s *statement, wire string, op Op, vars []string) {
	if s.executed {
		return
	}
	switch op {
	case AND:
		if alive(vars[0]) && alive(vars[1]) {
			signal(wire, eval(vars[0])&eval(vars[1]))
			s.executed = true
		}
	case OR:
		if alive(vars[0]) && alive(vars[1]) {
			signal(wire, (eval(vars[0]) | eval(vars[1])))
			s.executed = true
		}
	case NOT:
		if alive(vars[1]) {
			signal(wire, ^eval(vars[1]))
			s.executed = true
		}
	case LSHIFT:
		if alive(vars[0]) && alive(vars[1]) {
			signal(wire, eval(vars[0])<<eval(vars[1]))
			s.executed = true
		}
	case RSHIFT:
		if alive(vars[0]) && alive(vars[1]) {
			signal(wire, eval(vars[0])>>eval(vars[1]))
			s.executed = true
		}
	}
}

func signal(wire string, signal uint16) {
	wire = strings.TrimSpace(wire)
	wires[wire] = signal
}

//is receiving signal
func alive(s string) bool {
	s = strings.TrimSpace(s)
	if _, ok := wires[s]; ok {
		return true
	}

	if _, err := strconv.ParseUint(s, 10, 16); err == nil {
		return true
	}
	return false
}

func eval(s string) uint16 {
	s = strings.TrimSpace(s)

	//is it an integer?
	if val, err := strconv.ParseUint(s, 10, 16); err == nil {
		return uint16(val)
	}

	return wires[s]
}
