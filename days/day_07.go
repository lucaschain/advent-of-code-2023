package days

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"slices"
	"sort"

	"github.com/lucaschain/advent-of-code/helpers"
)

var powerTable = map[string]int{
	"2": 2, "3": 3, "4": 4, "5": 5,
	"6": 6, "7": 7, "8": 8, "9": 9,
	"T": 10, "J": 1, "Q": 12, "K": 13,
	"A": 14,
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	slices.Sort(a)
	slices.Sort(b)

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func isFiveOfAKind(cardCount []int, jokerCount int) bool {
	if jokerCount == 5 {
		return true
	}

	for _, count := range cardCount {
		if count+jokerCount == 5 {
			return true
		}
	}
	return false
}

func isFourOfAKind(cardCount []int, jokerCount int) bool {
	for _, count := range cardCount {
		if count+jokerCount == 4 {
			return true
		}
	}
	return false
}

func isFullHouse(cardCount []int, jokerCount int) bool {
	fullHouseSet := []int{3, 2}
	if sliceEqual(cardCount, fullHouseSet) {
		return true
	}

	if jokerCount == 1 {
		if sliceEqual(cardCount, []int{3, 1}) {
			return true
		}
		if sliceEqual(cardCount, []int{2, 2}) {
			return true
		}
	} else if jokerCount == 2 {
		if sliceEqual(cardCount, []int{2, 1}) {
			return true
		}
	}

	return false
}

func isThreeOfAKind(cardCount []int, jokerCount int) bool {
	for _, count := range cardCount {
		if count+jokerCount == 3 {
			return true
		}
	}
	return false
}

func isTwoPair(cardCount []int, jokerCount int) bool {
	if jokerCount == 1 {
		if sliceEqual(cardCount, []int{2, 1}) {
			return true
		}
		return sliceEqual(cardCount, []int{2, 2})
	}

	return sliceEqual(cardCount, []int{2, 2, 1})
}

func isPair(cardCount []int, jokerCount int) bool {
	hasJoker := jokerCount >= 1

	if hasJoker {
		return true
	}

	return slices.Contains(cardCount, 2)
}

type Hand struct {
	Cards string
	Bid   int
}

func (h Hand) UniqueCount() ([]int, int) {
	count := map[string]int{}
	jokerCount := 0

	for _, char := range h.Cards {
		if char == 'J' {
			jokerCount++
			continue
		}
		count[string(char)]++
	}

	set := []int{}

	for _, value := range count {
		set = append(set, value)
	}

	return set, jokerCount
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (h Hand) CombinationPower() int {
	uniqueCount, jokerCount := h.UniqueCount()
	power := 0

	matchers := []func([]int, int) bool{
		isFiveOfAKind,
		isFourOfAKind,
		isFullHouse,
		isThreeOfAKind,
		isTwoPair,
		isPair,
	}

	powerStart := len(matchers) + 1
	for i, matcher := range matchers {
		if matcher(uniqueCount, jokerCount) {
			powerStart -= i
			power += int(math.Pow(float64(10), float64(powerStart)))
			break
		}
	}

	return power
}

func (h Hand) HandPower() int {
	handPower := 0

	for i := 0; i < 5; i++ {
		cardPower := powerTable[h.Cards[i:i+1]]
		handPower += int(float64(cardPower) * math.Pow(float64(16), float64(7-i)))
	}

	return handPower
}

type HandList []Hand

func (h HandList) Len() int {
	return len(h)
}

func (h HandList) Less(i, j int) bool {
	if h[i].CombinationPower() != h[j].CombinationPower() {
		return h[i].CombinationPower() < h[j].CombinationPower()
	}

	return h[i].HandPower() < h[j].HandPower()
}

func (h HandList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func Day7() string {
	lines := helpers.Read("input/day7.txt")
	hands := HandList{}
	for _, line := range lines {
		hand := Hand{Cards: line[0:5], Bid: helpers.ToInt(line[6:])}
		hands = append(hands, hand)
	}

	sort.Sort(hands)

	totalWinnings := 0
	for i, hand := range hands {
		println(hand.Cards, i+1, hand.CombinationPower(), hand.HandPower())
		rank := i + 1
		totalWinnings += hand.Bid * rank
	}
	return fmt.Sprintf("Total winnings: %d", totalWinnings)
}
