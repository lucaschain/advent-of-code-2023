package days

import (
	"fmt"
	"time"

	"github.com/lucaschain/advent-of-code/helpers"
)

type Node struct {
	Name     string
	Children map[byte]string
}

func (n Node) IsStart() bool {
	return n.Name[len(n.Name)-1] == 'A'
}

func (n Node) IsEnd() bool {
	return n.Name[len(n.Name)-1] == 'Z'
}

func (n Node) Next(key byte) string {
	return n.Children[key]
}

type MapNodes map[string]Node

func nodeFromLine(line string) Node {
	extracted := helpers.ExtractInfo(`(?P<Name>[A-Z]{3}) = \((?P<L>[A-Z]{3}), (?P<R>[A-Z]{3})\)`, line)

	node := Node{
		Children: make(map[byte]string),
	}
	for property := range extracted {
		if property == "Name" {
			node.Name = extracted[property]
		} else {
			node.Children[property[0]] = extracted[property]
		}
	}

	return node
}

func areAllEndingNodes(nodes MapNodes) bool {
	for _, node := range nodes {
		if !node.IsEnd() {
			return false
		}
	}

	return true
}

func Day8() string {
	lines := helpers.Read("input/day8.txt")
	timeStart := time.Now()

	instructions := lines[0]

	fmt.Println("Instructions:", instructions)

	nodes := MapNodes{}
	startingNodes := MapNodes{}
	for _, line := range lines[2:] {
		node := nodeFromLine(line)
		nodes[node.Name] = node

		if node.IsStart() {
			startingNodes[node.Name] = node
		}
	}

	fmt.Println("Starting Nodes:", startingNodes)

	var steps int
	currentNodes := startingNodes
	instructionsLen := len(instructions)
	for {
		instruction := instructions[steps%instructionsLen]
		if areAllEndingNodes(currentNodes) {
			break
		}

		nextNodes := MapNodes{}
		for _, node := range currentNodes {
			nextNode := node.Next(instruction)
			nextNodes[nextNode] = nodes[nextNode]
		}
		currentNodes = nextNodes

		steps++
	}

	fmt.Println("Time spent:", time.Since(timeStart))
	return fmt.Sprintf("Steps: %d", steps)
}
