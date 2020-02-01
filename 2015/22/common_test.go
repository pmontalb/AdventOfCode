package main

import (
	"sort"
	"testing"
)

func TestCase1(t *testing.T) {
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
	magicList[0] = Magic{"Magic Missile", 53, 0, 4, 0, Effect{}}
	magicList[1] = Magic{"Drain", 73, 0, 2, 2, Effect{}}
	magicList[2] = Magic{"Shield", 113, 0, 0, 0, Effect{6, 7, 0, 0, true, 0}}
	magicList[3] = Magic{"Poison", 173, 0, 0, 0, Effect{6, 0, 3, 0, false, 0}}
	magicList[4] = Magic{"Recharge", 229, 0, 0, 0, Effect{5, 0, 0, 101, false, 0}}
	sort.Sort(Magics(magicList))

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
	PlayTurn(&player, &boss, &magicList[3], &effects, false)
	if len(effects) != 1 {
		t.Error()
	}
	if effects[magicList[3].id].remainingTurns != 5 {
		t.Error()
	}
	if player.currentHealthPoint != 2 {
		t.Error()
	}
	if player.currentManaStat != 77 {
		t.Error()
	}
	if player.armorStat != 0 {
		t.Error()
	}
	if boss.currentHealthPoint != 10 {
		t.Error()
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
	win, _ := PlayTurn(&player, &boss, &magicList[0], &effects, false)
	if !win {
		t.Error()
	}
	if len(effects) != 1 {
		t.Error()
	}
	if effects[magicList[3].id].remainingTurns != 3 {
		t.Error()
	}
	if player.currentHealthPoint != 2 {
		t.Error()
	}
	if player.currentManaStat != 24 {
		t.Error()
	}
	if player.armorStat != 0 {
		t.Error()
	}
	if boss.currentHealthPoint > 0 {
		t.Error()
	}
	if !boss.dead {
		t.Error()
	}
}

func TestCase2(t *testing.T) {
	player := Player{}
	player.id = "Player"
	player.initialHealthPoint = 10
	player.initialManaStat = 250
	player.reset()

	boss := Player{}
	boss.id = "Boss"
	boss.initialHealthPoint = 14
	boss.damageStat = 8
	boss.reset()

	magicList := make([]Magic, 5)
	magicList[0] = Magic{"Magic Missile", 53, 0, 4, 0, Effect{}}
	magicList[1] = Magic{"Drain", 73, 0, 2, 2, Effect{}}
	magicList[2] = Magic{"Shield", 113, 0, 0, 0, Effect{6, 7, 0, 0, true, 0}}
	magicList[3] = Magic{"Poison", 173, 0, 0, 0, Effect{6, 0, 3, 0, false, 0}}
	magicList[4] = Magic{"Recharge", 229, 0, 0, 0, Effect{5, 0, 0, 101, false, 0}}
	sort.Sort(Magics(magicList))

	effects := make(map[string]*Effect)

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
	PlayTurn(&player, &boss, &magicList[4], &effects, false)
	if len(effects) != 1 {
		t.Error()
	}
	if effects[magicList[4].id].remainingTurns != 4 {
		t.Error()
	}
	if player.currentHealthPoint != 2 {
		t.Error()
	}
	if player.currentManaStat != 122 {
		t.Error()
	}
	if player.armorStat != 0 {
		t.Error()
	}
	if boss.currentHealthPoint != 14 {
		t.Error()
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
	PlayTurn(&player, &boss, &magicList[2], &effects, false)
	if len(effects) != 2 {
		t.Error()
	}
	if effects[magicList[4].id].remainingTurns != 2 {
		t.Error()
	}
	if effects[magicList[2].id].remainingTurns != 5 {
		t.Error()
	}
	if player.currentHealthPoint != 1 {
		t.Error()
	}
	if player.currentManaStat != 211 {
		t.Error()
	}
	if player.armorStat != 7 {
		t.Error()
	}
	if boss.currentHealthPoint != 14 {
		t.Error()
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
	PlayTurn(&player, &boss, &magicList[1], &effects, false)
	if len(effects) != 1 {
		t.Error()
	}
	if effects[magicList[2].id].remainingTurns != 3 {
		t.Error()
	}
	if player.currentHealthPoint != 2 {
		t.Error()
	}
	if player.currentManaStat != 340 {
		t.Error()
	}
	if player.armorStat != 7 {
		t.Error()
	}
	if boss.currentHealthPoint != 12 {
		t.Error()
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
	PlayTurn(&player, &boss, &magicList[3], &effects, false)
	if len(effects) != 2 {
		t.Error()
	}
	if effects[magicList[2].id].remainingTurns != 1 {
		t.Error()
	}
	if effects[magicList[3].id].remainingTurns != 5 {
		t.Error()
	}
	if player.currentHealthPoint != 1 {
		t.Error()
	}
	if player.currentManaStat != 167 {
		t.Error()
	}
	if player.armorStat != 7 {
		t.Error()
	}
	if boss.currentHealthPoint != 9 {
		t.Error()
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
	win, _ := PlayTurn(&player, &boss, &magicList[0], &effects, false)
	if !win {
		t.Error()
	}
	if len(effects) != 1 {
		t.Error()
	}
	if effects[magicList[3].id].remainingTurns != 3 {
		t.Error()
	}
	if player.currentHealthPoint != 1 {
		t.Error()
	}
	if player.currentManaStat != 114 {
		t.Error()
	}
	if player.armorStat != 0 {
		t.Error()
	}
	if !boss.dead {
		t.Error()
	}
}

func TestCase3(t *testing.T) {
	player := Player{}
	player.id = "Player"
	player.initialHealthPoint = 50
	player.initialManaStat = 500
	player.reset()

	boss := Player{}
	boss.id = "Boss"
	boss.initialHealthPoint = 55
	boss.damageStat = 8
	boss.reset()

	magicList := make([]Magic, 5)
	magicList[0] = Magic{"Magic Missile", 53, 0, 4, 0, Effect{}}
	magicList[1] = Magic{"Drain", 73, 0, 2, 2, Effect{}}
	magicList[2] = Magic{"Shield", 113, 0, 0, 0, Effect{6, 7, 0, 0, true, 0}}
	magicList[3] = Magic{"Poison", 173, 0, 0, 0, Effect{6, 0, 3, 0, false, 0}}
	magicList[4] = Magic{"Recharge", 229, 0, 0, 0, Effect{5, 0, 0, 101, false, 0}}
	sort.Sort(Magics(magicList))

	effects := make(map[string]*Effect)

	/*
		{Player 50 500 7 52 4 0 false}
		{Boss 55 0 -1 0 8 0 true}
		{Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
		{Poison 173 0 0 0 {6 0 3 0 false 0}}
		{Recharge 229 0 0 0 {5 0 0 101 false 0}}
		{Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
		{Shield 113 0 0 0 {6 7 0 0 true 0}}
		{Poison 173 0 0 0 {6 0 3 0 false 0}}
		{Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
		{Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
		{Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
	*/
	win, _ := PlayTurn(&player, &boss, &magicList[0], &effects, false) // {Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
	if win || player.dead || boss.dead {
		t.Error()
	}
	win, _ = PlayTurn(&player, &boss, &magicList[3], &effects, false) // {Poison 173 0 0 0 {6 0 3 0 false 0}}
	if win || player.dead || boss.dead {
		t.Error()
	}
	win, _ = PlayTurn(&player, &boss, &magicList[4], &effects, false) // {Recharge 229 0 0 0 {5 0 0 101 false 0}}
	if win || player.dead || boss.dead {
		t.Error()
	}
	win, _ = PlayTurn(&player, &boss, &magicList[0], &effects, false) // {Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
	if win || player.dead || boss.dead {
		t.Error()
	}
	win, _ = PlayTurn(&player, &boss, &magicList[2], &effects, false) // {Shield 113 0 0 0 {6 7 0 0 true 0}}
	if win || player.dead || boss.dead {
		t.Error()
	}
	win, _ = PlayTurn(&player, &boss, &magicList[3], &effects, false) // {Poison 173 0 0 0 {6 0 3 0 false 0}}
	if win || player.dead || boss.dead {
		t.Error()
	}
	win, _ = PlayTurn(&player, &boss, &magicList[0], &effects, false) // {Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
	if win || player.dead || boss.dead {
		t.Error()
	}
	win, _ = PlayTurn(&player, &boss, &magicList[0], &effects, false) // {Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
	if win || player.dead || boss.dead {
		t.Error()
	}
	win, _ = PlayTurn(&player, &boss, &magicList[0], &effects, false) // {Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
	if !win {
		t.Error()
	}
	if player.currentHealthPoint != 7 {
		t.Error()
	}
	if boss.currentHealthPoint != -1 {
		t.Error()
	}
}

func TestCase4(t *testing.T) {
	player := Player{}
	player.id = "Player"
	player.initialHealthPoint = 50
	player.initialManaStat = 500
	player.reset()

	boss := Player{}
	boss.id = "Boss"
	boss.initialHealthPoint = 58
	boss.damageStat = 9
	boss.reset()

	magicList := make([]Magic, 5)
	magicList[0] = Magic{"Magic Missile", 53, 0, 4, 0, Effect{}}
	magicList[1] = Magic{"Drain", 73, 0, 2, 2, Effect{}}
	magicList[2] = Magic{"Shield", 113, 0, 0, 0, Effect{6, 7, 0, 0, true, 0}}
	magicList[3] = Magic{"Poison", 173, 0, 0, 0, Effect{6, 0, 3, 0, false, 0}}
	magicList[4] = Magic{"Recharge", 229, 0, 0, 0, Effect{5, 0, 0, 101, false, 0}}
	sort.Sort(Magics(magicList))

	effects := make(map[string]*Effect)

	/*
		{Player 50 500 2 95 0 0 false}
		{Boss 58 0 0 0 9 0 true}
		{Poison 173 0 0 0 {6 0 3 0 false 0}}
		{Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
		{Recharge 229 0 0 0 {5 0 0 101 false 0}}
		{Shield 113 0 0 0 {6 7 0 0 true 0}}
		{Poison 173 0 0 0 {6 0 3 0 false 0}}
		{Recharge 229 0 0 0 {5 0 0 101 false 0}}
		{Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
		{Shield 113 0 0 0 {6 7 0 0 true 0}}
		{Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
		{Poison 173 0 0 0 {6 0 3 0 false 0}}
		{Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
	*/
	win, _ := PlayTurn(&player, &boss, &magicList[3], &effects, false) // {Poison 173 0 0 0 {6 0 3 0 false 0}}
	win, _ = PlayTurn(&player, &boss, &magicList[0], &effects, false)  // {Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
	win, _ = PlayTurn(&player, &boss, &magicList[4], &effects, false)  // {Recharge 229 0 0 0 {5 0 0 101 false 0}}
	win, _ = PlayTurn(&player, &boss, &magicList[2], &effects, false)  // {Shield 113 0 0 0 {6 7 0 0 true 0}}
	win, _ = PlayTurn(&player, &boss, &magicList[3], &effects, false)  // {Poison 173 0 0 0 {6 0 3 0 false 0}}
	win, _ = PlayTurn(&player, &boss, &magicList[4], &effects, false)  // {Recharge 229 0 0 0 {5 0 0 101 false 0}}
	win, _ = PlayTurn(&player, &boss, &magicList[0], &effects, false)  // {Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
	win, _ = PlayTurn(&player, &boss, &magicList[2], &effects, false)  // {Shield 113 0 0 0 {6 7 0 0 true 0}}
	win, _ = PlayTurn(&player, &boss, &magicList[0], &effects, false)  // {Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
	win, _ = PlayTurn(&player, &boss, &magicList[3], &effects, false)  // {Poison 173 0 0 0 {6 0 3 0 false 0}}
	win, _ = PlayTurn(&player, &boss, &magicList[0], &effects, false)  // {Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
	if !win {
		t.Error()
	}
	if player.currentHealthPoint != 2 {
		t.Error()
	}
	if boss.currentHealthPoint != 0 {
		t.Error()
	}
}

// func TestCaseHard1(t *testing.T) {
// 	player := Player{}
// 	player.id = "Player"
// 	player.initialHealthPoint = 50
// 	player.initialManaStat = 500
// 	player.reset()

// 	boss := Player{}
// 	boss.id = "Boss"
// 	boss.initialHealthPoint = 55
// 	boss.damageStat = 8
// 	boss.reset()

// 	magicList := make([]Magic, 5)
// 	magicList[0] = Magic{"Magic Missile", 53, 0, 4, 0, Effect{}}
// 	magicList[1] = Magic{"Drain", 73, 0, 2, 2, Effect{}}
// 	magicList[2] = Magic{"Shield", 113, 0, 0, 0, Effect{6, 7, 0, 0, true, 0}}
// 	magicList[3] = Magic{"Poison", 173, 0, 0, 0, Effect{6, 0, 3, 0, false, 0}}
// 	magicList[4] = Magic{"Recharge", 229, 0, 0, 0, Effect{5, 0, 0, 101, false, 0}}
// 	sort.Sort(Magics(magicList))

// 	effects := make(map[string]*Effect)

// 	/*
// 		{Player 50 500 2 128 2 0 false}
// 		{Boss 55 0 0 0 8 0 true}
// 		{Poison 173 0 0 0 {6 0 3 0 false 0}}
// 		{Recharge 229 0 0 0 {5 0 0 101 false 0}}
// 		{Shield 113 0 0 0 {6 7 0 0 true 0}}
// 		{Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
// 		{Poison 173 0 0 0 {6 0 3 0 false 0}}
// 		{Recharge 229 0 0 0 {5 0 0 101 false 0}}
// 		{Shield 113 0 0 0 {6 7 0 0 true 0}}
// 		{Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
// 		{Poison 173 0 0 0 {6 0 3 0 false 0}}
// 		{Drain 73 0 2 2 {0 0 0 0 false 0}}

// 	*/
// 	win, _ := PlayTurn(&player, &boss, &magicList[3], &effects, true) // {Poison 173 0 0 0 {6 0 3 0 false 0}}
// 	if win || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, _ = PlayTurn(&player, &boss, &magicList[4], &effects, true) // {Recharge 229 0 0 0 {5 0 0 101 false 0}}
// 	if win || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, _ = PlayTurn(&player, &boss, &magicList[2], &effects, true) // {Shield 113 0 0 0 {6 7 0 0 true 0}}
// 	if win || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, _ = PlayTurn(&player, &boss, &magicList[0], &effects, true) // {Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
// 	if win || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, _ = PlayTurn(&player, &boss, &magicList[3], &effects, true) // {Poison 173 0 0 0 {6 0 3 0 false 0}}
// 	if win || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, _ = PlayTurn(&player, &boss, &magicList[4], &effects, true) // {Recharge 229 0 0 0 {5 0 0 101 false 0}}
// 	if win || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, _ = PlayTurn(&player, &boss, &magicList[2], &effects, true) // {Shield 113 0 0 0 {6 7 0 0 true 0}}
// 	if win || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, _ = PlayTurn(&player, &boss, &magicList[0], &effects, true) // {Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
// 	if win || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, _ = PlayTurn(&player, &boss, &magicList[3], &effects, true) // {Poison 173 0 0 0 {6 0 3 0 false 0}}
// 	if win || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, _ = PlayTurn(&player, &boss, &magicList[1], &effects, true) // {Drain 73 0 2 2 {0 0 0 0 false 0}}
// 	if !win {
// 		t.Error()
// 	}
// 	if player.currentHealthPoint != 2 {
// 		t.Error()
// 	}
// 	if boss.currentHealthPoint != 0 {
// 		t.Error()
// 	}
// }

// func TestCaseHard2(t *testing.T) {
// 	player := Player{}
// 	player.id = "Player"
// 	player.initialHealthPoint = 50
// 	player.initialManaStat = 500
// 	player.reset()

// 	boss := Player{}
// 	boss.id = "Boss"
// 	boss.initialHealthPoint = 55
// 	boss.damageStat = 8
// 	boss.reset()

// 	magicList := make([]Magic, 5)
// 	magicList[0] = Magic{"Magic Missile", 53, 0, 4, 0, Effect{}}
// 	magicList[1] = Magic{"Drain", 73, 0, 2, 2, Effect{}}
// 	magicList[2] = Magic{"Shield", 113, 0, 0, 0, Effect{6, 7, 0, 0, true, 0}}
// 	magicList[3] = Magic{"Poison", 173, 0, 0, 0, Effect{6, 0, 3, 0, false, 0}}
// 	magicList[4] = Magic{"Recharge", 229, 0, 0, 0, Effect{5, 0, 0, 101, false, 0}}
// 	sort.Sort(Magics(magicList))

// 	effects := make(map[string]*Effect)

// 	/*
// 		{Player 50 500 2 128 2 0 false}
// 		{Boss 55 0 0 0 8 0 true}
// 		{Poison 173 0 0 0 {6 0 3 0 false 0}}
// 		{Recharge 229 0 0 0 {5 0 0 101 false 0}}
// 		{Shield 113 0 0 0 {6 7 0 0 true 0}}
// 		{Poison 173 0 0 0 {6 0 3 0 false 0}}
// 		{Recharge 229 0 0 0 {5 0 0 101 false 0}}
// 		{Shield 113 0 0 0 {6 7 0 0 true 0}}
// 		{Poison 173 0 0 0 {6 0 3 0 false 0}}
// 		{Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
// 		{Magic Missile 53 0 4 0 {0 0 0 0 false 0}}

// 	*/
// 	win, lose := PlayTurn(&player, &boss, &magicList[3], &effects, true) // {Poison 173 0 0 0 {6 0 3 0 false 0}}
// 	if win || lose || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, lose = PlayTurn(&player, &boss, &magicList[4], &effects, true) // {Recharge 229 0 0 0 {5 0 0 101 false 0}}
// 	if win || lose || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, lose = PlayTurn(&player, &boss, &magicList[2], &effects, true) // {Shield 113 0 0 0 {6 7 0 0 true 0}}
// 	if win || lose || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, lose = PlayTurn(&player, &boss, &magicList[3], &effects, true) // {Poison 173 0 0 0 {6 0 3 0 false 0}}
// 	if win || lose || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, lose = PlayTurn(&player, &boss, &magicList[4], &effects, true) // {Recharge 229 0 0 0 {5 0 0 101 false 0}}
// 	if win || lose || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, lose = PlayTurn(&player, &boss, &magicList[2], &effects, true) // {Shield 113 0 0 0 {6 7 0 0 true 0}}
// 	if win || lose || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, lose = PlayTurn(&player, &boss, &magicList[3], &effects, true) // {Poison 173 0 0 0 {6 0 3 0 false 0}}
// 	if win || lose || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, lose = PlayTurn(&player, &boss, &magicList[0], &effects, true) // {Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
// 	if win || lose || player.dead || boss.dead {
// 		t.Error()
// 	}
// 	win, lose = PlayTurn(&player, &boss, &magicList[0], &effects, true) // {Magic Missile 53 0 4 0 {0 0 0 0 false 0}}
// 	if !win || lose {
// 		t.Error()
// 	}
// 	if player.currentHealthPoint != 11 {
// 		t.Error()
// 	}
// 	if boss.currentHealthPoint != -1 {
// 		t.Error()
// 	}
// }
