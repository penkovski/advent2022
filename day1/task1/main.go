package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var maxCalories int
	var currCalories int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currCalories > maxCalories {
				maxCalories = currCalories
			}
			currCalories = 0
			continue
		}

		cals, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalln("invalid input: ", err)
		}

		currCalories += cals
	}

	fmt.Println(maxCalories)
}
