// Exercise: https://adventofcode.com/2024/day/6

package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

func checkCoordinate(coordinates []int, grid [130][]string) string {
	y := coordinates[0]
	x := coordinates[1]
	if y < 0 || y > len(grid[0])-1 || x < 0 || x > len(grid)-1 {
		return "outside"
	}
	if grid[y][x] == "." {
		return "continue"
	} else if grid[y][x] == "#" {
		return "obstacle"
	} else {
		return "error"
	}
}

func positionAhead(direction int, coordinates []int) []int {
	y := coordinates[0]
	x := coordinates[1]
	switch direction {
	case 0:
		y--
	case 1:
		x++
	case 2:
		y++
	case 3:
		x--
	}
	return []int{y, x}
}

func part1(grid [130][]string, startingPoint []int) (int, time.Duration) {
	start := time.Now()

	positions := [][]int{startingPoint}
	direction := 0
	guard := startingPoint

	for true {
		status := checkCoordinate(positionAhead(direction, guard), grid)
		if status == "outside" {
			break
		}

		if status == "obstacle" {
			direction++
			if direction == 4 {
				direction = 0
			}
		}

		guard = positionAhead(direction, guard)
		if slices.IndexFunc(positions, func(position []int) bool {
			return position[0] == guard[0] && position[1] == guard[1]
		}) == -1 {
			positions = append(positions, guard)
		}
	}

	duration := time.Since(start)

	return len(positions), duration
}

func part2(grid [130][]string) (int, time.Duration) {
	start := time.Now()

	counter := 0

	duration := time.Since(start)

	return counter, duration
}

func main() {
	start := time.Now()

	file := readInput("input")
	lines := strings.Split(file, "\n")

	var grid [130][]string
	var startingPoint []int

	for y, line := range lines {
		characters := strings.Split(line, "")
		startingX := slices.Index(characters, "^")
		if startingX != -1 {
			startingPoint = []int{y, startingX}
		}
		grid[y] = characters
	}

	part1, part1Duration := part1(grid, startingPoint)
	// part2, part2Duration := part2(grid)

	duration := time.Since(start)

	fmt.Printf("✅ Part 1: The guard will visit %v distinct positions before leaving the mapped area (Execution time: %s)\n", part1, part1Duration)
	// fmt.Printf("✅ Part 2: The result is: %v (Execution time: %s)\n", part2, part2Duration)
	fmt.Printf("🚀 Total execution time: %s\n", duration)
}
