package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./src/day1/input.txt")
	if err != nil {
		log.Fatal("cannot open input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var total int64
	for scanner.Scan() {
		var currentNo, convErr = strconv.ParseInt(scanner.Text(), 10, 64)
		if convErr != nil {
			fmt.Println("cannot convert to number")
		}
		total += currentNo
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
