package main

import (
	"slices"

	"github.com/julijane/advent-of-code-2024/aoc"
)

type TrailInfo struct {
	ReachableEndpoints aoc.Coordinates
	Rating             int
}

type TrailByStartpos map[aoc.Coordinate]TrailInfo

type TrailsByStartval map[byte]TrailByStartpos

func calc(input *aoc.Input, doPart1, doPart2 bool) (any, any) {
	grid := input.Grid()

	trailsByStartval := make(TrailsByStartval)
	trailsByStartval['9'] = make(TrailByStartpos)

	gridPositions := grid.FindMultipleAll("0123456789")

	for _, pos := range gridPositions['9'] {
		trailsByStartval['9'][pos] = TrailInfo{
			ReachableEndpoints: aoc.Coordinates{pos},
			Rating:             1,
		}
	}

	for startVal := byte('8'); startVal >= '0'; startVal-- {
		trailsByStartval[startVal] = make(TrailByStartpos)

		for _, pos := range gridPositions[startVal] {
			reachableEndpoints := aoc.Coordinates{}
			rating := 0

			for _, dir := range aoc.DirsStraight {
				adjacentPos := pos.Add(dir)

				if trails, ok := trailsByStartval[startVal+1][adjacentPos]; ok {
					for _, trail := range trails.ReachableEndpoints {
						if !slices.Contains(reachableEndpoints, trail) {
							reachableEndpoints = append(reachableEndpoints, trail)
						}
					}

					rating += trails.Rating
				}
			}

			trailsByStartval[startVal][pos] = TrailInfo{
				ReachableEndpoints: reachableEndpoints,
				Rating:             rating,
			}
		}
	}

	sumPart1 := 0
	sumPart2 := 0

	for _, trail := range trailsByStartval['0'] {
		sumPart1 += len(trail.ReachableEndpoints)
		sumPart2 += trail.Rating
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
