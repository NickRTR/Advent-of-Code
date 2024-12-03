// Exercise: https://adventofcode.com/2024/day/3

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

func calculateOperation(operation string) int {
	operation = operation[4 : len(operation)-1]
	values := strings.Split(operation, ",")
	a, _ := strconv.Atoi(values[0])
	b, _ := strconv.Atoi(values[1])
	return a * b
}

func part1(file string) (int, time.Duration) {
	start := time.Now()

	re := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	operations := re.FindAllString(file, -1)

	result := 0
	for _, operation := range operations {
		result += calculateOperation(operation)
	}

	duration := time.Since(start)

	return result, duration
}

func part2(file string) (int, time.Duration) {
	start := time.Now()

	re := regexp.MustCompile(`(mul\([0-9]{1,3},[0-9]{1,3}\))|(do\(\))|(don't\(\))`)
	operations := re.FindAllString(file, -1)

	result := 0

	mode := "enabled"
	for _, operation := range operations {
		switch operation {
		case "do()":
			mode = "enabled"
		case "don't()":
			mode = "disabled"
		default:
			if mode == "enabled" {
				result += calculateOperation(operation)
			}
		}
	}

	duration := time.Since(start)

	return result, duration
}

func main() {
	start := time.Now()

	file := readInput("input")

	part1, part1Duration := part1(file)
	part2, part2Duration := part2(file)

	duration := time.Since(start)

	fmt.Printf("✅ Part 1: The sum of all valid operations is: %v (Execution time: %s)\n", part1, part1Duration)
	fmt.Printf("✅ Part 2: The sum of all valid and enabled operations is: %v (Execution time: %s)\n", part2, part2Duration)
	fmt.Printf("🚀 Total execution time: %s\n", duration)
}
