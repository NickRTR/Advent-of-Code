// Exercise: https://adventofcode.com/2022/day/4

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readInput(file string) string {
	// read input
	rawFile, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	return string(rawFile)
}

func main() {
	input := readInput("input.txt")

	counter := 0
	counter2 := 0

	// split input into pairs of two
	pairs := strings.Split(input, "\n")
	for _, pair := range pairs {
		if len(pair) > 0 {
			assignments := strings.Split(pair, ",")

			var fields []int

			for _, assignment := range assignments {
				field := strings.Split(assignment, "-")
				for i := 0; i < 2; i++ {
					f, _ := strconv.Atoi(field[i])
					fields = append(fields, f)
				}
			}

			// check if the fields in the assignments of each pair of elves contain each other
			if fields[2] >= fields[0] && fields[3] <= fields[1] {
				counter++
			} else if fields[0] >= fields[2] && fields[1] <= fields[3] {
				counter++
			}

			// Part 2
			if fields[0] <= fields[3] && fields[1] >= fields[2] {
				counter2++
			}
		}
	}
	fmt.Printf("In %v assignment pairs one range fully contains the other.\n", counter)
	fmt.Printf("In %v assignment pairs there are ranges overlapping\n", counter2)
}
