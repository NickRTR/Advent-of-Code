// Exercise: https://adventofcode.com/2023/day/5

package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func readInput(file string) string {
	rawFile, err := os.ReadFile("input copy")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func convertSeed(seeds []int, lines []string) []int {
	seedMap := map[int]int{}

	for _, line := range lines {
		instructions := strings.Split(line, " ")
		rangeLength, _ := strconv.Atoi(instructions[2])
		destination, _ := strconv.Atoi(instructions[0])
		source, _ := strconv.Atoi(instructions[1])
		for i := 0; i < rangeLength; i++ {
			for _, seed := range seeds {
				if seed == source+i {
					seedMap[source+i] = destination + i
				}
			}
		}
	}

	for _, seed := range seeds {
		if _, ok := seedMap[seed]; !ok {
			seedMap[seed] = seed
		}
	}

	newSeeds := []int{}

	for _, seed := range seeds {
		newSeeds = append(newSeeds, seedMap[seed])
	}

	return newSeeds
}

func part1(file string) int {
	sections := strings.Split(file, "\n\n")

	seeds := make([]int, 0, len(sections[0]))

	seedsString := strings.Split(sections[0], " ")

	for _, str := range seedsString[1:] {
		seed, _ := strconv.Atoi(str)
		seeds = append(seeds, seed)
	}

	for _, section := range sections[1:] {
		lines := strings.Split(section, "\n")
		seeds = convertSeed(seeds, lines[1:])
	}

	// find minimum location
	min := math.MaxInt64
	for _, seed := range seeds {
		if seed < min {
			min = seed
		}
	}

	return min
}

func main() {
	file := readInput(("input"))

	start := time.Now()
	fmt.Println("Part 1: The lowest location number that corresponds to any of the initial seed numbers is:", part1(file))
	fmt.Println("Execution Time (Part 1):", time.Now().Sub(start))
}
