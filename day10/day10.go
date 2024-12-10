package main

import (
	"slices"
	"strconv"

	"github.com/julijane/advent-of-code-2024/aoc"
)

type TrailInfo struct {
	ReachableNines aoc.Coordinates
	Rating         int
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
	grid := input.Grid()

	trailsByStartval := make(map[int]map[aoc.Coordinate]TrailInfo)
	trailsByStartval[9] = make(map[aoc.Coordinate]TrailInfo)

	ninePos := grid.FindAll('9')
	for _, pos := range ninePos {
		trailsByStartval[9][pos] = TrailInfo{
			ReachableNines: aoc.Coordinates{pos},
			Rating:         1,
		}
	}

	for startVal := 8; startVal >= 0; startVal-- {
		fieldVal := strconv.Itoa(startVal)[0]
		fieldPos := grid.FindAll(fieldVal)

		trailsByStartval[startVal] = make(map[aoc.Coordinate]TrailInfo)

		for _, pos := range fieldPos {
			reachable := aoc.Coordinates{}
			posRating := 0

			for _, dir := range aoc.DirsStraight {
				adjacentPos := pos.Add(dir)

				if trails, ok := trailsByStartval[startVal+1][adjacentPos]; ok {
					for _, trail := range trails.ReachableNines {
						if !slices.Contains(reachable, trail) {
							reachable = append(reachable, trail)
						}
					}

					posRating += trails.Rating
				}
			}

			trailsByStartval[startVal][pos] = TrailInfo{
				ReachableNines: reachable,
				Rating:         posRating,
			}
		}
	}

	sumPart1 := 0
	sumPart2 := 0

	for _, trail := range trailsByStartval[0] {
		sumPart1 += len(trail.ReachableNines)
		sumPart2 += trail.Rating
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
