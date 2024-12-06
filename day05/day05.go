package main

import (
	"slices"

	"github.com/julijane/advent-of-code-2024/aoc"
)

func isCorrectOrder(rules [][]int, pageList []int) bool {
	for _, rule := range rules {
		posFirstPage := slices.Index(pageList, rule[0])
		posSecondPage := slices.Index(pageList, rule[1])

		if posFirstPage != -1 && posSecondPage != -1 && posFirstPage > posSecondPage {
			return false
		}
	}
	return true
}

func doReorder(rules [][]int, pageList []int) {
	slices.SortFunc(pageList, func(i, j int) int {
		for _, rule := range rules {
			if rule[0] == i && rule[1] == j {
				return -1
			} else if rule[0] == j && rule[1] == i {
				return 1
			}
		}
		return 0
	})
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	blocks := input.TextBlocks()

	orderingRules := make([][]int, 0)
	for _, line := range blocks[0] {
		orderingRules = append(orderingRules, aoc.ExtractNumbers(line))
	}

	for _, line := range blocks[1] {
		update := aoc.ExtractNumbers(line)

		if isCorrectOrder(orderingRules, update) {
			sumPart1 += update[len(update)/2]
		} else {
			doReorder(orderingRules, update)
			sumPart2 += update[len(update)/2]
		}
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
