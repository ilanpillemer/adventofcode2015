package main

import (
	"fmt"
	"os"
)

type player struct {
	name   string
	hp     int
	damage int
	armour int
	mana   int
}
type spell int
type timer int

var effects = map[spell]timer{}
var inEffect map[spell]bool

const (
	None spell = iota
	MagicMissile
	Drain
	Shield
	Poison
	Recharge
)

func (s spell) String() string {
	switch s {
	case MagicMissile:
		return "Magic Missile"
	case Drain:
		return "Drain"
	case Shield:
		return "Shield"
	case Poison:
		return "Poison"
	case Recharge:
		return "Recharge"
	}
	return "None"
}

var b = &player{}
var p = &player{}

func main() {
	//example1()
	example2()
}

func example1() {
	p.name = "Player"
	p.hp = 10
	p.mana = 250

	b.name = "Boss"
	b.hp = 13
	b.damage = 8

	turn(Poison)
	turn(MagicMissile)

}

func example2() {
	p.name = "Player"
	p.hp = 10
	p.mana = 250

	b.name = "Boss"
	b.hp = 14
	b.damage = 8

	turn(Recharge)
	turn(Shield)
	turn(Drain)
	turn(Poison)
	turn(MagicMissile)

}

func stats() {
	fmt.Printf("- Player has %d hit points, %d armour, %d mana\n", p.hp, p.armour, p.mana)
	fmt.Printf("- Boss has %d hit points \n", b.hp)
}

func turn(s spell) {
	defer gameover()

	fmt.Println("-- Player turn --")
	stats()
	p.armour = 0
	inEffect = map[spell]bool{} //clear
	apply()
	p.cast(s, b)
	fmt.Println()
	fmt.Println("-- Boss turn --")
	stats()
	apply()
	b.attack(p)
	fmt.Println()
}

func apply() {
	defer gameover()
	for s, timer := range effects {
		switch s {
		case Shield:
			p.armour = 7
			//	if !inEffect[Shield] {
			timer--
			inEffect[Shield] = true
			if timer == 0 {
				delete(effects, Shield)
				delete(inEffect, Shield)
				fmt.Println("Shield wears off, decreasing armor by 7.")
				p.armour = 0
			}
			//		}
			fmt.Printf("Shield's timer is now %d.\n", timer)
			if timer != 0 {
				effects[Shield] = timer
			}
		case Poison:
			b.hp -= 3
			//	if !inEffect[Poison] {
			timer--
			inEffect[Poison] = true
			if timer == 0 {
				delete(effects, Poison)
				delete(inEffect, Poison)
				fmt.Println("Poison wears off")
			}
			//	}
			fmt.Printf("Poison deals 3 damage; its timer is now %d.\n", timer)
			if timer != 0 {
				effects[Poison] = timer
			}
		case Recharge:
			p.mana += 101
			//	if !inEffect[Recharge] {
			timer--
			inEffect[Recharge] = true
			if timer == 0 {
				delete(effects, Recharge)
				delete(inEffect, Recharge)
				fmt.Println("Recharge wears off")
			}
			//	}
			fmt.Printf("Recharge provides 101 mana; its timer is now %d.\n", timer)
			if timer != 0 {
				effects[Recharge] = timer
			}
		}
	}
}

func (p *player) cast(s spell, op *player) {
	defer gameover()
	fmt.Println("Player casts", s)
	switch s {
	case MagicMissile:
		p.mana -= 53
		op.hp -= 4
		fmt.Println("Magic Missile does 4 damage")
	case Drain:
		p.mana -= 73
		p.hp += 2
		op.hp -= 2
		fmt.Println("Drain does 2 damage, heals 2 points")
	case Shield:
		p.mana -= 113
		effects[Shield] = timer(6)
	case Poison:
		p.mana -= 173
		effects[Poison] = timer(6)
	case Recharge:
		p.mana -= 229
		effects[Recharge] = timer(5)
	}

}

func (p *player) attack(op *player) {
	defer gameover()
	d := p.damage - op.armour
	if d < 1 {
		d = 1
	}
	op.hp = op.hp - d
	fmt.Printf("The %s deals %d-%d = %d damage; the %s goes down to %d hit points. \n", p.name, p.damage, op.armour, p.damage-op.armour, op.name, op.hp)

}

func gameover() {
	if b.hp <= 0 {
		fmt.Println("fight over", b.name, "loses")
		os.Exit(0)
	}

	if p.hp <= 0 {
		fmt.Println("fight over", p.name, "loses")
		os.Exit(0)
	}

}
