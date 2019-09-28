package main

import (
	"bufio"
	"fmt"
	"os"
)

type char struct {
	name   string
	hp     int
	dscore int
	ascore int
}

var player, boss char

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		line := in.Text()
		fmt.Println(line)
	}

	actual(&player, &boss)
	fmt.Println(&player, &boss)
	player.name = "player"
	boss.name = "boss"
	for {
		turn()
	}

}

func example(p *char, b *char) {
	p.hp = 8
	p.dscore = 5
	p.ascore = 8
	b.hp = 12
	b.dscore = 7
	b.ascore = 2
}

func actual(p *char, b *char) {
	p.hp = 100
	// just did this manually
	p.dscore = 7
	p.ascore = 2
	//
	b.hp = 104
	b.dscore = 8
	b.ascore = 1
}

func turn() {
	attack(&player, &boss)
	attack(&boss, &player)
}

func attack(att *char, def *char) {
	d := att.dscore - def.ascore
	if d < 1 {
		d = 1
	}
	def.hp = def.hp - d
	fmt.Printf("The %s deals %d-%d = %d damage; the %s goes down to %d hit points. \n", att.name, att.dscore, def.ascore, att.dscore-def.ascore, def.name, def.hp)
	if att.hp <= 0 || def.hp <= 0 {
		fmt.Println("fight over", def.name, "loses")
		os.Exit(0)
	}
}
