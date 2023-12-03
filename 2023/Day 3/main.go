// Exercise: https://adventofcode.com/2023/day/2

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

func validateNumber(grid [140][]string, y int, x int, number string) int {
	end := x - 1
	start := end - len(number) + 1
	middle := len(number)/2 + start

	coordinates := [][]int{
		{y, start - 1},
		{y, end + 1},

		{y + 1, start - 1},
		{y + 1, start},
		{y + 1, middle},
		{y + 1, end},
		{y + 1, end + 1},

		{y - 1, start - 1},
		{y - 1, start},
		{y - 1, middle},
		{y - 1, end},
		{y - 1, end + 1},
	}

	for _, coordinate := range coordinates {
		if coordinate[0] >= 0 && coordinate[0] < len(grid) && coordinate[1] >= 0 && coordinate[1] < len(grid[0]) {
			isNumber, _ := regexp.MatchString(`[0-9]`, grid[coordinate[0]][coordinate[1]])
			if grid[coordinate[0]][coordinate[1]] != "." && !isNumber {
				convertedNumber, _ := strconv.Atoi(number)
				return convertedNumber
			}
		}
	}
	return 0
}

func part1(grid [140][]string) int {
	partNumberSum := 0

	for y, row := range grid {
		number := ""
		for x, character := range row {
			isNumber, _ := regexp.MatchString(`[0-9]`, character)
			if isNumber {
				number += grid[y][x]
			} else if number != "" || x == len(row) {
				// search for symbols around
				partNumberSum += validateNumber(grid, y, x, number)
				number = ""
			}
		}
		partNumberSum += validateNumber(grid, y, len(row), number)
	}

	return partNumberSum
}

func main() {
	file := readInput(("input"))

	lines := strings.Split(file, "\n")

	var grid [140][]string

	for y, line := range lines {
		characters := strings.Split(line, "")
		grid[y] = characters
	}

	println("Part 1: The sum of the part numbers of the engine schematic is:", part1(grid))
}
