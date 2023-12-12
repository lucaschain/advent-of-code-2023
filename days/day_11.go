package days

import (
	"fmt"
	"math"

	"github.com/lucaschain/advent-of-code/helpers"
)

var GROWTH = 1000000

func isEmptyLine(line string) bool {
	for _, char := range line {
		if char != '.' {
			return false
		}
	}
	return true
}

func getExpandingRowIndexes(lines []string) []int {
	expandingRowIndexes := []int{}
	for i, line := range lines {
		if isEmptyLine(line) {
			expandingRowIndexes = append(expandingRowIndexes, i)
		}
	}
	return expandingRowIndexes
}

func getExpandingColumnIndexes(lines []string) []int {
	expandingColumnIndexes := []int{}

	isExpandingMap := map[int]bool{}
	for i := 0; i < len(lines[0]); i++ {
		isExpandingMap[i] = true
	}

	for _, line := range lines {
		for i, char := range line {
			if char != '.' {
				isExpandingMap[i] = false
			}
		}
	}

	for i, isExpanding := range isExpandingMap {
		if isExpanding {
			expandingColumnIndexes = append(expandingColumnIndexes, i)
		}
	}

	return expandingColumnIndexes
}

func appendLines(lines []string, newLine string, times int) []string {
	for i := 0; i < times; i++ {
		lines = append(lines, newLine)
	}
	return lines
}

func GalaxySliceToIdSlice(galaxies []Galaxy) []int {
	ids := []int{}
	for _, galaxy := range galaxies {
		ids = append(ids, galaxy.Id)
	}
	return ids
}

type Galaxy struct {
	Id int
	X  int
	Y  int
}

func (g *Galaxy) DistanceTo(other *Galaxy) int {
	xDiff := g.X - other.X
	yDiff := g.Y - other.Y
	return int(math.Abs(float64(xDiff)) + math.Abs(float64(yDiff)))
}

func countLowerThan(slice []int, value int) int {
	count := 0
	for _, item := range slice {
		if item < value {
			count++
		}
	}
	return count
}

func actualPosition(x int, y int, expandingRows, expandingColumns []int) (int, int) {
	xExpansions := countLowerThan(expandingColumns, x)
	yExpansions := countLowerThan(expandingRows, y)
	newX := (x - xExpansions) + xExpansions*GROWTH
	newY := (y - yExpansions) + yExpansions*GROWTH
	return newX, newY
}

func Day11() string {
	lines := helpers.Read("input/day11.txt")

	expandingRows := getExpandingRowIndexes(lines)
	expandingColumns := getExpandingColumnIndexes(lines)

	galaxies := []Galaxy{}
	galaxyId := 1
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				actualX, actualY := actualPosition(x, y, expandingRows, expandingColumns)
				galaxies = append(galaxies, Galaxy{Id: galaxyId, X: actualX, Y: actualY})
				galaxyId++
			}
		}
	}

	pairCombinations := map[int]map[int][]Galaxy{}

	savePair := func(pair []Galaxy) bool {
		exists := false

		if pairCombinations[pair[0].Id] != nil {
			if pairCombinations[pair[0].Id][pair[1].Id] != nil {
				exists = true
			}
		}

		if pairCombinations[pair[1].Id] != nil {
			if pairCombinations[pair[1].Id][pair[0].Id] != nil {
				exists = true
			}
		}

		if exists {
			return false
		}

		if pair[0].Id < pair[1].Id {
			pairCombinations[pair[0].Id][pair[1].Id] = pair
		} else {
			pairCombinations[pair[1].Id][pair[0].Id] = pair
		}

		return true
	}

	totalPairs := 0
	for _, thisGalaxy := range galaxies {
		pairCombinations[thisGalaxy.Id] = map[int][]Galaxy{}
		for _, otherGalaxy := range galaxies {
			if thisGalaxy.Id == otherGalaxy.Id {
				continue
			}

			newPair := []Galaxy{thisGalaxy, otherGalaxy}
			if !savePair(newPair) {
				totalPairs++
			}
		}
	}

	shortestPathSum := 0
	for _, combinations := range pairCombinations {
		for _, pair := range combinations {
			shortestPathSum += pair[0].DistanceTo(&pair[1])
		}
	}

	return fmt.Sprintf("Shortest Path Sum %d, Total Pairs %d, Galaxies: %d", shortestPathSum, totalPairs, len(galaxies))
}
