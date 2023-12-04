// Exercise: https://adventofcode.com/2023/day/4

package main

import (
	"log"
	"os"
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
		winCount := countWins(line)
		score := 0
		for i := 0; i < winCount; i++ {
			if i == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
		sum += score
	}
	return sum
}

func countWins(line string) int {
	count := 0

	parts := strings.Split(line, ":")
	game := strings.Split(parts[1], "|")

	winners := strings.Split(strings.TrimSpace(game[0]), " ")
	ours := strings.Split(strings.TrimSpace(game[1]), " ")

	for _, winner := range winners {
		winner = strings.TrimSpace(winner)
		if winner != "" {
			for _, our := range ours {
				our = strings.TrimSpace(our)
				if our != "" {
					if winner == our {
						count++
					}
				}
			}
		}
	}

	return count
}

func part2(lines []string) int {
	var cards = make([]int, len(lines))

	for i := range cards {
		cards[i] = 1
	}

	for i, line := range lines {
		winCount := countWins(line)
		countTo := i + winCount
		// prevent counting beyond the end of the array
		if countTo >= len(lines)-1 {
			countTo = len(lines) - 1
		}
		for j := i + 1; j <= countTo; j++ {
			cards[j] += cards[i]
		}
	}

	// add up the total amount of cards
	sum := 0
	for _, card := range cards {
		sum += card
	}
	return sum
}

func main() {
	file := readInput(("input"))

	lines := strings.Split(file, "\n")

	println("Part 1: The total points the cards are worth is:", part1(lines))
	println("Part 2: The total amount of scratchcards is:", part2(lines))
}
