package days

import (
	"fmt"

	"github.com/lucaschain/advent-of-code/helpers"
)

type RuneConnectionMap map[rune][]Point

var ConnectionUp = Point{X: 0, Y: -1}
var ConnectionRight = Point{X: 1, Y: 0}
var ConnectionDown = Point{X: 0, Y: 1}
var ConnectionLeft = Point{X: -1, Y: 0}

var connectionMap = RuneConnectionMap{
	'|': {ConnectionUp, ConnectionDown},
	'-': {ConnectionRight, ConnectionLeft},
	'L': {ConnectionUp, ConnectionRight},
	'J': {ConnectionUp, ConnectionLeft},
	'F': {ConnectionRight, ConnectionDown},
	'7': {ConnectionDown, ConnectionLeft},
	'.': {},
	'S': {ConnectionUp, ConnectionRight, ConnectionDown, ConnectionLeft},
}

type Point struct {
	X int
	Y int
}

func (p *Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p *Point) String() string {
	return fmt.Sprintf("P{X: %d, Y: %d}", p.X, p.Y)
}

type Tile struct {
	Position    Point
	TileType    rune
	Connections []*Tile
}

func (t *Tile) String() string {
	return fmt.Sprintf("Tile{Position: %v, TileType: %v}", t.Position, string(t.TileType))
}

func (t *Tile) FindNext(from *Tile) *Tile {
	for _, connection := range t.Connections {
		if connection.Position != from.Position {
			return connection
		}
	}

	panic("No next tile found")
}

func (t *Tile) Is(other *Tile) bool {
	return t.Position == other.Position
}

func (t *Tile) ConnectsTo(other Tile) bool {
	tileConnections := connectionMap[t.TileType]
	otherConnections := connectionMap[other.TileType]

	for _, tileConnection := range tileConnections {
		for _, otherConnection := range otherConnections {
			target := tileConnection.Add(t.Position)
			source := otherConnection.Add(other.Position)
			if target == other.Position && source == t.Position {
				return true
			}
		}
	}

	return false
}

func getStartingPoint(lines []string) Point {
	for y, line := range lines {
		for x, tile := range line {
			if tile == 'S' {
				return Point{X: x, Y: y}
			}
		}
	}

	panic("Starting point not found")
}

func getSurroundingConnections(tile *Tile, lines []string) []*Tile {
	directions := []Point{
		ConnectionUp,
		ConnectionRight,
		ConnectionDown,
		ConnectionLeft,
	}

	connections := []*Tile{}
	for _, dir := range directions {
		otherPosition := tile.Position.Add(dir)
		other := helpers.GetFromGrid(otherPosition.X, otherPosition.Y, lines)
		otherTile := Tile{Position: otherPosition, TileType: other}

		if tile.ConnectsTo(otherTile) {
			connections = append(connections, &otherTile)
		}
	}

	if len(connections) == 1 {
		for _, dir := range directions {
			otherPosition := tile.Position.Add(dir)
			other := helpers.GetFromGrid(otherPosition.X, otherPosition.Y, lines)
			otherTile := Tile{Position: otherPosition, TileType: other}

			if tile.ConnectsTo(otherTile) {
				fmt.Printf("Tile %s connects to %s\n", tile.String(), otherTile.String())
			} else {
				fmt.Printf("Tile %s does not connect to %s\n", tile.String(), otherTile.String())
			}
		}

		panic("Tile has only one connection")
	}

	return connections
}

func Day10() string {
	lines := helpers.Read("input/day10.txt")

	startingPoint := getStartingPoint(lines)
	startingTile := &Tile{
		Position: startingPoint,
		TileType: 'S',
	}
	startingTile.Connections = getSurroundingConnections(startingTile, lines)

	if len(startingTile.Connections) != 2 {
		panic("Starting point needs to have exactly 2 connections")
	}

	steps := 1
	currentA := startingTile.Connections[0]
	currentB := startingTile.Connections[1]
	previousA := startingTile
	previousB := startingTile
	for {
		if currentA.ConnectsTo(*currentB) || currentA == currentB {
			return "Found a connection"
		}
		currentA.Connections = getSurroundingConnections(currentA, lines)
		currentB.Connections = getSurroundingConnections(currentB, lines)
		steps++

		nextA := currentA.FindNext(previousA)
		nextB := currentB.FindNext(previousB)
		fmt.Printf("A: %s -> %s\n", currentA.String(), nextA.String())
		fmt.Printf("B: %s -> %s\n", currentB.String(), nextB.String())

		if nextA.Is(nextB) {
			return fmt.Sprintf("Farthest point after %d steps", steps)
		}

		previousA = currentA
		previousB = currentB

		currentA = nextA
		currentB = nextB

	}
}
