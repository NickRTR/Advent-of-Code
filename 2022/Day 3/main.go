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

func main() {
	input := readInput("input")

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
