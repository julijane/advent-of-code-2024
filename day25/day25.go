package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

func calc(input *aoc.Input, _, doPart2 bool, param ...any) (any, any) {
	resultPart1 := 0
	resultPart2 := 0

	locks := [][5]int{}
	keys := [][5]int{}

	for _, block := range input.TextBlocks() {
		grid := aoc.NewGridFromStrings(block)

		if grid.Get(aoc.Coordinate{X: 0, Y: 0}, '~') == '#' {
			// lock
			lock := [5]int{}

			for col := range 5 {
				for row := 1; row < 7 && grid.Get(aoc.Coordinate{X: col, Y: row}, '~') == '#'; row++ {
					lock[col]++
				}
			}

			locks = append(locks, lock)
		} else {
			// key
			key := [5]int{}

			for col := range 5 {
				for row := 5; row >= 0 && grid.Get(aoc.Coordinate{X: col, Y: row}, '~') == '#'; row-- {
					key[col]++
				}
			}

			keys = append(keys, key)
		}
	}

	for _, lock := range locks {
	keyloop:
		for _, key := range keys {

			for col := range 5 {
				if lock[col]+key[col] > 5 {
					continue keyloop
				}
			}

			resultPart1++
		}
	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
