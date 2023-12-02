package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// PART 1
	// limits := map[string]int{
	// 	"red":   12,
	// 	"green": 13,
	// 	"blue":  14,
	// }
	// counter := 0
	// totalPossibleGames := 0
	// isGamePossible := true

	// PART 2
	gameMinimums := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	runningPowerTotal := 0

	for scanner.Scan() {
		currentLine := scanner.Text()
		fmt.Println(currentLine)

		// PART 1
		// counter++
		// isGamePossible = true

		// PART 2
		gameMinimums = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		gameStats := strings.Split(currentLine, ":")[1]
		// fmt.Printf("Game stats: %s\n", gameStats)
		sets := strings.Split(gameStats, ";")

		for i := 0; i < len(sets); i++ {
			fmt.Printf("Set: %s\n", sets[i])
			cubes := strings.Split(sets[i], ",")

			for j := 0; j < len(cubes); j++ {
				trimmedCube := strings.Trim(cubes[j], " ")
				// fmt.Printf("Cube: %s\n", trimmedCube)
				cubeValues := strings.Split(trimmedCube, " ")
				count := atoi(cubeValues[0])
				colour := cubeValues[1]

				// PART 1
				// if count > limits[colour] {
				// 	fmt.Println("Set makes game impossible")
				// 	isGamePossible = false
				// }

				// PART 2
				if count > gameMinimums[colour] {
					gameMinimums[colour] = count
				}
			}
		}

		// PART 1
		// if isGamePossible {
		// 	fmt.Println("Game is possible")
		// 	totalPossibleGames += counter
		// } else {
		// 	fmt.Println("Game is not possible")
		// }

		// fmt.Printf("Running total is %v\n", totalPossibleGames)

		// PART 2
		fmt.Printf("Minimum red %v\n", gameMinimums["red"])
		fmt.Printf("Minimum blue %v\n", gameMinimums["blue"])
		fmt.Printf("Minimum green %v\n", gameMinimums["green"])
		runningPowerTotal += (gameMinimums["red"] * gameMinimums["blue"] * gameMinimums["green"])
		fmt.Printf("Running total is %v\n", runningPowerTotal)
	}

	// PART 1
	// fmt.Printf("Final total is %v\n", totalPossibleGames)

	// PART 2
	fmt.Printf("Final total is %v\n", runningPowerTotal)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}
