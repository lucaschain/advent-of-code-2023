package days

import (
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strings"

	"github.com/lucaschain/advent-of-code/helpers"
)

func countUnknownSprings(line string) int {
	count := 0
	for _, c := range line {
		if c == '?' {
			count++
		}
	}
	return count
}

func groupPatternToRegex(groupSizes []int) *regexp.Regexp {
	pattern := ""
	for i, size := range groupSizes {
		if i == 0 {
			pattern += fmt.Sprintf(`^\.*#{%d}`, size)
			continue
		}

		pattern += fmt.Sprintf(`\.+#{%d}`, size)

		if i == len(groupSizes)-1 {
			pattern += `\.*$`
		}
	}
	return regexp.MustCompile(pattern)
}

func fillInPattern(pattern string, base2Pattern string) string {
	patternCount := 0
	expanded := ""
	for _, c := range pattern {
		if c == '?' {
			expanded += string(base2Pattern[patternCount])
			patternCount++
		} else {
			expanded += string(c)
		}
	}

	return expanded
}

func expandPattern(pattern string) []string {
	unknown := countUnknownSprings(pattern)

	if unknown == 0 {
		return []string{pattern}
	}

	possibilities := math.Pow(2, float64(unknown))
	expanded := []string{}
	for i := 0; i < int(possibilities); i++ {
		base2 := big.NewInt(int64(i)).Text(2)
		base2 = strings.ReplaceAll(base2, "0", ".")
		base2 = strings.ReplaceAll(base2, "1", "#")

		for len(base2) < unknown {
			base2 = "." + base2
		}

		expanded = append(expanded, fillInPattern(pattern, base2))
	}

	return expanded
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

func countLineArrangements(line string) int {
	linePattern, groupSizes := extractLine(line)
	expanded := expandPattern(linePattern)
	regexp := groupPatternToRegex(groupSizes)

	count := 0
	for _, e := range expanded {
		if regexp.MatchString(e) {
			count++
		}
	}

	return count
}

func Day12() string {
	lines := helpers.Read("input/day12.txt")

	arrangementsSum := 0
	for _, line := range lines {
		arrangementsSum += countLineArrangements(line)
	}

	return fmt.Sprintf("Sum of the arrangements: %d", arrangementsSum)
}
