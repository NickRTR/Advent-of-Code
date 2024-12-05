// Exercise: https://adventofcode.com/2024/day/5

package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func readInput(file string) string {
	rawFile, err := os.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func checkUpdate(rules map[int][]int, update []int) bool {
	for i, page := range update {
		rule, exists := rules[page]
		if exists {
			for _, value := range rule {
				if slices.Contains(update[i:], value) {
					return false
				}
			}
		}
	}
	return true
}

func fixUpdate(rules map[int][]int, update []int) int {
	sort.Slice(update, func(i, j int) bool {
		return slices.Contains(rules[update[i]], update[j])
	})
	return update[len(update)/2]
}

func part1(rules map[int][]int, updates [][]int) (int, time.Duration) {
	start := time.Now()

	result := 0
	for _, update := range updates {
		if checkUpdate(rules, update) {
			result += update[len(update)/2]
		}
	}

	duration := time.Since(start)

	return result, duration
}

func part2(rules map[int][]int, updates [][]int) (int, time.Duration) {
	start := time.Now()

	result := 0
	for _, update := range updates {
		if !checkUpdate(rules, update) {
			result += fixUpdate(rules, update)
		}
	}

	duration := time.Since(start)

	return result, duration
}

func main() {
	start := time.Now()

	file := readInput("input")

	sections := strings.Split(file, "\n\n")
	rulesSection := strings.Split(sections[0], "\n")
	updatesSection := strings.Split(sections[1], "\n")

	rules := make(map[int][]int)
	for _, rule := range rulesSection {
		r := strings.Split(rule, "|")
		left, _ := strconv.Atoi(r[0])
		right, _ := strconv.Atoi(r[1])
		rules[right] = append(rules[right], left)
	}

	var updates [][]int
	for _, update := range updatesSection {
		pages := strings.Split(update, ",")
		u := make([]int, len(pages))
		for i, page := range pages {
			p, _ := strconv.Atoi(page)
			u[i] = p
		}
		updates = append(updates, u)
	}

	part1, part1Duration := part1(rules, updates)
	part2, part2Duration := part2(rules, updates)

	duration := time.Since(start)

	fmt.Printf("✅ Part 1: The sum of the middle page numbers from the correctly-ordered updates is: %v (Execution time: %s)\n", part1, part1Duration)
	fmt.Printf("✅ Part 2: After ordering the wrongly-ordered updates, the sum of the middle page numbers is: %v (Execution time: %s)\n", part2, part2Duration)
	fmt.Printf("🚀 Total execution time: %s\n", duration)
}
