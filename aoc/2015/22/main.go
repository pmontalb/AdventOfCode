package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"../../../utils"
)

type Effect struct {
	maxTurns int

	armor        int
	damage       int
	manaRecovery int

	remainingTurns int
}

type Magic struct {
	id       string
	manaCost int

	armor  int
	damage int

	healthRecovery int
	effect         *Effect
}
type Magics []Magic

func (magics Magics) Len() int           { return len(magics) }
func (magics Magics) Swap(i, j int)      { magics[i], magics[j] = magics[j], magics[i] }
func (magics Magics) Less(i, j int) bool { return magics[i].manaCost < magics[j].manaCost }

type Player struct {
	id string

	initialHealthPoint int
	initialManaStat    int

	currentHealthPoint int
	currentManaStat    int

	damageStat int
	armorStat  int
	dead       bool
}

func (p *Player) reset() {
	p.currentHealthPoint = p.initialHealthPoint
	p.currentManaStat = p.initialManaStat
	p.dead = false
}
func (p *Player) takeDamage(damage int) {
	p.currentHealthPoint -= damage
	p.dead = p.currentHealthPoint <= 0
}

func (p *Player) attack(q *Player) {
	damage := p.damageStat - q.armorStat
	if damage < 1 {
		damage = 1
	}

	q.takeDamage(damage)
	//fmt.Printf("player(%s|atk=%d) attacking player(%s|def=%d) dealing (%d) damage: player(%s) hp(%d -> %d) dead(%v)\n", p.id, p.currentAttackStat, q.id, q.currentDefenseStat, damage, q.id, q.currentHealthPoint+damage, q.currentHealthPoint, q.dead)
}

func (p *Player) castSpell(m *Magic, q *Player) {
	damage := m.damage
	q.takeDamage(damage)

	p.currentManaStat -= m.manaCost
	//fmt.Printf("player(%s|atk=%d) attacking player(%s|def=%d) dealing (%d) damage: player(%s) hp(%d -> %d) dead(%v)\n", p.id, p.currentAttackStat, q.id, q.currentDefenseStat, damage, q.id, q.currentHealthPoint+damage, q.currentHealthPoint, q.dead)
}

func applyEffect(p1 *Player, p2 *Player, e *Effect) bool {
	if e != nil && (*e).remainingTurns > 0 {
		(*e).remainingTurns--
		p1.armorStat += (*e).armor
		p1.currentManaStat += (*e).manaRecovery

		p2.takeDamage((*e).damage)

		return (*e).remainingTurns != 0
	}
	return true
}

func wearOffEffect(p *Player, e *Effect) {
	p.armorStat -= (*e).armor
}

func playTurn(p1 *Player, p2 *Player, magic *Magic, effects *map[string]*Effect) bool {

	fmt.Printf("* START p1(%s|hp=%d) p2(%s|hp=%d) effects(%v)\n", p1.id, p1.currentHealthPoint, p2.id, p2.currentHealthPoint, *effects)

	p1.damageStat = magic.damage
	p1.armorStat = magic.armor
	p1.currentHealthPoint += magic.healthRecovery

	canUseEffect := true
	_, found := (*effects)[magic.id]
	if magic.effect != nil && !found {
		(*effects)[magic.id] = magic.effect
		magic.effect.remainingTurns = magic.effect.maxTurns
		canUseEffect = false
	}

	for id, effect := range *effects {
		if canUseEffect || id != magic.id {
			if !applyEffect(p1, p2, effect) {
				delete(*effects, id)
				wearOffEffect(p1, effect)
			}
		}
	}

	fmt.Printf("** SETUP p1(%s|hp=%d|mgk=%s|atk=%d|def=%d) effects(%v)\n", p1.id, p1.currentHealthPoint, magic.id, p1.damageStat, p1.armorStat, *effects)
	p1.castSpell(magic, p2)

	if p2.dead {
		return true
	}

	for id, effect := range *effects {
		if !applyEffect(p1, p2, effect) {
			delete(*effects, id)
			wearOffEffect(p1, effect)
		}
	}

	if p2.dead {
		return true
	}

	p2.attack(p1)
	fmt.Printf("*** END p1(%s|hp=%d|mn=%d) p2(%s|hp=%d) effects(%v)\n", p1.id, p1.currentHealthPoint, p1.currentManaStat, p2.id, p2.currentHealthPoint, *effects)
	return !p1.dead
}

func optimize(player *Player, boss *Player, magicList []Magic) int {

	return 0
}

