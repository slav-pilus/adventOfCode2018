package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type coordinates struct {
	Left int64
	Top  int64
}

type size struct {
	Width  int64
	Length int64
}

func Day3part2() {
	file, err := os.Open("./src/day3/input.txt")
	if err != nil {
		log.Fatal("cannot open input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	idMap := make(map[string]int64)

	fabric := make([][]string, 1000)
	for i := range fabric {
		fabric[i] = make([]string, 1000)
	}

	for scanner.Scan() {
		line := scanner.Text()
		id := getId(line)
		coordinates := GetCoordinates(line)
		size := GetSize(line)
		idMap[id] = size.Width * size.Length

		for row := coordinates.Top; row < coordinates.Top+size.Length; row++ {
			for column := coordinates.Left; column < coordinates.Left+size.Width; column++ {
				if len(fabric[row][column]) != 0 {
					fabric[row][column] = "x"
				} else {
					fabric[row][column] = id
				}
			}
		}
	}

	for key, value := range idMap {
		counter := 0
		for i := range fabric {
			for j := range fabric[i] {
				if fabric[i][j] == key {
					counter++
				}
			}
		}

		if int64(counter) == value {
			fmt.Println(key)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getId(line string) string {
	s := strings.Trim(strings.Split(line, "@")[0], " ")

	return s
}
