package days

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/lucaschain/advent-of-code/helpers"
)

type CubeSubset struct {
	Blue  int
	Green int
	Red   int
}

func getGameId(line string) int {
	gameLabel := strings.Split(line, ":")[0]

	id, err := strconv.Atoi(strings.Replace(gameLabel, "Game ", "", 1))

	if err != nil {
		log.Fatal(err)
	}

	return id
}

func getColorFromSubset(subsetPart string, colorName string) int {
	colorCount := 0
	subsets := strings.Split(subsetPart, ";")

	for _, subset := range subsets {
		subset = strings.Trim(subset, " ")
		subsetColors := strings.Split(subset, ",")
		for _, subsetColor := range subsetColors {
			if strings.Contains(subsetColor, colorName) {
				colorCountRegex := regexp.MustCompile(`[a-zA-Z ]+`)

				digits := colorCountRegex.ReplaceAll([]byte(subsetColor), []byte{})

				digit, _ := strconv.Atoi(string(digits))

				return digit
			}
		}
	}

	return colorCount
}

func getGameSubsets(line string) []CubeSubset {
	gameSubsetPart := strings.Trim(strings.Split(line, ":")[1], " ")

	subsets := strings.Split(gameSubsetPart, ";")

	var cubeSubsets []CubeSubset
	for _, subset := range subsets {
		subset = strings.Trim(subset, " ")

		blue := getColorFromSubset(subset, "blue")
		green := getColorFromSubset(subset, "green")
		red := getColorFromSubset(subset, "red")

		cubeSubsets = append(cubeSubsets, CubeSubset{
			Blue:  blue,
			Green: green,
			Red:   red,
		})
	}

	return cubeSubsets
}

var limits = CubeSubset{
	Blue:  14,
	Green: 13,
	Red:   12,
}

func (subset CubeSubset) Power() int {
	return subset.Blue * subset.Green * subset.Red
}

func (subset CubeSubset) String() string {
	return fmt.Sprintf("R: %d, G: %d, B: %d, P: %d", subset.Red, subset.Green, subset.Blue, subset.Power())
}

func (subset CubeSubset) IsPossible() bool {
	return subset.Blue <= limits.Blue && subset.Green <= limits.Green && subset.Red <= limits.Red
}

func minimumPossibleSubset(subsets []CubeSubset) CubeSubset {
	minimalSubset := CubeSubset{
		Blue:  0,
		Green: 0,
		Red:   0,
	}

	for _, subset := range subsets {
		if subset.Blue > minimalSubset.Blue {
			minimalSubset.Blue = subset.Blue
		}

		if subset.Green > minimalSubset.Green {
			minimalSubset.Green = subset.Green
		}

		if subset.Red > minimalSubset.Red {
			minimalSubset.Red = subset.Red
		}
	}

	return minimalSubset
}

func Day2() string {
	lines := helpers.Read("input/day2.txt")

	idSum := 0
	minimumSubsetPowerSum := 0
	for _, line := range lines {
		gameId := getGameId(line)
		gameSubsets := getGameSubsets(line)
		gamePossible := true
		minimumSubset := minimumPossibleSubset(gameSubsets)
		minimumSubsetPowerSum += minimumSubset.Power()

		for _, subset := range gameSubsets {
			if !subset.IsPossible() {
				gamePossible = false
				break
			}
		}

		if gamePossible {
			idSum += gameId
		}
	}

	return fmt.Sprintf("Possible ID sum: %d. Sum of minimal powers: %d", idSum, minimumSubsetPowerSum)
}
