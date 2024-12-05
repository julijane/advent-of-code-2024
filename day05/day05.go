package main

import (
	"slices"

	"github.com/julijane/advent-of-code-2024/aoc"
)

type (
	PageList      []int
	Updates       []PageList
	Rule          []int
	OrderingRules []Rule
)

func isCorrectOrder(rules OrderingRules, pageList PageList) bool {
	for i := 1; i < len(pageList); i++ {
		page := pageList[i]

		for _, rule := range rules {
			if rule[0] != page {
				continue
			}

			posSecondPage := slices.Index(pageList, rule[1])
			if posSecondPage != -1 && posSecondPage < i {
				return false
			}
		}
	}

	return true
}

func doReorder(rules OrderingRules, pageList PageList) {
	for !isCorrectOrder(rules, pageList) {
		for _, rule := range rules {
			posFirstPage := slices.Index(pageList, rule[0])
			posSecondPage := slices.Index(pageList, rule[1])

			if posFirstPage == -1 || posSecondPage == -1 {
				continue
			}

			if posSecondPage < posFirstPage {
				pageList[posFirstPage], pageList[posSecondPage] = pageList[posSecondPage], pageList[posFirstPage]
			}
		}
	}
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	blocks := input.TextBlocks()

	orderingRules := OrderingRules{}
	for _, line := range blocks[0] {
		orderingRules = append(orderingRules,
			aoc.ExtractNumbers(line),
		)
	}

	updates := Updates{}
	for _, line := range blocks[1] {
		updates = append(updates,
			aoc.ExtractNumbers(line),
		)
	}

	for _, update := range updates {
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
