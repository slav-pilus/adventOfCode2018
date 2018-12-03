package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func main() {
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

func GetSize(line string) size {
	s := strings.Split(strings.Split(strings.Split(line, "@")[1], ": ")[1], "x")
	width, _ := strconv.ParseInt(s[0], 10, 64)
	length, _ := strconv.ParseInt(s[1], 10, 64)

	return size{width, length}
}

func GetCoordinates(line string) coordinates {
	c := strings.Split(strings.Split(strings.Split(line, "@ ")[1], ":")[0], ",")
	left, _ := strconv.ParseInt(c[0], 10, 64)
	top, _ := strconv.ParseInt(c[1], 10, 64)

	return coordinates{Left: left, Top: top}
}
