package days

import (
	"fmt"
	"strings"

	"github.com/lucaschain/advent-of-code/helpers"
)

func containsDot(pattern string) bool {
	for _, c := range pattern {
		if c == '.' {
			return true
		}
	}
	return false
}

var cache = map[string]int{}

func countMatches(pattern string, groupSizes []int, patternIndex int, groupIndex, possibilitiesLeft int) int {
	cacheKey := fmt.Sprintf("%s-%d-%d-%d-%v", pattern, patternIndex, groupIndex, possibilitiesLeft, groupSizes)
	if _, ok := cache[cacheKey]; ok {
		return cache[cacheKey]
	}
	if groupIndex >= len(groupSizes) {
		if strings.Contains(pattern[patternIndex-1:], "#") {
			return 0
		}
		return 1
	}

	solutions := 0
	unknownCells := len(pattern) - patternIndex - possibilitiesLeft

	groupSizesLeft := len(groupSizes) - groupIndex
	runSize := groupSizesLeft + unknownCells
	currentGroupSize := groupSizes[groupIndex]
	for runIndex := 0; runIndex < runSize; runIndex++ {
		validChars := strings.Repeat(".", runIndex)
		validChars += strings.Repeat("#", currentGroupSize) + "."

		patternPart := pattern[patternIndex:]
		valid := true
		for i, validChar := range validChars {
			if i >= len(patternPart) {
				break
			}
			originalChar := rune(patternPart[i])

			if originalChar == '?' {
				continue
			}

			if originalChar != validChar {
				valid = false
				break
			}
		}

		if !valid {
			continue
		}

		solutions += countMatches(
			pattern,
			groupSizes,
			patternIndex+runIndex+currentGroupSize+1,
			groupIndex+1,
			possibilitiesLeft-currentGroupSize-1,
		)

	}

	cache[cacheKey] = solutions
	return solutions
}

func extractLine(line string) (string, []int) {
	parts := strings.Split(line, " ")
	pattern := parts[0]

	groupSizes := []int{}
	for _, c := range strings.Split(parts[1], ",") {
		groupSizes = append(groupSizes, helpers.ToInt(string(c)))
	}

	return pattern, groupSizes
}

func unfoldPattern(pattern string) string {
	unfolded := pattern
	for i := 1; i < 5; i++ {
		unfolded += "?" + pattern
	}
	return unfolded
}

func unfoldGroupSize(groupSizes []int) []int {
	unfolded := []int{}
	for i := 0; i < 5; i++ {
		unfolded = append(unfolded, groupSizes...)
	}
	return unfolded
}

func countLineArrangements(line string) int {
	linePattern, groupSizes := extractLine(line)

	linePattern = unfoldPattern(linePattern)
	groupSizes = unfoldGroupSize(groupSizes)

	return countMatches(linePattern, groupSizes, 0, 0, helpers.SliceSum(groupSizes)+len(groupSizes)-1)
}

func Day12() string {
	lines := helpers.Read("input/day12.txt")

	arrangementsSum := 0
	for _, line := range lines {
		arrangements := countLineArrangements(line)
		println(arrangements, line)
		arrangementsSum += arrangements
	}

	return fmt.Sprintf("Sum of the arrangements: %d", arrangementsSum)
}
