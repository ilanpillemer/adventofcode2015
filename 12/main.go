package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"unicode"
	"unicode/utf8"
)

//var ignore = "red"

var ignore = flag.String("i", "wtf", "ignore objects with this value")

func main() {
	flag.Parse()
	part2()
}

//part 2 converts from json and then walks the json, ignoring the specified object eg "red"
func part2() {
	var v interface{}
	decoder := json.NewDecoder(os.Stdin)
	err := decoder.Decode(&v)
	if err != nil {
		panic(err)
	}
	total := walk(v, 0)
	fmt.Printf("%.0f \n", total)
}

func walk(v interface{}, total float64) float64 {
	switch object := v.(type) {
	case map[string]interface{}:
		for _, v := range object {
			if str, ok := v.(string); ok {
				if str == *ignore { //if this is a red object, bail
					return total
				}
			}
		}

		for _, v := range object {
			switch v := v.(type) {
			case float64: //add up any numbers
				total += v
			case []interface{}: //if it has an array add up any of its numbers
				total = walk(v, total)
			case map[string]interface{}:
				total = walk(v, total)
			}
		}
		return total
	case []interface{}:
		for _, v := range object {
			switch v := v.(type) {
			case float64:
				total += v
			case []interface{}:
				total = walk(v, total)
			case map[string]interface{}:
				total = walk(v, total)
			}

		}
	}

	return total
}

//just lexes the byte array and finds all numbers and adds them..
//this doesnt work for part 2 as you cant look backwards or forwards
//in time easily and the "red" property can exist long long after or
//before the number in the byte array.
func part1() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanRunes)
	memory := 0
	total := 0
	inNumber := false
	current := ""
	for in.Scan() {
		c, _ := utf8.DecodeRuneInString(in.Text())
		switch {
		case c == '-':
			current = "-"
			inNumber = true
		case unicode.IsNumber(c):
			current += string(c)
			inNumber = true
		case inNumber:
			i := atoi(current)
			memory += i
			current = ""
			inNumber = false
		}
	}
	total = memory
	fmt.Println("total", total)
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
