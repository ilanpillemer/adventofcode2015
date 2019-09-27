package main

import (
	"fmt"
	"regexp"
	"testing"
)

const testmol = "HOH"

var testoptions = map[string]bool{}

func Test(t *testing.T) {

	do("H", "HO")

	for k := range testoptions {
		fmt.Println(k)
	}

}

func do(from, to string) {
	re := regexp.MustCompile(from)
	locs := re.FindAllStringIndex(testmol, -1)
	fmt.Println(locs)
	for _, loc := range locs {
		option := testmol[:loc[0]] + to + testmol[loc[1]:]
		testoptions[option] = true
	}
}
