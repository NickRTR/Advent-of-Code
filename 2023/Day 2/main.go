// Exercise: https://adventofcode.com/2023/day/2

package main

import (
	"log"
	"math"
	"os"
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

func getCount(line string) map[string]int {
	values := strings.Split(line, ":")
	count := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	sets := strings.Split(values[1], ";")
	for _, set := range sets {
		items := strings.Split(set, ",")
		for _, item := range items {
			value := strings.Split(item, " ")
			number, _ := strconv.ParseFloat(value[1], 64)
			count[value[2]] = int(math.Max(float64(count[value[2]]), number))
		}
	}
	return count
}

func part1(lines []string) int {
	possibleGames := 0

	for index, line := range lines {
		count := getCount(line)

		if count["red"] <= 12 && count["green"] <= 13 && count["blue"] <= 14 {
			possibleGames += index
		}
	}
	return possibleGames
}

func part2(lines []string) int {
	sum := 0
	for _, line := range lines {
		count := getCount(line)
		sum += count["red"] * count["green"] * count["blue"]
	}
	return sum
}

func main() {
	file := readInput(("input"))

	lines := strings.Split(file, "\n")

	println("Part 1: The sum of the indexes of all possible games is:", part1(lines))
	println("Part 2: The sum of the power of the minimum set of cubes for each games is:", part2(lines))
}
