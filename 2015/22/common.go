package main

import "math"

type Effect struct {
	maxTurns int

	armor        int
	damage       int
	manaRecovery int

	immediate      bool
	remainingTurns int
}

type Magic struct {
	id       string
	manaCost int

	armor  int
	damage int

	healthRecovery int
	effect         Effect
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

	p.armorStat = m.armor
	p.currentHealthPoint += m.healthRecovery
	p.currentManaStat -= m.manaCost
	//fmt.Printf("player(%s|atk=%d) attacking player(%s|def=%d) dealing (%d) damage: player(%s) hp(%d -> %d) dead(%v)\n", p.id, p.currentAttackStat, q.id, q.currentDefenseStat, damage, q.id, q.currentHealthPoint+damage, q.currentHealthPoint, q.dead)
}

type State struct {
	player     Player
	boss       Player
	effects    map[string]Effect
	magicChain []Magic
}

type StateMagicPair struct {
	state State
	magic Magic
}

func applyEffect(p1 *Player, p2 *Player, e *Effect) bool {
	if e.remainingTurns > 0 {
		e.remainingTurns--

		if p1.armorStat == 0 { // TODO: this is just a hack, it doesn't allow for multiple shield-like effects
			p1.armorStat += e.armor
		}
		p1.currentManaStat += e.manaRecovery

		p2.takeDamage(e.damage)

		return e.remainingTurns != 0
	}
	return true
}

func wearOff(p *Player, e *Effect) {
	p.armorStat -= (*e).armor
}

func playerOneTurn(p1 *Player, p2 *Player, magic *Magic, effects *map[string]*Effect) (bool, bool) {
	//fmt.Printf("* START p1(%s|hp=%d) p2(%s|hp=%d) effects(%v)\n", p1.id, p1.currentHealthPoint, p2.id, p2.currentHealthPoint, *effects)

	canUseEffect := true
	_, found := (*effects)[magic.id]
	illegalMove := found && (*effects)[magic.id].remainingTurns > 1 // cannot spell a magic which is not about to finish this turn
	if illegalMove {
		return false, true
	}

	if magic.effect.maxTurns > 0 {
		tmp := magic.effect
		canUseEffect = tmp.immediate || (found && (*effects)[magic.id].remainingTurns == 1)

		(*effects)[magic.id] = &tmp
		tmp.remainingTurns = tmp.maxTurns
	}

	for id, effect := range *effects {
		if canUseEffect || id != magic.id {
			if !applyEffect(p1, p2, effect) {
				delete(*effects, id)
				wearOff(p1, effect)
			}

			if canUseEffect && id == magic.id {
				effect.remainingTurns++ // when immediate, the effect remaining turns drops by one only at the next turn!
			}
		}
	}
	if p2.dead {
		return true, false
	}

	if p1.currentManaStat < magic.manaCost {
		return false, true
	}

	//fmt.Printf("** SETUP p1(%s|hp=%d|mgk=%s|atk=%d|def=%d) effects(%v)\n", p1.id, p1.currentHealthPoint, magic.id, p1.damageStat, p1.armorStat, *effects)
	p1.castSpell(magic, p2)
	//fmt.Printf("** AFTER p1(%s|hp=%d|mgk=%s|atk=%d|def=%d|mn=%d) p2(%s|hp=%d) effects(%v)\n", p1.id, p1.currentHealthPoint, magic.id, p1.damageStat, p1.armorStat, p1.currentManaStat, p2.id, p2.currentHealthPoint, *effects)

	if p2.dead {
		return true, false
	}
	return false, false
}

func playerTwoTurn(p1 *Player, p2 *Player, magic *Magic, effects *map[string]*Effect) (bool, bool) {
	for id, effect := range *effects {
		if !applyEffect(p1, p2, effect) {
			delete(*effects, id)
			wearOff(p1, effect)
		}
	}

	if p2.dead {
		return true, false
	}

	p2.attack(p1)
	//fmt.Printf("*** END p1(%s|hp=%d|mn=%d) p2(%s|hp=%d) effects(%v)\n", p1.id, p1.currentHealthPoint, p1.currentManaStat, p2.id, p2.currentHealthPoint, *effects)
	return false, p1.dead
}

func takeDamageOnHardMode(p1 *Player, hardMode bool) bool {
	if hardMode {
		p1.takeDamage(1)
	}
	return !p1.dead
}

func PlayTurn(p1 *Player, p2 *Player, magic *Magic, effects *map[string]*Effect, hardMode bool) (bool, bool) {
	if !takeDamageOnHardMode(p1, hardMode) {
		return false, true
	}

	playerOneWins, playerOneLose := playerOneTurn(p1, p2, magic, effects)
	if playerOneWins || playerOneLose {
		return playerOneWins, playerOneLose
	}

	// if !takeDamageOnHardMode(p1, hardMode) {
	// 	return false, true
	// }

	playerOneWins, playerOneLose = playerTwoTurn(p1, p2, magic, effects)
	if playerOneWins || playerOneLose {
		return playerOneWins, playerOneLose
	}

	return false, false
}

func Optimize(player *Player, boss *Player, magicList []Magic, hardMode bool) int {

	originalState := State{*player, *boss, nil, nil}
	originalState.effects = make(map[string]Effect)

	var currentState StateMagicPair
	var states []StateMagicPair
	for _, magic := range magicList {
		tmpState := State{}
		tmpState.player = originalState.player
		tmpState.boss = originalState.boss
		tmpState.effects = originalState.effects
		states = append(states, StateMagicPair{tmpState, magic})
	}

	minCost := math.MaxInt64
	for len(states) > 0 {
		currentState, states = states[0], states[1:]

		newEffects := make(map[string]*Effect)
		for id, effect := range currentState.state.effects {
			tmp := effect
			newEffects[id] = &tmp
		}
		newPlayer := currentState.state.player
		newBoss := currentState.state.boss
		newMagic := currentState.magic

		win, lost := PlayTurn(&newPlayer, &newBoss, &newMagic, &newEffects, hardMode)

		totalManaCost := 0
		for _, m := range currentState.state.magicChain {
			//fmt.Println(m)
			totalManaCost += m.manaCost
		}
		//fmt.Println(currentState.magic)
		totalManaCost += currentState.magic.manaCost

		if win {
			if newPlayer.dead {
				panic(0)
			}

			if totalManaCost < minCost {
				minCost = totalManaCost
			}
		}

		// stop branching from this node if:
		/*
			- won/lost
			- already found an optimum and new cost is higher than the optimum
		*/
		if lost || win || totalManaCost > minCost {
			continue
		}

		newState := State{newPlayer, newBoss, nil, nil}

		for _, m := range currentState.state.magicChain {
			newState.magicChain = append(newState.magicChain, m)
		}
		newState.magicChain = append(newState.magicChain, currentState.magic)

		newState.effects = make(map[string]Effect)
		for id, effect := range newEffects {
			newState.effects[id] = *effect
		}

		for _, magic := range magicList {
			states = append(states, StateMagicPair{newState, magic})
		}
	}
	return minCost
}
