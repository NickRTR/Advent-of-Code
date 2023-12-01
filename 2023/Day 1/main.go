// Exercise: https://adventofcode.com/2023/day/1

package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput(file string) string {
	rawFile, err := os.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func part1(lines []string) int {
	filteredLines := []string{}

	for _, line := range lines {
		filteredLines = append(filteredLines, regexp.MustCompile(`[^0-9]`).ReplaceAllString(line, ""))
	}

	result := 0

	for _, line := range filteredLines {
		runes := []rune(line)
		characters := string(runes[0]) + string(runes[len(runes)-1])
		value, _ := strconv.Atoi(characters)
		result += value
	}

	return result
}

func main() {
	file := readInput(("input"))

	lines := strings.Split(file, "\n")

	println("The sum of all calibration values is:", part1(lines))
}
