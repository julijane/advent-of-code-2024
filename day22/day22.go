package main

import (
	"fmt"

	"github.com/julijane/advent-of-code-2024/aoc"
)

func calc(input *aoc.Input, _, doPart2 bool, param ...any) (any, any) {
	resultPart1 := 0
	resultPart2 := 0

	cache := make(map[int]int)
	bananasPerSeq := make(map[string]int)

	for _, line := range input.Lines {
		secret := aoc.ExtractNumbers(line.Data)[0]

		bananaChanges := make([]int, 4)
		bananaChanges[3] = secret % 10

		seenSeq := make(map[string]struct{})

		for step := 0; step < 2000; step++ {
			oldSecret := secret

			if val, ok := cache[oldSecret]; ok {
				secret = val
			} else {

				newval := secret << 6 // * 64
				secret = (secret ^ newval) & 0xFFFFFF

				newval = secret >> 5 // / 32
				secret = (secret ^ newval) & 0xFFFFFF

				newval = secret << 11 // * 2048
				secret = (secret ^ newval) & 0xFFFFFF

				cache[oldSecret] = secret
			}

			for i := range 3 {
				bananaChanges[i] = bananaChanges[i+1]
			}

			bananas := secret % 10
			bananaChange := bananas - oldSecret%10
			bananaChanges[3] = bananaChange

			if step >= 2 {
				changesKey := fmt.Sprintf("%d,%d,%d,%d", bananaChanges[0], bananaChanges[1], bananaChanges[2], bananaChanges[3])

				if _, ok := seenSeq[changesKey]; ok {
					continue
				}
				seenSeq[changesKey] = struct{}{}

				if _, ok := bananasPerSeq[changesKey]; !ok {
					bananasPerSeq[changesKey] = 0
				}

				bananasPerSeq[changesKey] += bananas
			}
		}

		resultPart1 += secret
	}

	if !doPart2 {
		return resultPart1, nil
	}

	for seq := range bananasPerSeq {
		if bananasPerSeq[seq] > resultPart2 {
			resultPart2 = bananasPerSeq[seq]
		}
	}

	return resultPart1, resultPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, false)
	aoc.Run("sample2.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
