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
	lineTotal := ""
	total := 0
	numbers := [18]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for scanner.Scan() {
		firstIndex := -1
		lastIndex := -1
		firstInt := -1
		lastInt := -1
		lineTotal = ""
		currentLine := scanner.Text()
		fmt.Println(currentLine)

		for j := 0; j < len(numbers); j++ {
			i := strings.Index(currentLine, numbers[j])
			if i == -1 {
				continue
			}

			if firstIndex == -1 || i < firstIndex {
				firstIndex = i
				firstInt = getNumber(numbers[j])
			}

			if lastIndex == -1 || i > lastIndex {
				lastIndex = i
				lastInt = getNumber(numbers[j])
			}

			i = strings.LastIndex(currentLine, numbers[j])

			if firstIndex == -1 || i < firstIndex {
				firstIndex = i
				firstInt = getNumber(numbers[j])
			}

			if lastIndex == -1 || i > lastIndex {
				lastIndex = i
				lastInt = getNumber(numbers[j])
			}
		}

		fmt.Printf("First int = %v, Last int = %v\n", firstInt, lastInt)

		lineTotal = strconv.Itoa(firstInt) + strconv.Itoa(lastInt)
		fmt.Printf("Line total is %v\n", lineTotal)
		if lineInt, err := strconv.Atoi(lineTotal); err == nil {
			total += lineInt
		}

		fmt.Printf("Total is %v\n", total)
	}

	fmt.Printf("Final total is %v\n", total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getNumber(input string) int {
	if d, err := strconv.Atoi(input); err == nil {
		return d
	}

	return convertStringToNumber(input)
}

func convertStringToNumber(num string) int {
	switch num {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0
	}
}
