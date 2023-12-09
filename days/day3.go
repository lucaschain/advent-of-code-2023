package days

import (
	"fmt"
	"log"
	"strconv"

	"github.com/lucaschain/advent-of-code/helpers"
)

type Collision struct {
	Symbol rune
}

type CollisionMap map[string]Collision

type NumberPartStatus struct {
	Number     string
	Id         int
	Collisions map[string]Collision
}

func (n NumberPartStatus) AddCollisions(b CollisionMap) CollisionMap {
	var collisions CollisionMap
	if n.Collisions == nil {
		collisions = make(CollisionMap)
	} else {
		collisions = n.Collisions
	}
	for key, value := range b {
		collisions[key] = value
	}

	return collisions
}

func (n NumberPartStatus) IsPart() bool {
	return len(n.Collisions) > 0
}

func (n NumberPartStatus) String() string {
	return fmt.Sprintf("Number: %s, IsPart: %t, Collisions: %d", n.Number, n.IsPart(), len(n.Collisions))
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

func checkCollisions(x int, y int, lines []string) CollisionMap {
	dirs := [][]int{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}

	collisions := make(CollisionMap)
	for _, dir := range dirs {
		targetX := x + dir[0]
		targetY := y + dir[1]
		char := get(targetX, targetY, lines)
		if isSymbol(char) {
			collisionKey := fmt.Sprintf("%d.%d", targetX, targetY)
			collisions[collisionKey] = Collision{
				Symbol: char,
			}
		}
	}

	return collisions
}

func Day3() string {
	lines := helpers.Read("input/day3.txt")

	var numberStatuses []NumberPartStatus
	var currentNumberStatus NumberPartStatus
	currentId := 0

	for y, line := range lines {
		for x, char := range line {
			if isDigit(char) {
				currentNumberStatus.Number += string(char)
				currentNumberStatus.Collisions = currentNumberStatus.AddCollisions(checkCollisions(x, y, lines))
			} else {
				if currentNumberStatus.Number != "" {
					numberStatuses = append(numberStatuses, currentNumberStatus)
					currentNumberStatus = NumberPartStatus{}
					currentNumberStatus.Id = currentId
					currentId += 1
				}
			}
		}
	}

	var partNumberSum int
	engines := make(map[string][]int)
	for _, numberStatus := range numberStatuses {
		if numberStatus.IsPart() {
			intNumber, err := strconv.Atoi(numberStatus.Number)
			if err != nil {
				log.Fatal("Error converting string to int")
			}

			partNumberSum += intNumber

			for collisionKey, collision := range numberStatus.Collisions {
				if collision.Symbol == '*' {
					engines[collisionKey] = append(engines[collisionKey], intNumber)
				}
			}
		}
	}

	var engineWithPairsSum int

	for _, engine := range engines {
		if len(engine) == 2 {
			engineWithPairsSum += engine[0] * engine[1]
		}
	}

	return fmt.Sprintf("Part Number Sum: %d. Engine with pair sum: %d", partNumberSum, engineWithPairsSum)
}
