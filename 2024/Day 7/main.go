// Exercise: https://adventofcode.com/2024/day/7

package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Equation struct {
	result int
	values []int
}

func readInput(file string) string {
	rawFile, err := os.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func solve(interimResults []int, values []int, usingConcatenation bool) []int {
	if len(values) == 0 {
		return interimResults
	}
	var results []int
	for _, result := range interimResults {
		results = append(results, result*values[0])
		results = append(results, result+values[0])
		if usingConcatenation {
			concatenated, _ := strconv.Atoi(fmt.Sprintf("%v%v", result, values[0]))
			results = append(results, concatenated)
		}
	}
	return solve(results, values[1:], usingConcatenation)
}

func part1(equations []Equation) (int, time.Duration) {
	start := time.Now()

	sum := 0
	for _, equation := range equations {
		results := solve(equation.values[:1], equation.values[1:], false)
		if slices.Index(results, equation.result) != -1 {
			sum += equation.result
		}
	}

	duration := time.Since(start)

	return sum, duration
}

func part2(equations []Equation) (int, time.Duration) {
	start := time.Now()

	sum := 0
	for _, equation := range equations {
		results := solve(equation.values[:1], equation.values[1:], true)
		if slices.Index(results, equation.result) != -1 {
			sum += equation.result
		}
	}

	duration := time.Since(start)

	return sum, duration
}

func main() {
	start := time.Now()

	file := readInput("input")
	lines := strings.Split(file, "\n")

	var equations []Equation

	for _, equation := range lines {
		resultSection, valueSection, _ := strings.Cut(equation, ":")

		result, _ := strconv.Atoi(resultSection)

		values := strings.Split(strings.Trim(valueSection, " "), " ")
		var v []int
		for _, i := range values {
			number, _ := strconv.Atoi(i)
			v = append(v, number)
		}
		equations = append(equations, Equation{result, v})
	}

	part1, part1Duration := part1(equations)
	part2, part2Duration := part2(equations)

	duration := time.Since(start)

	fmt.Printf("✅ Part 1: The total calibration result of all possible equations is: %v (Execution time: %s)\n", part1, part1Duration)
	fmt.Printf("✅ Part 2: The total calibration result of all possible equations, including the concatenation operator is: %v (Execution time: %s)\n", part2, part2Duration)
	fmt.Printf("🚀 Total execution time: %s\n", duration)
}
