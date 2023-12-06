package days

import (
	"fmt"
	"log"
	"strconv"

	"github.com/lucaschain/advent-of-code/helpers"
)

type NumberPartStatus struct {
	Number string
	IsPart bool
}

func (n NumberPartStatus) String() string {
	return fmt.Sprintf("Number: %s, IsPart: %t", n.Number, n.IsPart)
}

func isDigit(char rune) bool {
	_, err := strconv.Atoi(string(char))
	return err == nil
}

func get(x int, y int, lines []string) rune {
	if y >= 0 && x >= 0 && y < len(lines) && x < len(lines[y]) {
		return rune(lines[y][x])
	}

	return 0
}

func isSymbol(char rune) bool {
	if char == 0 {
		return false
	}
	if isDigit(char) {
		return false
	}

	if char == '.' {
		return false
	}

	return true
}

func isPart(x int, y int, lines []string) bool {
	dirs := [][]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	for _, dir := range dirs {
		char := get(x+dir[0], y+dir[1], lines)
		if isSymbol(char) {
			return true
		}
	}

	return false
}

func Day3() string {
	lines := helpers.Read("input/day3.txt")

	var numberStatuses []NumberPartStatus
	var currentNumberStatus NumberPartStatus
	var engines map[int]map[int]int

	for y, line := range lines {
		for x, char := range line {
			if isDigit(char) {
				currentNumberStatus.Number += string(char)
				currentNumberStatus.IsPart = currentNumberStatus.IsPart || isPart(x, y, lines)

				if currentNumberStatus.IsPart && char == '*' {
					engines[x][y] += 1
				}
			} else {
				if currentNumberStatus.Number != "" {
					numberStatuses = append(numberStatuses, currentNumberStatus)
					currentNumberStatus = NumberPartStatus{}
				}
			}
		}
	}

	var partNumberSum int
	for _, numberStatus := range numberStatuses {
		if numberStatus.IsPart {
			intNumber, err := strconv.Atoi(numberStatus.Number)
			if err != nil {
				log.Fatal("Error converting string to int")
			}

			partNumberSum += intNumber
		}
	}

	var engineWithPairsSum int
	for i, engine := range engines {
		for j, pair := range engine {
			if pair == 2 {
				engineWithPairsSum += i * j
			}
		}
	}

	return fmt.Sprintf("Part Number Sum: %d. Engine with pair sum: %d", partNumberSum, engineWithPairsSum)
}
