package days

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/lucaschain/advent-of-code/helpers"
)

func cleanEmptySliceItems(slice []string) []string {
	var cleanSlice []string
	for _, item := range slice {
		if strings.Trim(item, " ") != "" {
			cleanSlice = append(cleanSlice, item)
		}
	}

	return cleanSlice
}

func numberSetsFromString(numberSetsString string) ([]string, []string) {
	numberSets := strings.Split(numberSetsString, "|")

	winningNumbersPart := cleanEmptySliceItems(strings.Split(numberSets[0], " "))
	ourNumbersPart := cleanEmptySliceItems(strings.Split(numberSets[1], " "))

	return winningNumbersPart, ourNumbersPart
}

func calculateCopies(winningNumbersPart []string, ourNumbersPart []string) int {
	matches := 0
	for _, winningNumber := range winningNumbersPart {
		for _, ourNumber := range ourNumbersPart {
			if winningNumber == ourNumber {
				matches++
			}
		}
	}

	return matches
}

func calculatePoints(copyCount int) int {
	return int(math.Pow(2, float64(copyCount-1)))
}

func updateCopies(currentCard int, copyCount int, copies map[int]int) map[int]int {
	copyStart := currentCard + 1
	copyEnd := copyStart + copyCount
	for i := copyStart; i < copyEnd; i++ {
		copies[i] += 1
	}
	return copies
}

func processLine(line string) int {
	numberSetsString := strings.Trim(
		strings.Split(line, ":")[1],
		" ",
	)
	winningNumbersPart, ourNumbersPart := numberSetsFromString(numberSetsString)
	return calculateCopies(winningNumbersPart, ourNumbersPart)
}

func Day4() string {
	lines := helpers.Read("input/day4.txt")

	var uniqueCardPointSum int
	var totalCards int
	copies := map[int]int{}
	start := time.Now()
	for currentCardNumber, line := range lines {
		copiesOfThisCard := copies[currentCardNumber]
		fmt.Printf("Processing line %d - %d copies\n", currentCardNumber, copiesOfThisCard)
		for i := 0; i <= copiesOfThisCard; i++ {
			totalCards++
			copyCount := processLine(line)
			copies = updateCopies(
				currentCardNumber,
				copyCount,
				copies,
			)

			if i == 0 {
				uniqueCardPointSum += calculatePoints(copyCount)
			}
		}
	}

	fmt.Printf("Took %s\n", time.Since(start))

	return fmt.Sprintf("Card Points Sum: %d, total cards: %d", uniqueCardPointSum, totalCards)
}
