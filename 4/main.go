package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"strings"
)

var prefix = flag.String("p", "abcdef", "prefix")

func main() {
	flag.Parse()
	solved := false
	guess := ""
	i := -1
	for !solved {
		i++
		h := md5.New()
		guess = *prefix + itoa(i)
		h.Write([]byte(guess))
		result := h.Sum(nil)
		hex := fmt.Sprintf("%x", result)
		if strings.HasPrefix(hex, "000000") {
			solved = true
		}
	}
	fmt.Println(guess)
}

func itoa(i int) string {
	return fmt.Sprintf("%d", i)
}
