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

func part1(lines []string) int {
	possibleGames := 0

	for _, line := range lines {
		values := strings.Split(line, ":")
		game, _ := strconv.Atoi(strings.Split(values[0], " ")[1])
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
		if count["red"] <= 12 && count["green"] <= 13 && count["blue"] <= 14 {
			possibleGames += game
		}
	}
	return possibleGames
}

func main() {
	file := readInput(("input"))

	lines := strings.Split(file, "\n")

	println("The sum of the indexes of all possible games is:", part1(lines))
}
