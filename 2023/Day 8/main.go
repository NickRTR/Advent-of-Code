// Exercise: https://adventofcode.com/2023/day/8

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(file string) string {
	rawFile, err := os.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func part1(nodes map[string][]string, instructions []string) int {
	current := "AAA"
	steps := 0

	for i := 0; current != "ZZZ"; i++ {
		steps += 1
		if i == len(instructions) {
			i = 0
		}
		if instructions[i] == "L" {
			current = nodes[current][0]
		} else {
			current = nodes[current][1]
		}
	}

	return steps
}

// Helper function to calculate Greatest Common Divisor (GCD)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Helper function to calculate Least Common Multiple (LCM)
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func part2(nodes map[string][]string, instructions []string) int {
	points := []string{}

	for key := range nodes {
		if string(key[2]) == "A" {
			points = append(points, key)
		}
	}

	results := []int{}

	for _, value := range points {
		steps := 0
		current := value
		for i := 0; string(current[2]) != "Z"; i++ {
			steps += 1
			if i == len(instructions) {
				i = 0
			}
			if instructions[i] == "L" {
				current = nodes[current][0]
			} else {
				current = nodes[current][1]
			}
		}
		results = append(results, steps)
	}

	lcmValue := results[0]
	for _, value := range results[1:] {
		lcmValue = lcm(lcmValue, value)
	}

	return lcmValue
}

func main() {
	file := readInput("input")

	lines := strings.Split(file, "\n")
	instructions := strings.Split(lines[0], "")

	nodes := make(map[string][]string)

	for _, line := range lines[2:] {
		characters := strings.Split(line, "")
		target := strings.Join(characters[0:3], "")
		left := strings.Join(characters[7:10], "")
		right := strings.Join(characters[12:15], "")
		nodes[target] = []string{left, right}
	}

	fmt.Println("Part 1: Required steps to reach the target ZZZ:", part1(nodes, instructions))
	fmt.Println("Part 2: Required steps to have all nodes end on a Z:", part2(nodes, instructions))
}
