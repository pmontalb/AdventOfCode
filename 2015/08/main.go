package main

import "../utils"

import "fmt"

import "strconv"

func countStringData(str string) int {
	stringDataCount := 0
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "\"" {
			// fmt.Println("excluding", str[i:i+1])
			continue
		} else if string(str[i]) == "\\" {
			switch string(str[i+1]) {
			case "\"":
				// fmt.Println("excluding", string(str[i]), "but including", string(str[i+1]))
				i++
				stringDataCount++
			case "\\":
				// fmt.Println("excluding", string(str[i]), "but including", string(str[i+1]))
				i++
				stringDataCount++
			case "x":
				// fmt.Println("excluding", str[i:i+4], "but including one char")
				i += 3
				stringDataCount++
			}
			continue
		}
		// fmt.Println("including", str[i:i+1])
		stringDataCount++
	}
	if len(str) <= stringDataCount {
		panic(str)
	}
	return stringDataCount
}

func countEncodedStringData(str string) int {
	encodedStringDataCount := 2 // two initial quotes have to be kept
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "\"" {
			encodedStringDataCount += 2 // " -> \"
			continue
		} else if string(str[i]) == "\\" {
			encodedStringDataCount += 2 // \ -> \\
			continue
		}
		encodedStringDataCount++
	}
	if len(str) >= encodedStringDataCount {
		panic(str)
	}
	return encodedStringDataCount
}

func main() {
	lines := utils.GetLines("input")

	totalCodeCount := 0
	totalStringDataCount := 0
	totalEncodedStringDataCount := 0
	for _, line := range lines {
		fmt.Println(line)
		totalCodeCount += len(line)
		totalStringDataCount += countStringData(line)
		totalEncodedStringDataCount += countEncodedStringData(line)
	}
	fmt.Println(totalCodeCount - totalStringDataCount)
	utils.WriteLines([]string{strconv.Itoa(totalCodeCount - totalStringDataCount)}, "output_1")

	fmt.Println(totalEncodedStringDataCount - totalCodeCount)
	utils.WriteLines([]string{strconv.Itoa(totalEncodedStringDataCount - totalCodeCount)}, "output_2")
}
