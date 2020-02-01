package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../../utils"
)

type Instruction struct {
	id             string
	value          int
	targetRegister *int

	computer *Computer
	apply    func() bool
}
type Computer struct {
	registers      []int
	instructionSet []Instruction

	currentInstructionIndex int
}

func (i *Instruction) parse(str string) {
	assignTargetRegister := func(s string) {
		switch strings.TrimSpace(s) {
		case "a":
			i.targetRegister = &i.computer.registers[0]
		case "b":
			i.targetRegister = &i.computer.registers[1]
		default:
			panic(s)
		}
	}
	tokens := strings.Split(str, " ")
	i.id = tokens[0]

	switch tokens[0] {
	case "hlf":
		if len(tokens) != 2 {
			panic(0)
		}
		assignTargetRegister(tokens[1])
		i.apply = func() bool {
			*i.targetRegister /= 2
			return true
		}
	case "tpl":
		if len(tokens) != 2 {
			panic(0)
		}
		assignTargetRegister(tokens[1])
		i.apply = func() bool {
			*i.targetRegister *= 3
			return true
		}
	case "inc":
		if len(tokens) != 2 {
			panic(0)
		}
		assignTargetRegister(tokens[1])
		i.apply = func() bool {
			*i.targetRegister++
			return true
		}
	case "jmp":
		i.apply = func() bool {
			i.computer.currentInstructionIndex += i.value
			return false
		}
		value, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}
		i.value = value
		if i.value == 0 {
			panic(0) // infinite loop
		}
	case "jie":
		if len(tokens) != 3 {
			panic(0)
		}
		assignTargetRegister(tokens[1][:len(tokens[1])-1])
		i.apply = func() bool {
			*i.targetRegister++
			return true
		}

		i.apply = func() bool {
			if *i.targetRegister%2 == 0 {
				if i.value == 0 {
					panic(0) // infinite loop
				}
				i.computer.currentInstructionIndex += i.value
				return false
			}
			return true
		}
		value, err := strconv.Atoi(tokens[2])
		if err != nil {
			panic(err)
		}
		i.value = value
	case "jio":
		if len(tokens) != 3 {
			panic(0)
		}
		assignTargetRegister(tokens[1][:len(tokens[1])-1])
		i.apply = func() bool {
			if *i.targetRegister == 1 {
				if i.value == 0 {
					panic(0) // infinite loop
				}
				i.computer.currentInstructionIndex += i.value
				return false
			}
			return true
		}
		value, err := strconv.Atoi(tokens[2])
		if err != nil {
			panic(err)
		}
		i.value = value
	default:
		panic(str)
	}
}

func (i *Instruction) execute() bool {
	return i.apply()
}

func (c *Computer) executeNext() bool {
	if c.currentInstructionIndex >= len(c.instructionSet) {
		return false
	}
	for true {
		if c.currentInstructionIndex >= len(c.instructionSet) {
			//fmt.Printf("idx(%d) cannot run as it exceeds instruction set nubmer(%d)\n", c.currentInstructionIndex, len(c.instructionSet))
			break
		}
		//fmt.Printf("idx(%d) running (%v) reg(%v)\n", c.currentInstructionIndex, c.instructionSet[c.currentInstructionIndex], c.registers)
		if c.instructionSet[c.currentInstructionIndex].execute() {
			break
		}
	}
	c.currentInstructionIndex++
	return true
}

func (c *Computer) reset() {
	for i := 0; i < len(c.registers); i++ {
		c.registers[i] = 0
	}
	c.currentInstructionIndex = 0
}

func main() {
	lines := utils.GetLines("input")

	computer := Computer{}
	computer.registers = make([]int, 2)

	for _, line := range lines {
		instr := Instruction{}
		instr.computer = &computer

		instr.parse(line)

		computer.instructionSet = append(computer.instructionSet, instr)
	}

	for true {
		if !computer.executeNext() {
			break
		}
	}
	fmt.Println(computer.registers[1])
	utils.WriteIntegerOutput(computer.registers[1], "1")

	computer.reset()
	computer.registers[0] = 1

	for true {
		if !computer.executeNext() {
			break
		}
	}
	fmt.Println(computer.registers[1])
	utils.WriteIntegerOutput(computer.registers[1], "2")
}
