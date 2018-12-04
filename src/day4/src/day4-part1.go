package src

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day4part1() string {

	file, err := os.Open("./src/day4/input.txt")
	if err != nil {
		log.Fatal("cannot open input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sleepTotalsMap map[string]int64
	sleepTotalsMap = make(map[string]int64)

	var sleepPatternMap map[string][]int64
	sleepPatternMap = make(map[string][]int64)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sort.Strings(lines)

	var guard string
	var asleepTime int64
	var wakeUpTime int64

	for i := range lines {
		if strings.Contains(lines[i], "Guard") {
			guard = getGuardNo(lines[i])
		} else if strings.Contains(lines[i], "falls asleep") {
			asleepTime, _ = getMinutes(lines[i])
		} else if strings.Contains(lines[i], "wakes up") {
			wakeUpTime, _ = getMinutes(lines[i])

			if _, present := sleepPatternMap[guard]; !present {
				sleepPatternMap[guard] = make([]int64, 60)
			}

			for i := asleepTime; i < wakeUpTime; i++ {
				sleepPatternMap[guard][i] = sleepPatternMap[guard][i] + 1
			}
			sleepTotalsMap[guard] = sleepTotalsMap[guard] + (wakeUpTime - asleepTime)
		}
	}

	var bestSleeper string
	var bestSleepsTotal int64

	for key, val := range sleepTotalsMap {
		if bestSleepsTotal < val {
			bestSleeper = key
			bestSleepsTotal = val
		}
	}
	var topSleepMinute int64
	var topSleepMinuteIndex int

	for minute := range sleepPatternMap[bestSleeper] {
		if topSleepMinute < sleepPatternMap[bestSleeper][minute] {
			topSleepMinute = sleepPatternMap[bestSleeper][minute]
			topSleepMinuteIndex = minute
		}
	}

	answer, _ := strconv.Atoi(bestSleeper)

	return strconv.Itoa(answer * topSleepMinuteIndex)
}

func getMinutes(line string) (int64, error) {
	return strconv.ParseInt(strings.Split(strings.Split(line, ":")[1], "]")[0], 10, 64)
}

func getGuardNo(line string) string {
	return strings.Split(strings.Split(line, "#")[1], " ")[0]
}
