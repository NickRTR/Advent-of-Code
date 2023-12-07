package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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

var cardValues = map[string]int{
	"2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "T": 9, "J": 10, "Q": 11, "K": 12, "A": 13,
}

func getCardValue(card string) int {
	return cardValues[card]
}

func getHandTypeValue(cards []string) int {
	values := map[string]int{}

	for _, card := range cards {
		values[card] += 1
	}

	switch len(values) {
	case 1:
		return 7 // Five of a kind
	case 5:
		return 1 // High card
	default:
		for _, value := range values {
			switch value {
			case 4:
				return 6 // Four of a kind
			case 3:
				if len(values) == 2 {
					return 5 // Full house
				}
				return 4 // Three of a kind
			case 2:
				if len(values) == 3 {
					return 3 // Two pair
				}
			}
		}
		return 2
	}
}

func getHandValue(hand string) int {
	cards := strings.Split(hand, "")

	cardValue := getHandTypeValue(cards)

	for _, card := range cards {
		value := getCardValue(card)
		cardValue = cardValue*100 + value
	}

	return cardValue
}

func part1(lines []string) int {
	handValues := map[int]int{}
	sum := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")

		bid, _ := strconv.Atoi(parts[1])
		hand := parts[0]

		value := getHandValue(hand)

		handValues[value] += bid
	}

	keys := make([]int, 0, len(handValues))
	for key := range handValues {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for i, key := range keys {
		sum += (i + 1) * handValues[key]
	}

	return sum
}

func main() {
	file := readInput("input")

	lines := strings.Split(file, "\n")

	fmt.Println("Part 1: The total winnings are:", part1(lines))
}
