package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type register string

var registers = map[register]int{"a": 1, "b": 0}

var a, b register

var ptr = 0

func hlf(r register) {
	registers[r] = registers[r] / 2
}
func tpl(r register) {
	registers[r] = registers[r] * 3
}
func inc(r register) {
	registers[r] = registers[r] + 1
}

func jmp(offset int) {
	ptr += offset
}
func jie(r register, offset int) {
	if registers[r]%2 == 0 {
		ptr += offset
	} else {
		ptr++
	}

}
func jio(r register, offset int) {
	if registers[r] == 1 {
		ptr += offset
	} else {
		ptr++
	}
}

func main() {
	program := []string{}
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		program = append(program, line)
	}
	fmt.Println(program)
	for {
		if ptr >= len(program) {
			fmt.Println("a", registers["a"])
			fmt.Println("b", registers["b"])
			os.Exit(0)
		}

		fields := strings.Fields(program[ptr])
		instr := fields[0]
		switch instr {
		case "hlf":
			hlf(register(fields[1]))
		case "tpl":
			tpl(register(fields[1]))
		case "inc":
			inc(register(strings.TrimSpace(fields[1])))
		case "jmp":
			jmp(atoi(fields[1]))
			continue
		case "jie":
			jie(register(strings.TrimSuffix(fields[1], ",")), atoi(fields[2]))
			continue
		case "jio":
			jio(register(strings.TrimSuffix(fields[1], ",")), atoi(fields[2]))
			continue
		default:
			panic("hey ho the merry o")
		}
		ptr++
	}
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
