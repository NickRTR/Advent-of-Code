package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

func readInput(file string) string {
	// read input
	rawFile, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func alphabeticalIndex(input rune) int {
	index := int(unicode.ToLower(input)) - 96
	if unicode.IsUpper(input) {
		index += 26
	}
	return index
}

func part1(input string) {
	rucksacks := strings.Split(input, "\n")

	prioritySum := 0

	for _, r := range rucksacks {
		items := strings.Split(r, "")
		compartment1 := items[:(len(items) / 2)]
		compartment2 := items[(len(items) / 2):]

		counter := make(map[string]bool)

		for _, i := range compartment1 {
			counter[i] = true
		}

		for _, i := range compartment2 {
			if counter[i] {
				prioritySum += alphabeticalIndex(rune(i[0]))
				break
			}
		}
	}

	fmt.Printf("Sum of the priorities of the duplicate item types is %v.\n", prioritySum)
}

func createMap() {

}

func part2(input string) {
	lines := strings.Split(input, "\n")
	var groups [][]string

	// divide input into groups of three Elves
	for i := 0; i < len(lines); i += 3 {
		end := i + 3

		if end > len(lines) {
			end = len(lines)
		}

		groups = append(groups, lines[i:end])
	}

	prioritySum := 0

	for _, g := range groups {
		if len(g) == 3 {
			var counter [3]map[rune]bool

			for i := 0; i < 3; i++ {
				counter[i] = make(map[rune]bool)

				for _, item := range g[i] {
					counter[i][item] = true
				}
			}

			for item := range counter[0] {
				if counter[1][item] && counter[2][item] {
					prioritySum += alphabeticalIndex(item)
					break
				}
			}
		}
	}

	fmt.Printf("Sum of the priorities of the group badges is %v.\n", prioritySum)
}

func main() {
	input := readInput("input")
	part1(input)
	part2(input)
}
