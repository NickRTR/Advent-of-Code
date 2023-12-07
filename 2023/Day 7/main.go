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

func sortByValues(m map[string]int) string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys[0]
}

var cardValues = map[string]int{
	"2": 1, "3": 2, "4": 3, "5": 4, "6": 5, "7": 6, "8": 7, "9": 8, "T": 9, "J": 10, "Q": 11, "K": 12, "A": 13,
}

var cardValuesWithJokers = map[string]int{
	"J": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "T": 9, "Q": 10, "K": 11, "A": 12,
}

func getCardValue(card string, joker bool) int {
	if joker {
		return cardValuesWithJokers[card]
	}
	return cardValues[card]
}

func getHandTypeValue(cards []string, joker bool) int {
	values := map[string]int{}

	jokerCount := 0

	for _, card := range cards {
		if joker && card == "J" {
			jokerCount += 1
			continue
		}
		values[card] += 1
	}

	if jokerCount == 5 {
		return 7
	}

	if joker {
		values[sortByValues(values)] += jokerCount
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

func getHandValue(hand string, joker bool) int {
	cards := strings.Split(hand, "")

	cardValue := getHandTypeValue(cards, joker)

	for _, card := range cards {
		value := getCardValue(card, joker)
		cardValue = cardValue*100 + value
	}

	return cardValue
}

func solution(lines []string, joker bool) int {
	handValues := map[int]int{}
	sum := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")

		bid, _ := strconv.Atoi(parts[1])
		hand := parts[0]

		value := getHandValue(hand, joker)

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

	fmt.Println("Part 1: The total winnings are:", solution(lines, false))
	fmt.Println("Part 2: The total winnings with jokers are:", solution(lines, true))
}
