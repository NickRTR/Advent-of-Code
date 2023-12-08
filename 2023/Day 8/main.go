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
}
