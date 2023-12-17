// Exercise: https://adventofcode.com/2023/day/9

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) string {
	rawFile, err := os.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func containsOnlyZeros(line []int) bool {
	for _, value := range line {
		if value != 0 {
			return false
		}
	}
	return true
}

func calculateDifferences(line []int) [][]int {
	var newLines [][]int

	newLines = append(newLines, line)

	for i := 0; !containsOnlyZeros(newLines[len(newLines)-1]); i++ {
		newLine := make([]int, len(newLines[i])-1)
		for j := 0; j < len(newLines[i])-1; j++ {
			newLine[j] = newLines[i][j+1] - newLines[i][j]
		}
		newLines = append(newLines, newLine)
	}

	return newLines
}

func solution(values [][]int, part int) int {
	sum := 0

	for _, line := range values {
		newLines := calculateDifferences(line)

		var result int
		for i := len(newLines) - 1; i >= 0; i-- {
			l := newLines[i]
			if part == 1 {
				result = l[len(l)-1] + result
			} else if part == 2 {
				result = l[0] - result
			}
		}

		sum += result
	}

	return sum
}

func main() {
	file := readInput("input")

	lines := strings.Split(file, "\n")

	values := make([][]int, len(lines))
	for i, line := range lines {
		v := strings.Split(line, " ")
		for _, value := range v {
			converted, _ := strconv.Atoi(value)
			values[i] = append(values[i], converted)
		}
	}

	fmt.Println("Part 1: The sum of the extrapolated values is:", solution(values, 1))
	fmt.Println("Part 2: The sum of the extrapolated values is:", solution(values, 2))
}
