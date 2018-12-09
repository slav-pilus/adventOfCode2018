package src

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	id int
	x  int
	y  int
}

func Day6Part1() {
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

	idNearestCountMap := make(map[int]int)

	for row := range grid {
		for col := range grid[row] {
			nearestId := findNearest(col, row, coordinates, offsetX, offsetY)
			grid[row][col] = nearestId
			if val, exists := idNearestCountMap[nearestId]; exists {
				idNearestCountMap[nearestId] = val + 1
			} else {
				idNearestCountMap[nearestId] = 1
			}

		}
	}

	for row := range grid {

		for col := range grid[row] {
			if row == 0 {
				idNearestCountMap[grid[row][col]] = 0
			}
			if row == len(grid)-1 {
				idNearestCountMap[grid[row][col]] = 0

			}
			if col == len(grid[row])-1 {
				idNearestCountMap[grid[row][col]] = 0
			}
			if col == 0 {
				idNearestCountMap[grid[row][col]] = 0
			}
		}
	}

	highestCount := -1
	for _, val := range idNearestCountMap {
		if val > highestCount {
			highestCount = val
		}
	}

	fmt.Println(highestCount)
}

func findNearest(x int, y int, coordinates []coordinate, offsetX int, offsetY int) (nearestId int) {
	distances := make(map[int]int)
	nearestDistance := 99999999

	for i := range coordinates {
		distance := getTaxicabDistance(y, x, coordinates[i].y-offsetY, coordinates[i].x-offsetX)
		distances[coordinates[i].id] = distance

		if distance < nearestDistance {
			nearestDistance = distance
		}
	}

	counter := 0
	for key, val := range distances {
		if val == nearestDistance {
			counter++
			nearestId = key
		}
		if counter > 1 {
			return -1
		}
	}

	return nearestId
}

func getTaxicabDistance(firstX int, firstY int, secondX int, secondY int) int {
	return Abs(firstX-secondX) + Abs(secondY-firstY)
}

func getGrid(coordinates []coordinate) (grid [][]int, offsetX int, offsetY int) {
	top := 99999999
	bottom := 0
	left := 99999999
	right := 0

	for c := range coordinates {
		if coordinates[c].x < left {
			left = coordinates[c].x
		}
		if coordinates[c].x > right {
			right = coordinates[c].x
		}
		if coordinates[c].y < top {
			top = coordinates[c].y
		}
		if coordinates[c].y > bottom {
			bottom = coordinates[c].y
		}
	}

	grid = make([][]int, (bottom-top)+top)
	for i := range grid {
		grid[i] = make([]int, (right-left)+left)
	}
	return grid, left, top
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
