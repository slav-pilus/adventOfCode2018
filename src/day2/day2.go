package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func buildFrequencyMap(word string) map[int]int {
	var frequencyMap map[int]int
	frequencyMap = make(map[int]int)

	stringArray := strings.Split(word, "")
	sort.Strings(stringArray)

	blockCount := 0
	lastCharacter := stringArray[0]

	for _, element := range stringArray {
		if lastCharacter == element {
			blockCount++
		} else {
			if value, exists := frequencyMap[blockCount]; exists {
				frequencyMap[blockCount] = value + 1
			} else {
				frequencyMap[blockCount] = 1
			}

			blockCount = 1
			lastCharacter = element
		}
	}

	if value, exists := frequencyMap[blockCount]; exists {
		frequencyMap[blockCount] = value + 1
	} else {
		frequencyMap[blockCount] = 1
	}

	return frequencyMap
}

func main() {
	file, err := os.Open("./src/day2/input.txt")
	if err != nil {
		log.Fatal("cannot open input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	twoCount := 0
	threeCount := 0

	for scanner.Scan() {
		frequencyMap := buildFrequencyMap(scanner.Text())

		if _, exists := frequencyMap[2]; exists {
			twoCount++
		}
		if _, exists := frequencyMap[3]; exists {
			threeCount++
		}
	}

	fmt.Println(twoCount * threeCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
