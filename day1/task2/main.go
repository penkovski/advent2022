package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var elves []int

	i := 0
	currCalories := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			elves = append(elves, currCalories)
			currCalories = 0
			i++
			continue
		}

		cals, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln("invalid input: ", err)
		}

		currCalories += cals
	}

	// sort elves calories
	sort.Ints(elves)

	// sum the lastThree three elves calories
	lastThree := 0
	elvesCount := len(elves)
	for i := 0; i < 3; i++ {
		lastThree += elves[elvesCount-1-i]
	}

	fmt.Println(lastThree)
}
