package days

import (
	"fmt"
	"strings"

	"github.com/lucaschain/advent-of-code/helpers"
)

func diffSequence(sequence []int) []int {
	diffSequence := []int{}
	diffSize := len(sequence) - 1
	for i := 0; i < diffSize; i++ {
		diffSequence = append(diffSequence, sequence[i+1]-sequence[i])
	}
	return diffSequence
}

func extractSequence(line string) []int {
	sequence := []int{}
	digits := strings.Split(line, " ")
	for _, digit := range digits {
		if digit == "" {
			continue
		}
		sequence = append(sequence, helpers.ToInt(digit))
	}
	return sequence
}

func hasOnlyZeroes(sequence []int) bool {
	for _, digit := range sequence {
		if digit != 0 {
			return false
		}
	}
	return true
}

func extrapolateBackward(sequence []int) int {
	diffSequence := diffSequence(sequence)
	firstNumber := sequence[0]
	if hasOnlyZeroes(diffSequence) {
		return firstNumber
	}

	return firstNumber - extrapolateBackward(diffSequence)
}

func extrapolateForward(sequence []int) int {
	diffSequence := diffSequence(sequence)
	lastNumber := sequence[len(sequence)-1]
	if hasOnlyZeroes(diffSequence) {
		return lastNumber
	}

	return extrapolateForward(diffSequence) + lastNumber
}

func Day9() string {
	lines := helpers.Read("input/day9.txt")
	extrapolateForwardSum := 0
	extrapolateBackwardSum := 0
	for _, line := range lines {
		sequence := extractSequence(line)
		extrapolateForwardSum += extrapolateForward(sequence)
		extrapolateBackwardSum += extrapolateBackward(sequence)
	}
	return fmt.Sprintf(
		"extrapolate forward sum %d, extrapolate backward sum %d",
		extrapolateForwardSum,
		extrapolateBackwardSum,
	)
}
