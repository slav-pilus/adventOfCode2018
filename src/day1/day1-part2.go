package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var total int64
	var frequencyMap map[int64]int64
	frequencyMap = make(map[int64]int64)

	for {
		file, err := os.Open("./src/day1/input.txt")
		scanner := bufio.NewScanner(file)

		if err != nil {
			log.Fatal("cannot open input file")
		}

		defer file.Close()

		for scanner.Scan() {
			var currentNo, convErr = strconv.ParseInt(scanner.Text(), 10, 64)
			if convErr != nil {
				fmt.Println("cannot convert to number")
			}

			total += currentNo
			if _, exists := frequencyMap[total]; exists {
				fmt.Println("duplication found : " + string(total))
				fmt.Println(total)
				os.Exit(0)
			} else {
				frequencyMap[total] = 1
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
