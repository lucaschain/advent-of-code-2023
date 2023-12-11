package main

import (
	"fmt"

	"github.com/lucaschain/advent-of-code/days"
)

func main() {
	exercises := []func() string{
		days.Day1,
		days.Day2,
		days.Day3,
		days.Day4,
		days.Day5,
		days.Day6,
		days.Day7,
		days.Day8,
		days.Day9,
	}

	for day, exercise := range exercises {
		fmt.Printf("DAY %d: %s\n", day+1, exercise())
	}
}
