package main

import "fmt"

type ing struct {
	cap int
	dur int
	fla int
	tex int
	cal int
}

//Frosting: capacity 4, durability -2, flavor 0, texture 0, calories 5
//Candy: capacity 0, durability 5, flavor -1, texture 0, calories 8
//Butterscotch: capacity -1, durability 0, flavor 5, texture 0, calories 6
//Sugar: capacity 0, durability 0, flavor -2, texture 2, calories 1

var frosting = ing{
	cap: 4,
	dur: -2,
	fla: 0,
	tex: 0,
	cal: 5,
}

var candy = ing{
	cap: 0,
	dur: 5,
	fla: -1,
	tex: 0,
	cal: 8,
}

var butters = ing{
	cap: -1,
	dur: 0,
	fla: 5,
	tex: 0,
	cal: 6,
}

var sugar = ing{
	cap: 0,
	dur: 0,
	fla: -2,
	tex: 2,
	cal: 1,
}

//Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
//Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
var Butterscotch = ing{
	cap: -1,
	dur: -2,
	fla: 6,
	tex: 3,
	cal: 8,
}

var Cinnamon = ing{
	cap: 2,
	dur: 3,
	fla: -2,
	tex: -1,
	cal: 3,
}

func main() {

	//fmt.Println(scoreTest(44,56))
	max := 0
	for i := 1; i < 100; i++ {
		for j := 1; j < 100; j++ {
			for k := 1; k < 100; k++ {
				for l := 1; l < 100; l++ {
					if i+j+k+l == 100 {
						//	fmt.Println(i, j, k, l, "score", score(i, j, k, l))
						most := score(i, j, k, l)
						if most > max {
							max = most
						}
					}
				}
			}
		}
	}
	fmt.Println(max)
}

//A capacity of 44*-1 + 56*2 = 68
//A durability of 44*-2 + 56*3 = 80
//A flavor of 44*6 + 56*-2 = 152
//A texture of 44*3 + 56*-1 = 76

func scoreTest(f, c int) int {
	capacity := Butterscotch.cap*f + Cinnamon.cap*c
	durability := Butterscotch.dur*f + Cinnamon.dur*c
	flavor := Butterscotch.fla*f + Cinnamon.fla*c
	texture := Butterscotch.tex*f + Cinnamon.tex*c
	return val(capacity) * val(durability) * val(flavor) * val(texture)
}

func score(f, c, b, s int) int {
	capacity := frosting.cap*f + candy.cap*c + butters.cap*b + sugar.cap*s
	durability := frosting.dur*f + candy.dur*c + butters.dur*b + sugar.dur*s
	flavor := frosting.fla*f + candy.fla*c + butters.fla*b + sugar.fla*s
	texture := frosting.tex*f + candy.tex*c + butters.tex*b + sugar.tex*s
	cal := frosting.cal*f + candy.cal*c + butters.cal*b + sugar.cal*s

	if cal != 500 {
		return 0
	}

	return val(capacity) * val(durability) * val(flavor) * val(texture)
}

func val(v int) int {
	if v > 0 {
		return v
	}
	return 0
}

//Frosting: capacity 4, durability -2, flavor 0, texture 0, calories 5
//Candy: capacity 0, durability 5, flavor -1, texture 0, calories 8
//Butterscotch: capacity -1, durability 0, flavor 5, texture 0, calories 6
//Sugar: capacity 0, durability 0, flavor -2, texture 2, calories 1
