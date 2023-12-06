package days

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/lucaschain/advent-of-code/helpers"
)

var digitNames = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func firstDigit(input string) string {
	for i, char := range input {
		if unicode.IsDigit(char) {
			return string(char)
		}

		for digit, digitName := range digitNames {
			if strings.HasPrefix(input[i:], digitName) {
				return strconv.Itoa(digit)
			}
		}
	}

	return "0"
}

func lastDigit(input string) string {
	for i := len(input) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(input[i])) {
			return string(input[i])
		}

		for digit, digitName := range digitNames {
			if strings.HasPrefix(input[i:], digitName) {
				return strconv.Itoa(digit)
			}
		}
	}

	return "0"
}

func Day1() string {
	lines := helpers.Read("input/day1.txt")

	sum := 0
	for _, line := range lines {
		first := firstDigit(line)
		last := lastDigit(line)

		pair, _ := strconv.Atoi(first + last)

		sum += pair
	}

	return strconv.Itoa(sum)
}
