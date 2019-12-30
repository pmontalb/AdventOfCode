package main

import (
	"fmt"
	"strconv"
	"strings"
	"../../../utils"
)

// SignalCombiner stores the operation between signal
type SignalCombiner func(uint16, uint16) uint16

// Signal is a helper class that stores the operations to do at runtime
type Signal struct {
	inputSignal1    *Signal
	inputSignal2    *Signal
	operation       SignalCombiner
	value           uint16
	id              string
	alreadyComputed bool
}

// GetValue recursively computes the signal value
func (signal *Signal) GetValue() uint16 {
	if signal.alreadyComputed {
		return signal.value
	}

	if signal.inputSignal1 == nil && signal.inputSignal2 == nil {
		_, err := strconv.ParseInt(signal.id, 10, 64)
		if err != nil {
			panic(signal)
		}
		return signal.value
	}

	if signal.inputSignal2 == nil {
		return signal.operation(signal.inputSignal1.GetValue(), 0)
	}

	signal.alreadyComputed = true
	ret := signal.operation(signal.inputSignal1.GetValue(), signal.inputSignal2.GetValue())
	signal.value = ret

	return signal.value
}

func getOrCreateSymbolicSignal(id string, signalMap *map[string](*Signal)) *Signal {
	var signal *Signal
	if _, found := (*signalMap)[id]; !found {
		signal = new(Signal)
		signal.id = id
	} else {
		signal = (*signalMap)[id]
	}
	(*signalMap)[id] = signal
	return signal
}

func getOrCreateNumericalSignal(id string, value uint16, signalMap *map[string](*Signal)) *Signal {
	var signal *Signal = getOrCreateSymbolicSignal(id, signalMap)
	signal.inputSignal1 = nil
	signal.inputSignal2 = nil
	signal.value = value

	return signal
}

func getOrCreateSignal(id string, signalMap *map[string](*Signal)) *Signal {
	signalValue, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		return getOrCreateNumericalSignal(id, uint16(signalValue), signalMap)
	}

	return getOrCreateSymbolicSignal(id, signalMap)
}

func parseSignal(line string, signalMap *map[string](*Signal), debug int) {
	tokens := strings.Split(line, " ")
	if tokens[len(tokens)-2] != "->" {
		panic(tokens)
	}

	var signal *Signal = getOrCreateSignal(tokens[len(tokens)-1], signalMap)
	if debug > 0 {
		fmt.Println("Processing signal", signal.id)
	}
	if len(tokens) == 3 {
		// signal assignment
		signal.inputSignal1 = getOrCreateSignal(tokens[0], signalMap)
		signal.inputSignal2 = nil
		signal.operation = func(signal1 uint16, signal2 uint16) uint16 { return signal1 }
		if debug > 1 {
			fmt.Printf("Signal(%s) assigning value(%d)\n", signal.id, signal.value)
		}

	} else if len(tokens) == 5 {
		// binary operation
		signal.inputSignal1 = getOrCreateSignal(tokens[0], signalMap)
		signal.inputSignal2 = getOrCreateSignal(tokens[2], signalMap)
		switch tokens[1] {
		case "AND":
			signal.operation = func(signal1 uint16, signal2 uint16) uint16 { return signal1 & signal2 }
			if debug > 2 {
				fmt.Printf("Signal(%s) (%s|%d) & (%s|%d) -> %d\n", signal.id,
					signal.inputSignal1.id, signal.inputSignal1.GetValue(),
					signal.inputSignal2.id, signal.inputSignal2.GetValue(),
					signal.GetValue())
			}
		case "OR":
			signal.operation = func(signal1 uint16, signal2 uint16) uint16 { return signal1 | signal2 }
			if debug > 2 {
				fmt.Printf("Signal(%s) (%s|%d) | (%s|%d) -> %d\n", signal.id,
					signal.inputSignal1.id, signal.inputSignal1.GetValue(),
					signal.inputSignal2.id, signal.inputSignal2.GetValue(),
					signal.GetValue())
			}
		case "LSHIFT":
			signal.operation = func(signal1 uint16, signal2 uint16) uint16 { return signal1 << signal2 }
			if debug > 2 {
				fmt.Printf("Signal(%s) (%s|%d) << (%s|%d) -> %d\n", signal.id,
					signal.inputSignal1.id, signal.inputSignal1.GetValue(),
					signal.inputSignal2.id, signal.inputSignal2.GetValue(),
					signal.GetValue())
			}
		case "RSHIFT":
			signal.operation = func(signal1 uint16, signal2 uint16) uint16 { return signal1 >> signal2 }
			if debug > 2 {
				fmt.Printf("Signal(%s) (%s|%d) >> (%s|%d) -> %d\n", signal.id,
					signal.inputSignal1.id, signal.inputSignal1.GetValue(),
					signal.inputSignal2.id, signal.inputSignal2.GetValue(),
					signal.GetValue())
			}
		default:
			panic(tokens)
		}
	} else if len(tokens) == 4 {
		// binary operation
		signal.inputSignal1 = getOrCreateSymbolicSignal(tokens[1], signalMap)
		switch tokens[0] {
		case "NOT":
			signal.inputSignal2 = nil
			signal.operation = func(signal1 uint16, signal2 uint16) uint16 { return ^signal1 }
			if debug > 2 {
				fmt.Printf("Signal(%s) ~(%s|%d) -> %d\n", signal.id,
					signal.inputSignal1.id, signal.inputSignal1.GetValue(),
					signal.GetValue())
			}
		default:
			panic(tokens)
		}
	}
}

func main() {
	signalMap := make(map[string](*Signal))
	lines := utils.GetLines("input")
	for _, line := range lines {
		parseSignal(line, &signalMap, 0)
	}
	fmt.Println("a=", signalMap["a"].GetValue())
	utils.WriteLines([]string{strconv.Itoa(int(signalMap["a"].GetValue()))}, "output_1")

	// reset signals
	signalMap["b"].value = signalMap["a"].GetValue()
	signalMap["b"].alreadyComputed = true

	for id, signal := range signalMap {
		if id == "b" {
			continue
		}
		signal.alreadyComputed = false
	}

	fmt.Println("a=", signalMap["a"].GetValue())
	utils.WriteLines([]string{strconv.Itoa(int(signalMap["a"].GetValue()))}, "output_2")
}
