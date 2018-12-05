package src

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day5Part1() string {

	file, err := os.Open("./src/day5/input.txt")
	if err != nil {
		log.Fatal("cannot open input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var letters []string

	for scanner.Scan() {
		letters = strings.Split(scanner.Text(), "")
	}

	run := true

	for run {
		for i := range letters {
			if i+1 == len(letters) {
				break
			}
			current := letters[i]
			next := letters[i+1]
			if strings.ToUpper(current) == strings.ToUpper(next) && current != next {
				letters = append(letters[:i], letters[i+2:]...)
				run = true
				break
			}
			run = false
		}
	}
	return strconv.Itoa(len(letters))
}
