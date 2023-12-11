package days

import (
	"fmt"

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

	stepsForEachGhost := []int{}
	for _, node := range startingNodes {
		stepsForEachGhost = append(stepsForEachGhost, processNode(node, instructions))
	}
	result := helpers.LcmSlice(stepsForEachGhost)

	return fmt.Sprintf("Steps: %d", result)
}

func processNode(startingNode *Node, instructions string) int {
	var steps int
	currentNode := startingNode
	instructionsLen := len(instructions)
	for {
		instruction := instructions[steps%instructionsLen]
		next := currentNode.Next(instruction)

		if next == nil {
			panic(fmt.Sprintf("Node %s has no next node", currentNode.Name))
		}

		currentNode = next

		steps++
		if currentNode.IsEnd {
			return steps
		}
	}
}
