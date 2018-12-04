package src

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day4part2() string {

	file, err := os.Open("./src/day4/input.txt")
	if err != nil {
		log.Fatal("cannot open input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
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
		}
	}

	var topGuard, topMinuteIndex int
	var topMinute int64
	for key, value := range sleepPatternMap {
		for minute := range value {
			if value[minute] > topMinute {
				topMinuteIndex = minute
				topMinute = value[minute]
				topGuard, _ = strconv.Atoi(key)
			}
		}
	}

	return strconv.Itoa(topGuard * topMinuteIndex)
}
