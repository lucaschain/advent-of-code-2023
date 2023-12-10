package days

import (
	"fmt"
	"math"
	"time"

	"github.com/lucaschain/advent-of-code/helpers"
)

type Node struct {
	Name     string
	Children []string
	Left     *Node
	Right    *Node
	IsStart  bool
	IsEnd    bool
}

func (n Node) Next(key byte) *Node {
	if key == 'L' {
		return n.Left
	}
	return n.Right
}

type MapNodes map[string]*Node

func nodeFromLine(line string) Node {
	extracted := helpers.ExtractInfo(`(?P<Name>[A-Z]{3}) = \((?P<L>[A-Z]{3}), (?P<R>[A-Z]{3})\)`, line)

	nodeName := extracted["Name"]
	lastChar := nodeName[len(nodeName)-1]
	node := Node{
		Name:     nodeName,
		IsStart:  (lastChar == 'A'),
		IsEnd:    (lastChar == 'Z'),
		Children: []string{extracted["L"], extracted["R"]},
	}
	return node
}

func areAllEndingNodes(nodes []*Node) bool {
	for _, node := range nodes {
		if !node.IsEnd {
			return false
		}
	}

	return true
}

func Day8() string {
	lines := helpers.Read("input/day8.txt")
	timeStart := time.Now()

	instructions := lines[0]
	intInstructions := []int{}
	for _, instruction := range instructions {
		if instruction == 'L' {
			intInstructions = append(intInstructions, 0)
		} else if instruction == 'R' {
			intInstructions = append(intInstructions, 1)
		} else {
			panic("Invalid instruction")
		}
	}

	nodes := MapNodes{}
	startingNodes := []*Node{}
	for _, line := range lines[2:] {
		node := nodeFromLine(line)
		nodes[node.Name] = &node

		if node.IsStart {
			startingNodes = append(startingNodes, &node)
		}
	}

	for _, node := range nodes {
		leftNode, rightNode := nodes[node.Children[0]], nodes[node.Children[1]]
		node.Left = leftNode
		node.Right = rightNode
	}

	var steps int
	currentNodes := startingNodes
	instructionsLen := len(instructions)
	for {
		if steps%10000000 == 0 {
			stepMilestone := math.Pow(10, 12)
			sinceStart := time.Since(timeStart)
			timePassed := time.Duration(steps + 1)
			stepsPerMinute := float64(timePassed) / sinceStart.Minutes()
			timeToFinishInHours := (stepMilestone / stepsPerMinute) / 60
			fmt.Printf(
				`

Step %d
  Time Spent %s
  Estimate to 14 digits: %f hours`,
				steps,
				sinceStart,
				timeToFinishInHours,
			)
		}
		instruction := instructions[steps%instructionsLen]
		nextNodes := []*Node{}
		for _, node := range currentNodes {
			next := node.Next(instruction)

			if next == nil {
				panic(fmt.Sprintf("Node %s has no next node", node.Name))
			}

			nextNodes = append(nextNodes, next)
		}
		currentNodes = nextNodes

		if areAllEndingNodes(currentNodes) {
			break
		}
		steps++
	}

	fmt.Println("Time spent:", time.Since(timeStart))
	return fmt.Sprintf("Steps: %d", steps)
}
