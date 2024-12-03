// Exercise: https://adventofcode.com/2024/day/

package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func readInput(file string) string {
	rawFile, err := os.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func part1() (int, time.Duration) {
	start := time.Now()

	duration := time.Since(start)

	return 0, duration
}

func part2() (int, time.Duration) {
	start := time.Now()

	duration := time.Since(start)

	return 0, duration
}

func main() {
	start := time.Now()

	file := readInput("input")

	part1, part1Duration := part1()
	part2, part2Duration := part2()

	duration := time.Since(start)

	fmt.Printf("✅ Part 1: The result is: %v (Execution time: %s)\n", part1, part1Duration)
	fmt.Printf("✅ Part 2: The result is: %v (Execution time: %s)\n", part2, part2Duration)
	fmt.Printf("🚀 Total execution time: %s\n", duration)
}
