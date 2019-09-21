package main

import (
	"flag"
	"fmt"
	"strings"
)

var password = flag.String("p", "abcdefgh", "old password")

func main() {
	flag.Parse()
	valid := false
	next := []byte(*password)
	for !valid {
		next = inc(next)
		valid = isValid(next)
	}
	fmt.Printf("%s\n", next)

}
func inc(s []byte) []byte {
	for i := len(s) - 1; i > -1; i-- {
		if s[i] == 'z' {
			s[i] = 'a'
			continue
		}
		s[i] = s[i] + 1
		break
	}
	return s
}

func isValid(s []byte) bool {
	return hasStraight(s) && visFine(s) && hasTwoPairs(s)
}

func hasStraight(s []byte) bool {
	count := 1
	var prev byte
	for i, c := range s {
		if i == 0 {
			prev = c
			continue
		}
		if (c - prev) == 1 {
			count++
		} else {
			count = 1
		}
		if count == 3 {
			return true
		}
		prev = c
	}
	return false
}

func visFine(s []byte) bool {
	for _, c := range string(s) {
		switch c {
		case 'o', 'i', 'l':
			return false
		}
	}
	return true
}

func hasTwoPairs(s []byte) bool {
	//could also generate
	pairs := []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg", "hh", "ii", "jj", "kk", "ll", "mm", "nn", "oo", "pp", "qq", "rr", "ss", "tt", "uu", "vv", "ww", "xx", "yy", "zz"}
	total := 0

	for _, pair := range pairs {
		v := strings.Count(string(s), pair)
		total += v
	}

	return total >= 2
}

//	fmt.Printf("%s\n", inc([]byte("abc")))
//	fmt.Printf("%s\n", inc([]byte("abz")))
//	fmt.Printf("%t\n", hasStraight([]byte("abz")))
//	fmt.Printf("%t\n", hasStraight([]byte("abc")))
//	fmt.Printf("%t\n", hasStraight([]byte("abdghi")))
//	fmt.Printf("%t\n", visFine([]byte("abdghi")))
//	fmt.Printf("%t\n", visFine([]byte("abdgh")))
//	fmt.Printf("%t\n", hasTwoPairs([]byte("aa")))
//	fmt.Printf("%t\n", hasTwoPairs([]byte("aaa")))
//	fmt.Printf("%t\n", hasTwoPairs([]byte("aaaa")))
//	fmt.Printf("%t\n", hasTwoPairs([]byte("aaba")))
//	fmt.Printf("%t\n", hasTwoPairs([]byte("aababb")))
