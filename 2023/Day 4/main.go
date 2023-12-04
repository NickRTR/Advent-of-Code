// Exercise: https://adventofcode.com/2023/day/4

package main

import (
	"log"
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
	sum := 0

	for _, line := range lines {
		parts := strings.Split(line, ":")
		game := strings.Split(parts[1], "|")

		winValues := strings.Split(game[0], " ")
		var winners []int

		for _, number := range winValues {
			if number != "" {
				converted, _ := strconv.Atoi(number)
				winners = append(winners, converted)
			}
		}

		gameValues := strings.Split(game[1], " ")
		var values []int

		for _, number := range gameValues {
			if number != "" {
				converted, _ := strconv.Atoi(number)
				values = append(values, converted)
			}
		}

		cardValue := 0

		for _, value := range values {
			for _, winner := range winners {
				if value == winner {
					if cardValue == 0 {
						cardValue = 1
					} else {
						cardValue *= 2
					}
				}
			}
		}
		sum += cardValue
	}
	return sum
}

func main() {
	file := readInput(("input"))

	lines := strings.Split(file, "\n")

	println("Part 1: The total points the cards are worth is:", part1(lines))
}
