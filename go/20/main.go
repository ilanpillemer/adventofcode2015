package main

import "fmt"

var store = make([]int, 3400000)

//34000000
func main() {
	//part1()
	for i := 1; i < 3400000; i++ {
		for j, count := i, 0; j < 3400000 && count < 50; j += i {
			store[j] += (i * 11)
			count++
		}
	}

	sum := 0
	i := 0
	for sum < 34000000 {
		i++
		sum = store[i]
	}
	fmt.Println(i, sum)

}

func part1() {
	for i := 1; i < 3400000; i++ {
		for j := i; j < 3400000; j += i {
			store[j] += (i * 10)
		}
	}

	sum := 0
	i := 0
	for sum < 34000000 {
		i++
		sum = store[i]
	}
	fmt.Println(i, sum)
}

//below is too slow
func factors(n int) []int {
	f := []int{}
	for i := 1; i < n+1; i++ {
		if n%i == 0 {
			f = append(f, i)
		}
	}
	return f
}

func sum(n []int) int {
	sum := 0
	for _, x := range n {
		sum += (x * 10)
	}
	return sum
}