func test() {
	player := Player{}
	player.id = "Player"
	player.initialHealthPoint = 10
	player.initialManaStat = 250
	player.reset()

	boss := Player{}
	boss.id = "Boss"
	boss.initialHealthPoint = 13
	boss.damageStat = 8
	boss.reset()

	magicList := make([]Magic, 5)
	magicList[0] = Magic{"Magic Missile", 53, 0, 4, 0, nil}
	magicList[1] = Magic{"Drain", 73, 0, 2, 2, nil}
	magicList[2] = Magic{"Shield", 113, 0, 0, 0, &Effect{6, 7, 0, 0, 0}}
	magicList[3] = Magic{"Poison", 173, 0, 0, 0, &Effect{6, 0, 3, 0, 0}}
	magicList[4] = Magic{"Recharge", 229, 0, 0, 0, &Effect{5, 0, 0, 101, 0}}
	sort.Sort(Magics(magicList))

	// test 1
	effects := make(map[string]*Effect)

	/*
		-- Player turn --
		- Player has 10 hit points, 0 armor, 250 mana
		- Boss has 13 hit points
		Player casts Poison.

		-- Boss turn --
		- Player has 10 hit points, 0 armor, 77 mana
		- Boss has 13 hit points
		Poison deals 3 damage; its timer is now 5.
		Boss attacks for 8 damage.
	*/
	playTurn(&player, &boss, &magicList[3], &effects)
	if len(effects) != 1 {
		panic(0)
	}
	if effects[magicList[3].id].remainingTurns != 5 {
		panic(0)
	}
	if player.currentHealthPoint != 2 {
		panic(0)
	}
	if player.currentManaStat != 77 {
		panic(0)
	}
	if player.armorStat != 0 {
		panic(0)
	}
	if boss.currentHealthPoint != 10 {
		panic(0)
	}

	/*
		-- Player turn --
		- Player has 2 hit points, 0 armor, 77 mana
		- Boss has 10 hit points
		Poison deals 3 damage; its timer is now 4.
		Player casts Magic Missile, dealing 4 damage.

		-- Boss turn --
		- Player has 2 hit points, 0 armor, 24 mana
		- Boss has 3 hit points
		Poison deals 3 damage. This kills the boss, and the player wins.
	*/
	win := playTurn(&player, &boss, &magicList[0], &effects)
	if !win {
		panic(0)
	}
	if len(effects) != 1 {
		panic(0)
	}
	if effects[magicList[3].id].remainingTurns != 3 {
		panic(0)
	}
	if player.currentHealthPoint != 2 {
		panic(0)
	}
	if player.currentManaStat != 24 {
		panic(0)
	}
	if player.armorStat != 0 {
		panic(0)
	}
	if boss.currentHealthPoint > 0 {
		panic(0)
	}
	if !boss.dead {
		panic(0)
	}

	//fmt.Println("--------------------------")

	// test 2
	player.reset()

	boss.initialHealthPoint = 14
	boss.reset()

	effects = make(map[string]*Effect)

	/*
		-- Player turn --
		- Player has 10 hit points, 0 armor, 250 mana
		- Boss has 14 hit points
		Player casts Recharge.

		-- Boss turn --
		- Player has 10 hit points, 0 armor, 21 mana
		- Boss has 14 hit points
		Recharge provides 101 mana; its timer is now 4.
		Boss attacks for 8 damage!
	*/
	playTurn(&player, &boss, &magicList[4], &effects)
	if len(effects) != 1 {
		panic(0)
	}
	if effects[magicList[4].id].remainingTurns != 4 {
		panic(0)
	}
	if player.currentHealthPoint != 2 {
		panic(0)
	}
	if player.currentManaStat != 122 {
		panic(0)
	}
	if player.armorStat != 0 {
		panic(0)
	}
	if boss.currentHealthPoint != 14 {
		panic(0)
	}

	/*
			-- Player turn --
		- Player has 2 hit points, 0 armor, 122 mana
		- Boss has 14 hit points
		Recharge provides 101 mana; its timer is now 3.
		Player casts Shield, increasing armor by 7.

		-- Boss turn --
		- Player has 2 hit points, 7 armor, 110 mana
		- Boss has 14 hit points
		Shield's timer is now 5.
		Recharge provides 101 mana; its timer is now 2.
		Boss attacks for 8 - 7 = 1 damage!
	*/
	playTurn(&player, &boss, &magicList[2], &effects)
	if len(effects) != 2 {
		panic(0)
	}
	if effects[magicList[4].id].remainingTurns != 2 {
		panic(0)
	}
	if effects[magicList[2].id].remainingTurns != 5 {
		panic(0)
	}
	if player.currentHealthPoint != 1 {
		panic(0)
	}
	if player.currentManaStat != 211 {
		panic(0)
	}
	if player.armorStat != 7 {
		panic(0)
	}
	if boss.currentHealthPoint != 14 {
		panic(0)
	}

	/*
			-- Player turn --
		- Player has 1 hit point, 7 armor, 211 mana
		- Boss has 14 hit points
		Shield's timer is now 4.
		Recharge provides 101 mana; its timer is now 1.
		Player casts Drain, dealing 2 damage, and healing 2 hit points.

		-- Boss turn --
		- Player has 3 hit points, 7 armor, 239 mana
		- Boss has 12 hit points
		Shield's timer is now 3.
		Recharge provides 101 mana; its timer is now 0.
		Recharge wears off.
		Boss attacks for 8 - 7 = 1 damage!
	*/
	playTurn(&player, &boss, &magicList[1], &effects)
	if len(effects) != 1 {
		panic(0)
	}
	if effects[magicList[2].id].remainingTurns != 3 {
		panic(0)
	}
	if player.currentHealthPoint != 2 {
		panic(0)
	}
	if player.currentManaStat != 340 {
		panic(0)
	}
	if player.armorStat != 7 {
		panic(0)
	}
	if boss.currentHealthPoint != 12 {
		panic(0)
	}

	/*
			-- Player turn --
		- Player has 2 hit points, 7 armor, 340 mana
		- Boss has 12 hit points
		Shield's timer is now 2.
		Player casts Poison.

		-- Boss turn --
		- Player has 2 hit points, 7 armor, 167 mana
		- Boss has 12 hit points
		Shield's timer is now 1.
		Poison deals 3 damage; its timer is now 5.
		Boss attacks for 8 - 7 = 1 damage!
	*/
	playTurn(&player, &boss, &magicList[3], &effects)
	if len(effects) != 2 {
		panic(0)
	}
	if effects[magicList[2].id].remainingTurns != 1 {
		panic(0)
	}
	if effects[magicList[3].id].remainingTurns != 5 {
		panic(0)
	}
	if player.currentHealthPoint != 1 {
		panic(0)
	}
	if player.currentManaStat != 167 {
		panic(0)
	}
	if player.armorStat != 7 {
		panic(0)
	}
	if boss.currentHealthPoint != 9 {
		panic(0)
	}

	/*
			-- Player turn --
		- Player has 1 hit point, 7 armor, 167 mana
		- Boss has 9 hit points
		Shield's timer is now 0.
		Shield wears off, decreasing armor by 7.
		Poison deals 3 damage; its timer is now 4.
		Player casts Magic Missile, dealing 4 damage.

		-- Boss turn --
		- Player has 1 hit point, 0 armor, 114 mana
		- Boss has 2 hit points
		Poison deals 3 damage. This kills the boss, and the player wins.
	*/
	win = playTurn(&player, &boss, &magicList[0], &effects)
	if !win {
		panic(0)
	}
	if len(effects) != 1 {
		panic(0)
	}
	if effects[magicList[3].id].remainingTurns != 3 {
		panic(0)
	}
	if player.currentHealthPoint != 1 {
		panic(0)
	}
	if player.currentManaStat != 114 {
		panic(0)
	}
	if player.armorStat != 0 {
		panic(0)
	}
	if !boss.dead {
		panic(0)
	}
}

func main() {
	test()

	lines := utils.GetLines("input")

	player := Player{}
	player.id = "Player"
	player.initialHealthPoint = 50
	player.initialManaStat = 500
	player.reset()

	boss := Player{}
	boss.id = "Boss"
	boss.initialHealthPoint, _ = strconv.Atoi(strings.TrimSpace(strings.Split(lines[0], ":")[1]))
	boss.damageStat, _ = strconv.Atoi(strings.TrimSpace(strings.Split(lines[1], ":")[1]))
	boss.reset()

	magicList := make([]Magic, 5)
	magicList[0] = Magic{"Magic Missile", 53, 0, 4, 0, nil}
	magicList[1] = Magic{"Drain", 73, 0, 2, 2, nil}
	magicList[2] = Magic{"Shield", 113, 0, 0, 0, &Effect{6, 7, 0, 0, 0}}
	magicList[3] = Magic{"Poison", 173, 0, 0, 0, &Effect{6, 0, 3, 0, 0}}
	magicList[4] = Magic{"Recharge", 229, 0, 0, 0, &Effect{5, 0, 0, 101, 0}}
	sort.Sort(Magics(magicList))
}
