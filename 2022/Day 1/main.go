// Exercise: https://adventofcode.com/2022/day/1

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// read input
	rawFile, err := ioutil.ReadFile("input")
	if err != nil {
		log.Panicln(err)
	}
	file := string(rawFile)

	// process input
	elfs := strings.Split(file, "\n\n")
	var elfSums []int
	for _, elf := range elfs {
		calories := strings.Split(elf, "\n")
		sum := 0
		for _, calorie := range calories {
			// Parse string to int
			if len(calorie) > 0 {
				i, err := strconv.Atoi(calorie)
				if err != nil {
					log.Fatalln("Error while converting input value string to int!" + calorie)
				}
				sum += i
			}
		}
		elfSums = append(elfSums, sum)
	}
	// Sort elf and Print result
	sort.Ints(elfSums)
	fmt.Printf("Elf with most food carries %v calories.", elfSums[len(elfSums)-1])
}
