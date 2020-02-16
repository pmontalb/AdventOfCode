package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"../utils"
)

func main() {
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
	magicList[0] = Magic{"Magic Missile", 53, 0, 4, 0, Effect{}}
	magicList[1] = Magic{"Drain", 73, 0, 2, 2, Effect{}}
	magicList[2] = Magic{"Shield", 113, 0, 0, 0, Effect{6, 7, 0, 0, true, 0}}
	magicList[3] = Magic{"Poison", 173, 0, 0, 0, Effect{6, 0, 3, 0, false, 0}}
	magicList[4] = Magic{"Recharge", 229, 0, 0, 0, Effect{5, 0, 0, 101, false, 0}}
	sort.Sort(Magics(magicList))

	minMana := Optimize(&player, &boss, magicList, false)
	fmt.Println(minMana)
	utils.WriteIntegerOutput(minMana, "1")

	player.reset()
	boss.reset()
	minManaHard := Optimize(&player, &boss, magicList, true)
	fmt.Println(minManaHard)
	utils.WriteIntegerOutput(minManaHard, "2")
}
