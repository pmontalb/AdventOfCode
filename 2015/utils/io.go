package utils

import (
	"bufio"
	"os"
	"strconv"
)

// GetLines returns a vector of lines from the given file
// return nil if the file doesn't exist or cannot be opened
func GetLines(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil
	}

	return lines
}

// GetFirstLine returns the first line of a file
func GetFirstLine(filePath string) string {
	return GetLines(filePath)[0]
}

// WriteLines writes the lines vector into the given filePath
// NB: override the file if it exists
func WriteLines(lines []string, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		writer.WriteString(line + "\n")
	}

	writer.Flush()
}

// WriteIntegerOutput writes the integer into "output_${outputIdx}"
func WriteIntegerOutput(out int, outputIdx string) {
	file, err := os.Create("output" + outputIdx)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(strconv.Itoa(out) + "\n")
	writer.Flush()
}

// WriteSingleStringOutput writes the string into "output_${outputIdx}"
func WriteSingleStringOutput(out string, outputIdx string) {
	file, err := os.Create("output" + outputIdx)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(out + "\n")
	writer.Flush()
}

// ReadUntilLast returns str[:-1]
func ReadUntilLast(str string) string {
	return str[:len(str)-1]
}
