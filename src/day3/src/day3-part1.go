package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinates struct {
	Left int64
	Top  int64
}

type Size struct {
	Width  int64
	Length int64
}

func Day3() {
	file, err := os.Open("./src/day3/input.txt")
	if err != nil {
		log.Fatal("cannot open input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	fabric := make([][]int, 1000)
	for i := range fabric {
		fabric[i] = make([]int, 1000)
	}

	for scanner.Scan() {
		line := scanner.Text()
		coordinates := GetCoordinates(line)
		size := GetSize(line)

		for row := coordinates.Top; row < coordinates.Top+size.Length; row++ {
			for column := coordinates.Left; column < coordinates.Left+size.Width; column++ {
				fabric[row][column] = fabric[row][column] + 1
			}
		}
	}

	counter := 0

	for i := range fabric {
		for j := range fabric[i] {
			if fabric[i][j] > 1 {
				counter++
			}
		}
	}

	fmt.Println(counter)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
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
