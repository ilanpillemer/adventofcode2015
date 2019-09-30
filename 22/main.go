package main

import (
	"flag"
	"fmt"
)

type player struct {
	name   string
	hp     int
	damage int
	armour int
	mana   int
	spent  int
}
type spell int
type timer int

const (
	None spell = iota
	MagicMissile
	Drain
	Shield
	Poison
	Recharge
)

var verbose = flag.Bool("v", false, "verbose")

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
	//	fmt.Println("example 1")
	//	example1()
	//	fmt.Println("------------------")
	//	fmt.Println("example 2")
	//	example2()
	//
	//	fmt.Println()
	//	fmt.Println()

	flag.Parse()
	//	example3()
	//	os.Exit(0)

	min := 10000000
	p, b, effects := ini()
	over := false
	for i := 0; i < 100000000; i++ {
	fight:
		for {
			options := possible(effects)

			if p.mana < 53 {
				delete(options, MagicMissile)
			}
			if p.mana < 73 {
				delete(options, Drain)
			}
			if p.mana < 113 {
				delete(options, Shield)
			}
			if p.mana < 173 {
				delete(options, Poison)
			}
			if p.mana < 229 {
				delete(options, Recharge)
			}
			if len(options) == 0 {
				if *verbose {
					fmt.Println("TRY AGAIN")
				}
				p, b, effects = ini()
				continue
			}
			for k, _ := range options {

				if over = turn(k, p, b, effects); over {
					if b.hp <= 0 {
						if *verbose {
							fmt.Println("spent", p.spent)
						}
						if p.spent < min {
							min = p.spent
						}
						if *verbose {
							fmt.Println("min", min)
						}
						p, b, effects = ini()
						break fight
					}
					// try again
					if *verbose {
						fmt.Println("TRY AGAIN")
					}
					p, b, effects = ini()
				}
				if p.mana < 53 {
					delete(options, MagicMissile)
				}
				if p.mana < 73 {
					delete(options, Drain)
				}
				if p.mana < 113 {
					delete(options, Shield)
				}
				if p.mana < 173 {
					delete(options, Poison)
				}
				if p.mana < 229 {
					delete(options, Recharge)
				}
				if len(options) == 0 {
					if *verbose {
						fmt.Println("TRY AGAIN")
					}
					p, b, effects = ini()
					break fight
				}
				if p.spent >= 973 {
					p, b, effects = ini()
					break fight
				}

			}
		}
	}
	fmt.Println("MIN:", min)
}

func ini() (*player, *player, map[spell]timer) {
	p := &player{}
	b := &player{}
	p.name = "Player"
	p.hp = 50
	p.mana = 500

	b.name = "Boss"
	b.hp = 51
	b.damage = 9
	effects := map[spell]timer{}
	return p, b, effects
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

func example3() {
	p, b, effects := ini()
	turn(Poison, p, b, effects)
	turn(Recharge, p, b, effects)
	turn(MagicMissile, p, b, effects)
	turn(Drain, p, b, effects)
	turn(Shield, p, b, effects)
	turn(Poison, p, b, effects)
	turn(MagicMissile, p, b, effects)
	turn(MagicMissile, p, b, effects) //turn(Drain, p, b, effects)
	turn(MagicMissile, p, b, effects)
	fmt.Println(p.spent)
}

func possible(effects map[spell]timer) map[spell]bool {
	options := map[spell]bool{MagicMissile: true, Drain: true, Shield: true, Poison: true, Recharge: true}
	for k, _ := range effects {
		delete(options, k)
	}

	return options
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
	if *verbose {
		fmt.Printf("- Player has %d hit points, %d armour, %d mana\n", p.hp, p.armour, p.mana)
		fmt.Printf("- Boss has %d hit points \n", b.hp)
	}
}

func turn(s spell, p *player, b *player, effects map[spell]timer) bool {
	if *verbose {
		fmt.Println("-- Player turn --")
	}
	stats(p, b)
	p.armour = 0
	//inEffect = map[spell]bool{} //clear
	if over := apply(p, b, effects); over {
		return true
	}

	if over := p.cast(s, b, effects); over {
		return true
	}
	if *verbose {
		fmt.Println()
		fmt.Println("-- Boss turn --")
	}
	stats(p, b)
	if over := apply(p, b, effects); over {
		return true
	}
	if over := b.attack(p); over {
		return true
	}
	if *verbose {
		fmt.Println()
	}
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
				if *verbose {
					fmt.Println("Shield wears off, decreasing armor by 7.")
				}
				p.armour = 0
			}
			if *verbose {
				fmt.Printf("Shield's timer is now %d.\n", timer)
			}
			if timer != 0 {
				effects[Shield] = timer
			}
		case Poison:
			b.hp -= 3
			timer--
			if timer == 0 {
				delete(effects, Poison)
				if *verbose {
					fmt.Println("Poison wears off")
				}
			}
			if *verbose {
				fmt.Printf("Poison deals 3 damage; its timer is now %d.\n", timer)
			}
			if timer != 0 {
				effects[Poison] = timer
			}
		case Recharge:
			p.mana += 101
			timer--
			if timer == 0 {
				delete(effects, Recharge)
				if *verbose {
					fmt.Println("Recharge wears off")
				}
			}
			if *verbose {
				fmt.Printf("Recharge provides 101 mana; its timer is now %d.\n", timer)
			}
			if timer != 0 {
				effects[Recharge] = timer
			}
		}
	}
	return gameover(p, b)
}

func (p *player) cast(s spell, op *player, effects map[spell]timer) bool {
	if *verbose {
		fmt.Println("Player casts", s)
	}
	switch s {
	case MagicMissile:
		p.mana -= 53
		p.spent += 53
		op.hp -= 4
		if *verbose {
			fmt.Println("Magic Missile does 4 damage")
		}
		if over := gameover(p, op); over {
			return true
		}
	case Drain:
		p.mana -= 73
		p.spent += 73
		p.hp += 2
		op.hp -= 2
		if *verbose {
			fmt.Println("Drain does 2 damage, heals 2 points")
		}
		if over := gameover(p, op); over {
			return true
		}
	case Shield:
		p.mana -= 113
		p.spent += 113
		effects[Shield] = timer(6)
	case Poison:
		p.mana -= 173
		p.spent += 173
		effects[Poison] = timer(6)
	case Recharge:
		p.mana -= 229
		p.spent += 229
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
	if *verbose {
		fmt.Printf("The %s deals %d-%d = %d damage; the %s goes down to %d hit points. \n", p.name, p.damage, op.armour, p.damage-op.armour, op.name, op.hp)
	}
	return gameover(p, op)
}

func gameover(p *player, b *player) bool {
	if b.hp <= 0 {
		if *verbose {
			fmt.Println("fight over", b.name, "loses")
		}
		return true
	}

	if p.hp <= 0 {
		if *verbose {
			fmt.Println("fight over", p.name, "loses")
		}
		return true
	}

	return false

}
