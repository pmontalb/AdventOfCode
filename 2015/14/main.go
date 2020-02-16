package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"../utils"
)

type Reindeer struct {
	id                   string
	speed                int
	flyingTime           int
	restingTime          int
	position             int
	score                int
	remainingFlyingTime  int
	remainingRestingTime int
}

func (r *Reindeer) Parse(line string) {
	tokens := strings.Split(line, " ")
	r.id = tokens[0]
	r.speed, _ = strconv.Atoi(tokens[3])
	r.flyingTime, _ = strconv.Atoi(tokens[6])
	r.restingTime, _ = strconv.Atoi(tokens[13])

	r.remainingFlyingTime = r.flyingTime
	r.remainingRestingTime = r.restingTime
}

func (r *Reindeer) Advance(dt int) {
	r.position += r.speed * dt
	r.remainingFlyingTime -= dt
}
func (r *Reindeer) Rest(dt int) {
	r.remainingRestingTime -= dt
}
func (r *Reindeer) Restart() {
	r.remainingFlyingTime = r.flyingTime
	r.remainingRestingTime = r.restingTime
}

type Reindeers []Reindeer

func (reindeers Reindeers) Len() int           { return len(reindeers) }
func (reindeers Reindeers) Swap(i, j int)      { reindeers[i], reindeers[j] = reindeers[j], reindeers[i] }
func (reindeers Reindeers) Less(i, j int) bool { return reindeers[i].position < reindeers[j].position }

type ReindeerScore []Reindeer

func (reindeers ReindeerScore) Len() int           { return len(reindeers) }
func (reindeers ReindeerScore) Swap(i, j int)      { reindeers[i], reindeers[j] = reindeers[j], reindeers[i] }
func (reindeers ReindeerScore) Less(i, j int) bool { return reindeers[i].score < reindeers[j].score }

func updateReindeer(reindeer *Reindeer) {
	const dt = 1
	if reindeer.remainingFlyingTime > 0 {
		reindeer.Advance(dt)
	} else {
		if reindeer.remainingRestingTime > 0 {
			reindeer.Rest(dt)
		} else {
			reindeer.Restart()

			reindeer.Advance(dt)
		}
	}
}

func evaluate(finalTime int, reindeers *[]Reindeer) {
	for time := 1; time <= finalTime; time++ {
		for i := 0; i < len(*reindeers); i++ {
			// r := (*reindeers)[i]
			updateReindeer(&((*reindeers)[i]))
			// fmt.Printf("t(%d) r(%s) p(%d -> %d) ft(%d -> %d) rt(%d -> %d)\n", time, r.id,
			// 	r.position, (*reindeers)[i].position,
			// 	r.remainingFlyingTime, (*reindeers)[i].remainingFlyingTime,
			// 	r.remainingRestingTime, (*reindeers)[i].remainingRestingTime)
		}
		sort.Sort(sort.Reverse(Reindeers(*reindeers)))
		(*reindeers)[0].score++
	}
}

func main() {
	lines := utils.GetLines("input")

	var reindeers []Reindeer
	for _, line := range lines {
		var r Reindeer
		r.Parse(line)
		reindeers = append(reindeers, r)
	}
	//fmt.Println(reindeers)
	evaluate(2503, &reindeers)

	sort.Sort(sort.Reverse(Reindeers(reindeers)))
	fmt.Println(reindeers[0].position)
	utils.WriteIntegerOutput(reindeers[0].position, "1")

	sort.Sort(sort.Reverse(ReindeerScore(reindeers)))
	fmt.Println(reindeers[0].score)
	utils.WriteIntegerOutput(reindeers[0].score, "2")
}
