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
	rawFile, err := os.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func convertSeed(seeds []int, lines []string) []int {
	newSeeds := []int{}

	for _, line := range lines {
		instructions := strings.Split(line, " ")
		rangeLength, _ := strconv.Atoi(instructions[2])
		destination, _ := strconv.Atoi(instructions[0])
		source, _ := strconv.Atoi(instructions[1])
		distance := destination - source
		for i, seed := range seeds {
			if seed >= source && seed <= source+rangeLength {
				newSeeds = append(newSeeds, seed+distance)
				// remove converted seeds from seeds slice
				if i == 0 {
					seeds = seeds[1:]
				} else if i < len(seeds)-1 {
					seeds = append(seeds[:i], seeds[i+1:]...)
				} else {
					seeds = seeds[:len(seeds)-1]
				}
			}
		}
	}

	for _, seed := range seeds {
		newSeeds = append(newSeeds, seed)
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
	min := seeds[0]
	for _, seed := range seeds {
		if seed < min {
			min = seed
		}
	}

	return min
}

type interval struct {
	start int
	count int
}

func removeSeeds(seeds []interval, i int) []interval {
	// remove converted seeds from seeds slice
	if i == 0 {
		seeds = seeds[1:]
	} else if i < len(seeds)-1 {
		seeds = append(seeds[:i], seeds[i+1:]...)
	} else {
		seeds = seeds[:len(seeds)-1]
	}
	return seeds
}

func convertSeedByInterval(seeds []interval, lines []string) []interval {
	newSeeds := []interval{}

	// - seed range is entirely before mapping range
	// - seed range starts before mapping range but ends inside mapping range
	// - seed range starts inside mapping range and ends inside mapping range
	// - seed range starts inside mapping range but ends outside mapping range
	// - seed range is entirely after mapping range
	// - seed range starts before mapping range and ends after mapping range

	for _, line := range lines {
		instructions := strings.Split(line, " ")
		rangeLength, _ := strconv.Atoi(instructions[2])
		destination, _ := strconv.Atoi(instructions[0])
		source, _ := strconv.Atoi(instructions[1])
		distance := destination - source
		for i := 0; i < len(seeds); i++ {
			seed := seeds[i]
			if seed.start <= source {
				if seed.start+seed.count <= source+rangeLength && seed.start+seed.count >= source {
					newSeeds = append(newSeeds, interval{source + distance, seed.count})
					seeds = removeSeeds(seeds, i)
					seeds = append(seeds[:i], interval{seed.start, source - seed.start})
				} else if seed.start+seed.count >= source+rangeLength {
					seed.start = source
					seed.count -= source - seed.start
					maxRange := int(math.Min(float64(seed.count), float64(source+rangeLength-seed.start)))
					newSeeds = append(newSeeds, interval{source + distance, maxRange})
					// add seeds after destination
					seeds = append(seeds[:i], interval{seed.start + rangeLength, seed.count - rangeLength})
				}
				// add seeds before destination
				seeds = append(seeds[:i], interval{seed.start, source - seed.start})
			} else if seed.start >= source && seed.start <= source+rangeLength {
				maxRange := int(math.Min(float64(seed.count), float64(source+rangeLength-seed.start)))
				newSeeds = append(newSeeds, interval{seed.start + distance, maxRange})
				seeds = removeSeeds(seeds, i)
				if maxRange < seed.count {
					seeds = append(seeds[:i], interval{seed.start + maxRange, seed.count - maxRange})
				}
			}
		}
	}

	for _, seed := range seeds {
		newSeeds = append(newSeeds, seed)
	}

	return newSeeds
}

func part2(file string) int {
	sections := strings.Split(file, "\n\n")

	seeds := []interval{}

	seedsString := strings.Split(sections[0], " ")

	for i := 1; i <= len(seedsString[1:]); i += 2 {
		start, _ := strconv.Atoi(seedsString[i])
		count, _ := strconv.Atoi(seedsString[i+1])
		seeds = append(seeds, interval{start, count})
	}

	for _, section := range sections[1:] {
		lines := strings.Split(section, "\n")
		seeds = convertSeedByInterval(seeds, lines[1:])
	}

	// find minimum location
	min := seeds[0].start
	for _, seed := range seeds {
		if seed.start < min {
			min = seed.start
		}
	}

	return min
}

func main() {
	file := readInput(("input"))

	start := time.Now()
	fmt.Println("Part 1: The lowest location number that corresponds to any of the initial seed numbers is:", part1(file))
	fmt.Println("Execution Time (Part 1):", time.Now().Sub(start))
	fmt.Println()
	start = time.Now()
	fmt.Println("Part 2: The lowest location number that corresponds to any of the initial seed numbers is:", part2(file))
	fmt.Println("Execution Time (Part 2):", time.Now().Sub(start))
}
