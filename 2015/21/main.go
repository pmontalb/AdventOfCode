package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"../../../utils"
)

type Equipment struct {
	id           string
	equipType    string
	attackBoost  int
	defenceBoost int

	cost int
}

type Player struct {
	id string

	initialHealthPoint int
	initialAttackStat  int
	initialDefenceStat int

	currentHealthPoint int
	currentAttackStat  int
	currentDefenseStat int

	dead bool

	equipmentCost int
	equipment     map[string]*Equipment
}

func (p *Player) attack(q *Player) {
	damage := p.currentAttackStat - q.currentDefenseStat
	if damage < 1 {
		damage = 1
	}
	q.currentHealthPoint -= damage
	q.dead = q.currentHealthPoint <= 0
	//fmt.Printf("player(%s|atk=%d) attacking player(%s|def=%d) dealing (%d) damage: player(%s) hp(%d -> %d) dead(%v)\n", p.id, p.currentAttackStat, q.id, q.currentDefenseStat, damage, q.id, q.currentHealthPoint+damage, q.currentHealthPoint, q.dead)
}

func (p *Player) equip() {
	p.currentAttackStat = p.initialAttackStat
	p.currentDefenseStat = p.initialDefenceStat
	p.equipmentCost = 0
	for _, e := range p.equipment {
		p.currentAttackStat += e.attackBoost
		p.currentDefenseStat += e.defenceBoost
		p.equipmentCost += e.cost
	}
}

func (p *Player) addEquipment(e *Equipment) {
	//fmt.Printf("player(%s) equipping(%s|%s)\n", p.id, e.equipType, e.id)
	p.equipment[e.equipType] = e
}
func (p *Player) unEquip(e *Equipment) {
	delete(p.equipment, e.id)
}

func (p *Player) reset() {
	p.equipment = make(map[string]*Equipment)
	p.currentAttackStat = p.initialAttackStat
	p.currentDefenseStat = p.initialDefenceStat
	p.resetHealthPoint()
}
func (p *Player) resetHealthPoint() {
	p.currentHealthPoint = p.initialHealthPoint
	p.dead = false
}

func playerOneWins(p1 *Player, p2 *Player) bool {
	nTurns := 0
	for {
		nTurns++
		//fmt.Printf("turn #%d:\n", nTurns)
		p1.attack(p2)
		if p2.dead {
			//fmt.Printf("player(%s) dies after %d turns [other player(%s) remaining hp(%d)]\n", p2.id, nTurns, p1.id, p1.currentHealthPoint)
			return true
		}

		p2.attack(p1)
		if p1.dead {
			//fmt.Printf("player(%s) dies after %d turns [other player(%s) remaining hp(%d)]\n", p1.id, nTurns, p2.id, p2.currentHealthPoint)
			return false
		}
	}
}

func test() {
	p1 := Player{"p1", 8, 5, 5, 8, 5, 5, false, 0, nil}
	p2 := Player{"p2", 12, 7, 2, 12, 7, 2, false, 0, nil}

	p1.reset()
	p2.reset()

	if !playerOneWins(&p1, &p2) {
		panic(0)
	}

	weapons := []Equipment{
		{"Dagger", "Weapon", 4, 0, 8},
		{"Shortsword", "Weapon", 5, 0, 10},
		{"Warhammer", "Weapon", 6, 0, 25},
		{"Longsword", "Weapon", 7, 0, 40},
		{"Greataxe", "Weapon", 8, 0, 74},
	}

	armors := []Equipment{
		{"Leather", "Armor", 0, 1, 13},
		{"Chainmail", "Armor", 0, 2, 31},
		{"Splintmail", "Armor", 0, 3, 53},
		{"Bandedmail", "Armor", 0, 4, 75},
		{"Platemail", "Armor", 0, 5, 102},
	}

	rings := []Equipment{
		{"Damage1", "Ring", 1, 0, 25},
		{"Damage2", "Ring", 2, 0, 50},
		{"Damage3", "Ring", 3, 0, 100},
		{"Defence1", "Ring", 0, 1, 20},
		{"Defence2", "Ring", 0, 2, 40},
		{"Defence3", "Ring", 0, 3, 80},
	}

	p1 = Player{"p1", 100, 0, 0, 0, 0, 0, false, 0, nil}
	p2 = Player{"p2", 109, 8, 2, 109, 8, 2, false, 0, nil}
	p1.reset()
	p2.reset()

	p1.addEquipment(&weapons[0])
	p1.addEquipment(&armors[0])

	rings[0].equipType = "Ring1"
	p1.addEquipment(&rings[0])
	rings[1].equipType = "Ring2"
	p1.addEquipment(&rings[1])

	p1.equip()

	if playerOneWins(&p1, &p2) {
		panic(0)
	}
	if p2.currentHealthPoint != 34 {
		panic(0)
	}

	p1.reset()
	p2.reset()

	p1.addEquipment(&weapons[0])
	p1.addEquipment(&armors[0])
	rings[1].equipType = "Ring1"
	p1.addEquipment(&rings[1])
	rings[2].equipType = "Ring2"
	p1.addEquipment(&rings[2])

	p1.equip()

	if playerOneWins(&p1, &p2) {
		panic(0)
	}
	if p2.currentHealthPoint != 4 {
		panic(0)
	}

}

