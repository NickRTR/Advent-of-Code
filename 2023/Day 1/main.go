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

func part2(lines []string) int {
	dictionary := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}

	formattedLines := []string{}

	for _, line := range lines {
		word := ""
		for _, char := range line {
			word += string(char)
			for key, value := range dictionary {
				if strings.Contains(word, key) {
					word = regexp.MustCompile(`[^0-9]`).ReplaceAllString(word, "")
					word += value
				}
			}
		}
		formattedLines = append(formattedLines, word)
	}

	return part1(formattedLines)
}

func main() {
	file := readInput(("input"))

	lines := strings.Split(file, "\n")

	println("The sum of all calibration values is:", part1(lines))
	println("The sum of all calibration values with spelled out numbers is:", part2(lines))
}
