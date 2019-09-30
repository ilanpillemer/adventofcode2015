package main

import (
	"fmt"
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

//var effects = map[spell]timer{}
//var inEffect map[spell]bool

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

func main() {
	fmt.Println("example 1")
	example1()
	fmt.Println("------------------")
	fmt.Println("example 2")
	example2()
}

func example1() {
	p := &player{}
	b := &player{}
	p.name = "Player"
	p.hp = 10
	p.mana = 250

	b.name = "Boss"
	b.hp = 13
	b.damage = 8
	effects := map[spell]timer{}
	turn(Poison, p, b, effects)
	turn(MagicMissile, p, b, effects)

}

func example2() {
	p := &player{}
	b := &player{}
	p.name = "Player"
	p.hp = 10
	p.mana = 250

	b.name = "Boss"
	b.hp = 14
	b.damage = 8
	effects := map[spell]timer{}
	if over := turn(Recharge, p, b, effects); over {
		return
	}
	if over := turn(Shield, p, b, effects); over {
		return
	}
	if over := turn(Drain, p, b, effects); over {
		return
	}
	if over := turn(Poison, p, b, effects); over {
		return
	}
	if over := turn(MagicMissile, p, b, effects); over {
		return
	}

}

func stats(p *player, b *player) {
	fmt.Printf("- Player has %d hit points, %d armour, %d mana\n", p.hp, p.armour, p.mana)
	fmt.Printf("- Boss has %d hit points \n", b.hp)
}

func turn(s spell, p *player, b *player, effects map[spell]timer) bool {

	fmt.Println("-- Player turn --")
	stats(p, b)
	p.armour = 0
	//inEffect = map[spell]bool{} //clear
	if over := apply(p, b, effects); over {
		return true
	}

	if over := p.cast(s, b, effects); over {
		return true
	}
	fmt.Println()
	fmt.Println("-- Boss turn --")
	stats(p, b)
	if over := apply(p, b, effects); over {
		return true
	}
	if over := b.attack(p); over {
		return true
	}
	fmt.Println()
	return gameover(p, b)
}

func apply(p *player, b *player, effects map[spell]timer) bool {

	for s, timer := range effects {
		switch s {
		case Shield:
			p.armour = 7
			timer--
			if timer == 0 {
				delete(effects, Shield)
				//			delete(inEffect, Shield)
				fmt.Println("Shield wears off, decreasing armor by 7.")
				p.armour = 0
			}

			fmt.Printf("Shield's timer is now %d.\n", timer)
			if timer != 0 {
				effects[Shield] = timer
			}
		case Poison:
			b.hp -= 3
			timer--
			if timer == 0 {
				delete(effects, Poison)
				fmt.Println("Poison wears off")
			}
			fmt.Printf("Poison deals 3 damage; its timer is now %d.\n", timer)
			if timer != 0 {
				effects[Poison] = timer
			}
		case Recharge:
			p.mana += 101
			timer--
			if timer == 0 {
				delete(effects, Recharge)
				fmt.Println("Recharge wears off")
			}
			fmt.Printf("Recharge provides 101 mana; its timer is now %d.\n", timer)
			if timer != 0 {
				effects[Recharge] = timer
			}
		}
	}
	return gameover(p, b)
}

func (p *player) cast(s spell, op *player, effects map[spell]timer) bool {
	fmt.Println("Player casts", s)
	switch s {
	case MagicMissile:
		p.mana -= 53
		op.hp -= 4
		fmt.Println("Magic Missile does 4 damage")
		if over := gameover(p, op); over {
			return true
		}
	case Drain:
		p.mana -= 73
		p.hp += 2
		op.hp -= 2
		fmt.Println("Drain does 2 damage, heals 2 points")
		if over := gameover(p, op); over {
			return true
		}
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
	return gameover(p, op)
}

func (p *player) attack(op *player) bool {
	d := p.damage - op.armour
	if d < 1 {
		d = 1
	}
	op.hp = op.hp - d
	fmt.Printf("The %s deals %d-%d = %d damage; the %s goes down to %d hit points. \n", p.name, p.damage, op.armour, p.damage-op.armour, op.name, op.hp)
	return gameover(p, op)
}

func gameover(p *player, b *player) bool {
	if b.hp <= 0 {
		fmt.Println("fight over", b.name, "loses")
		return true
	}

	if p.hp <= 0 {
		fmt.Println("fight over", p.name, "loses")
		return true
	}

	return false

}
