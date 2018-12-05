package src

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day5Part2() string {

	file, err := os.Open("./src/day5/input.txt")
	if err != nil {
		log.Fatal("cannot open input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var polimer string
	var letters []string
	alpahbet := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	for scanner.Scan() {
		//letters = strings.Split(scanner.Text(), "")

		polimer = scanner.Text()
	}

	lowest, _ := strconv.Atoi(Day5Part1())

	for outer := range alpahbet {
		trimmed := strings.Replace(polimer, alpahbet[outer], "", -1)
		trimmed = strings.Replace(trimmed, strings.ToUpper(alpahbet[outer]), "", -1)
		letters = strings.Split(trimmed, "")

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

		if lowest > len(letters) {
			lowest = len(letters)
		}
	}
	return strconv.Itoa(lowest)
}
