package days

import (
	"fmt"
	"strings"
	"sync"

	"github.com/lucaschain/advent-of-code/helpers"
)

type MapRule struct {
	DestinationRangeStart int
	SourceRangeStart      int
	RangeLength           int
}

type Map struct {
	From  string
	To    string
	Rules []MapRule
}

func (m Map) String() string {
	return fmt.Sprintf("Map from %s to %s with %d rules", m.From, m.To, len(m.Rules))
}

func (m Map) AddRule(rule MapRule) Map {
	m.Rules = append(m.Rules, rule)
	return m
}

func extractSeeds(line string, seedCallback func(int)) {
	seedsPart := strings.Trim(strings.Split(strings.Split(line, ":")[1], "|")[0], " ")
	seeds := strings.Split(seedsPart, " ")

	wg := sync.WaitGroup{}
	totalSeeds := len(seeds)
	wg.Add(totalSeeds / 2)
	for i := 0; i < totalSeeds; i += 2 {
		seedStart := helpers.ToInt(seeds[i])
		seedLength := helpers.ToInt(seeds[i+1])
		seedEnd := seedStart + seedLength

		go func(seedStart, seedLength, seedEnd int) {
			defer wg.Done()

			for seed := seedStart; seed <= seedEnd; seed++ {
				seedCallback(seed)
			}
		}(seedStart, seedLength, seedEnd)
	}
	wg.Wait()
}

func isMapStart(line string) (string, string) {
	matches := helpers.ExtractInfo(`(?P<From>[a-z]+)-to-(?P<To>[a-z]+) map:`, line)

	if len(matches) == 0 {
		return "", ""
	}

	return matches["From"], matches["To"]
}

func startNewMap(from string, to string) Map {
	return Map{
		From: strings.Trim(from, " "),
		To:   strings.Trim(to, " "),
	}
}

func appendCurrentMapIfExists(maps map[string]Map, currentMap Map) map[string]Map {
	if currentMap.From != "" && currentMap.To != "" {
		maps[currentMap.From] = currentMap
	}

	return maps
}

func extractRule(line string) MapRule {
	line = strings.Trim(line, " ")
	splitLine := strings.Split(line, " ")
	return MapRule{
		DestinationRangeStart: helpers.ToInt(splitLine[0]),
		SourceRangeStart:      helpers.ToInt(splitLine[1]),
		RangeLength:           helpers.ToInt(splitLine[2]),
	}
}

func getLocation(fromMapName string, maps map[string]Map, source int) int {
	locationMap, ok := maps[fromMapName]
	if !ok {
		panic(fmt.Sprintf("map %s not found", fromMapName))
	}

	for _, rule := range locationMap.Rules {
		sourceRangeStop := rule.SourceRangeStart + rule.RangeLength
		if source >= rule.SourceRangeStart && source < sourceRangeStop {
			offset := source - rule.SourceRangeStart
			return rule.DestinationRangeStart + offset
		}
	}

	return source
}

func followSeed(seed int, maps map[string]Map) int {
	soil := getLocation("seed", maps, seed)
	fertilizer := getLocation("soil", maps, soil)
	water := getLocation("fertilizer", maps, fertilizer)
	light := getLocation("water", maps, water)
	temperature := getLocation("light", maps, light)
	humidity := getLocation("temperature", maps, temperature)
	return getLocation("humidity", maps, humidity)
}

func Day5() string {
	lines := helpers.Read("input/day5.txt")

	maps := make(map[string]Map)
	var currentMap Map
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		from, to := isMapStart(line)
		if from != "" || to != "" {
			maps = appendCurrentMapIfExists(maps, currentMap)
			currentMap = startNewMap(from, to)
			continue
		}

		currentMap = currentMap.AddRule(extractRule(line))
	}
	maps = appendCurrentMapIfExists(maps, currentMap)

	lowestLocation := int(^uint(0) >> 1)
	seedLocator := func(seed int) {
		location := followSeed(seed, maps)
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	extractSeeds(lines[0], seedLocator)

	return fmt.Sprintf("Lowest location: %d", lowestLocation)
}
