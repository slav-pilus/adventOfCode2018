package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./src/day2/input.txt")
	if err != nil {
		log.Fatal("cannot open input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	for _, element := range words {
		for _, element2 := range words {
			characterArray := strings.Split(element, "")
			characterArray2 := strings.Split(element2, "")

			diffCount := 0

			for i := range characterArray {
				if characterArray[i] != characterArray2[i] {
					diffCount++
				}
			}

			if diffCount == 1 {
				for i := range characterArray {
					if characterArray[i] == characterArray2[i] {
						fmt.Print(characterArray[i])
					}
				}
				os.Exit(0)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
