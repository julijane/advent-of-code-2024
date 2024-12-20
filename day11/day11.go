package main

import (
	"strconv"

	"github.com/julijane/advent-of-code-2024/aoc"
)

type StoneLevel struct {
	Stone int
	Level int
}

var cache map[StoneLevel]int

func calcForStoneAtLevel(stone StoneLevel, maxLevel int) int {
	if stone.Level > maxLevel {
		return 1
	}

	if result, ok := cache[stone]; ok {
		return result
	}

	newStones := []StoneLevel{}

	if stone.Stone == 0 {
		newStones = append(newStones,
			StoneLevel{Stone: 1, Level: stone.Level + 1})
	} else {
		s := strconv.Itoa(stone.Stone)
		if len(s)%2 == 0 {
			newStones = append(newStones,
				StoneLevel{Stone: aoc.Atoi(s[:len(s)/2]), Level: stone.Level + 1},
				StoneLevel{Stone: aoc.Atoi(s[len(s)/2:]), Level: stone.Level + 1},
			)
		} else {
			newStones = append(newStones,
				StoneLevel{Stone: stone.Stone * 2024, Level: stone.Level + 1},
			)
		}
	}

	result := 0
	for _, newStone := range newStones {
		result += calcForStoneAtLevel(newStone, maxLevel)
	}

	cache[stone] = result
	return result
}

func calc(input *aoc.Input, _, _ bool, _ ...any) (any, any) {
	stones := aoc.ExtractNumbers(input.PlainLines()[0])

	sumPart1 := 0
	cache = make(map[StoneLevel]int)
	for _, stone := range stones {
		sumPart1 += calcForStoneAtLevel(StoneLevel{Stone: stone, Level: 1}, 25)
	}

	sumPart2 := 0
	cache = make(map[StoneLevel]int)
	for _, stone := range stones {
		sumPart2 += calcForStoneAtLevel(StoneLevel{Stone: stone, Level: 1}, 75)
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