func gridSearch(weapons *[]Equipment, armors *[]Equipment, rings *[]Equipment, p1 *Player, p2 *Player, minimize bool) int {

	optimalCost := math.MaxInt64
	if !minimize {
		optimalCost = 0
	}
	for _, weapon := range *weapons {
		p1.reset()
		p2.reset()

		p1.addEquipment(&weapon)
		for _, armor := range *armors {
			p1.addEquipment(&armor)
			for i, ring1 := range *rings {
				ring1.equipType = "Ring1"
				p1.addEquipment(&ring1)
				for j, ring2 := range *rings {
					if i == j {
						continue
					}
					ring2.equipType = "Ring2"
					p1.addEquipment(&ring2)

					p1.equip()

					p1Wins := playerOneWins(p1, p2)
					p2.reset()
					p1.resetHealthPoint()

					if minimize {
						if !p1Wins {
							continue
						}

						if p1.equipmentCost < optimalCost {
							optimalCost = p1.equipmentCost
							//fmt.Printf("*NEW optimalCost(%d)\n", optimalCost)
						} else {
							//fmt.Printf("currentObj(%d) optimalCost(%d)\n", p1.equipmentCost, optimalCost)
						}
					} else {
						if p1Wins {
							continue
						}

						if p1.equipmentCost > optimalCost {
							optimalCost = p1.equipmentCost
							//fmt.Printf("*NEW optimalCost(%d){%v}\n", optimalCost, eq.print())
						} else {
							//fmt.Printf("currentObj(%d) optimalCost(%d){%v}\n", p1.equipmentCost, optimalCost, eq.print())
						}
					}
				}
			}
		}
	}

	return optimalCost
}

func main() {
	test()

	weapons := []Equipment{
		{"Dagger", "Weapon", 4, 0, 8},
		{"Shortsword", "Weapon", 5, 0, 10},
		{"Warhammer", "Weapon", 6, 0, 25},
		{"Longsword", "Weapon", 7, 0, 40},
		{"Greataxe", "Weapon", 8, 0, 74},
	}

	armors := []Equipment{
		{"None", "Armor", 0, 0, 0},
		{"Leather", "Armor", 0, 1, 13},
		{"Chainmail", "Armor", 0, 2, 31},
		{"Splintmail", "Armor", 0, 3, 53},
		{"Bandedmail", "Armor", 0, 4, 75},
		{"Platemail", "Armor", 0, 5, 102},
	}

	rings := []Equipment{
		{"None", "Ring", 0, 0, 0},
		{"Damage1", "Ring", 1, 0, 25},
		{"Damage2", "Ring", 2, 0, 50},
		{"Damage3", "Ring", 3, 0, 100},
		{"Defence1", "Ring", 0, 1, 20},
		{"Defence2", "Ring", 0, 2, 40},
		{"Defence3", "Ring", 0, 3, 80},
	}

	p1 := Player{"p1", 100, 0, 0, 0, 0, 0, false, 0, nil}
	p1.reset()

	lines := utils.GetLines("input")
	if len(lines) != 3 {
		panic(lines)
	}
	p2 := Player{}
	p2.id = "p2"
	p2.initialHealthPoint, _ = strconv.Atoi(strings.TrimSpace(strings.Split(lines[0], ":")[1]))
	p2.initialAttackStat, _ = strconv.Atoi(strings.TrimSpace(strings.Split(lines[1], ":")[1]))
	p2.initialDefenceStat, _ = strconv.Atoi(strings.TrimSpace(strings.Split(lines[2], ":")[1]))
	p2.reset()

	minCost := gridSearch(&weapons, &armors, &rings, &p1, &p2, true)
	fmt.Println(minCost)
	utils.WriteIntegerOutput(minCost, "1")

	maxCost := gridSearch(&weapons, &armors, &rings, &p1, &p2, false)
	fmt.Println(maxCost)
	utils.WriteIntegerOutput(maxCost, "2")
}
