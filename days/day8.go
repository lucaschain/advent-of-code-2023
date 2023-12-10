package days

import (
	"fmt"
	"time"

	"github.com/lucaschain/advent-of-code/helpers"
)

type Node struct {
	Name     string
	Children []string
	IsStart  bool
	IsEnd    bool
}

func (n Node) Next(key byte) string {
	if key == 'L' {
		return n.Children[0]
	} else {
		return n.Children[1]
	}
}

type MapNodes map[string]Node

func nodeFromLine(line string) Node {
	extracted := helpers.ExtractInfo(`(?P<Name>[A-Z]{3}) = \((?P<L>[A-Z]{3}), (?P<R>[A-Z]{3})\)`, line)

	nodeName := extracted["Name"]
	lastChar := nodeName[len(nodeName)-1]
	node := Node{
		Children: []string{extracted["L"], extracted["R"]},
		Name:     nodeName,
		IsStart:  (lastChar == 'A'),
		IsEnd:    (lastChar == 'Z'),
	}
	return node
}

func areAllEndingNodes(nodes MapNodes) bool {
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
	startingNodes := MapNodes{}
	for _, line := range lines[2:] {
		node := nodeFromLine(line)
		nodes[node.Name] = node

		if node.IsStart {
			startingNodes[node.Name] = node
		}
	}

	fmt.Println("Starting Nodes:", startingNodes)

	var steps int
	currentNodes := startingNodes
	instructionsLen := len(instructions)
	for {
		if steps%1000000 == 0 {
			fmt.Println("Steps:", steps)
		}
		instruction := instructions[steps%instructionsLen]
		nextNodes := MapNodes{}
		for _, node := range currentNodes {
			nextNode := node.Next(instruction)
			nextNodes[nextNode] = nodes[nextNode]
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
