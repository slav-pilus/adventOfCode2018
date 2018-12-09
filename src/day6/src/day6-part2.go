package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day6Part2() {
	file, err := os.Open("./src/day6/input.txt")
	if err != nil {
		log.Fatal("cannot open input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var coordinates []coordinate

	counter := 0
	for scanner.Scan() {
		coordinateArray := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(strings.TrimSpace(coordinateArray[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(coordinateArray[1]))
		coordinates = append(coordinates, coordinate{counter, x, y})
		counter++
	}
	grid, offsetX, offsetY := getGrid(coordinates)

	underCounter := 0
	for row := range grid {
		for col := range grid[row] {
			total := getDistanceTotal(col, row, coordinates, offsetX, offsetY)
			if total < 10000 {
				underCounter++
			}
		}
	}

	fmt.Println(underCounter)
}

func getDistanceTotal(x int, y int, coordinates []coordinate, offsetX int, offsetY int) (distanceTotal int) {
	for i := range coordinates {
		distance := getTaxicabDistance(y, x, coordinates[i].y-offsetY, coordinates[i].x-offsetX)
		distanceTotal += distance
	}

	return distanceTotal
}
