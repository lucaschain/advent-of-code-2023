package days

import (
	"fmt"
	"strings"

	"github.com/lucaschain/advent-of-code/helpers"
)

func extractRecords(line string) []int {
	recordsPart := strings.Trim(strings.Split(line, ":")[1], " ")
	splitLine := strings.Split(recordsPart, " ")
	records := []int{}

	for _, record := range splitLine {
		if record == "" {
			continue
		}
		records = append(records, helpers.ToInt(record))
	}

	return records
}

func extractRecordsConsideringKerning(line string) int {
	recordsPart := strings.Trim(strings.Split(line, ":")[1], " ")
	strNumber := strings.ReplaceAll(recordsPart, " ", "")

	return helpers.ToInt(strNumber)
}

func countWaysToBeat(totalRaceTimeMs, record int) int {
	numberOfWays := 0

	for i := 0; i < totalRaceTimeMs; i++ {
		distance := i * (totalRaceTimeMs - i)

		if distance > record {
			numberOfWays++
		}
	}

	return numberOfWays
}

func Day6() string {
	lines := helpers.Read("input/day6.txt")
	totalRaceTimeMs := extractRecordsConsideringKerning(lines[0])
	raceRecord := extractRecordsConsideringKerning(lines[1])

	numberOfWays := countWaysToBeat(totalRaceTimeMs, raceRecord)

	return fmt.Sprintf("Possible ways to win game %d", numberOfWays)
}
