package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

func walkFrom(g *aoc.Grid, c aoc.Coordinate, fieldVal byte, walkedFields map[aoc.Coordinate]struct{}) (int, int, int) {
	sumFields := 1
	sumPerimeterLength := 0
	sumCorners := 0

	walkedFields[c] = struct{}{}

	for dir := 0; dir < 8; dir += 2 {
		adjacentToWalk := c.Add(aoc.DirsAll[dir])
		adjacentAfter := c.Add(aoc.DirsAll[(dir+2)%8])
		adjacentDiagonal := c.Add(aoc.DirsAll[(dir+1)%8])

		contentToWalk := g.Get(adjacentToWalk, '~')
		contentAfter := g.Get(adjacentAfter, '~')
		contentDiagonal := g.Get(adjacentDiagonal, '~')

		if (contentToWalk == fieldVal && contentAfter == fieldVal && contentDiagonal != fieldVal) ||
			(contentToWalk != fieldVal && contentAfter != fieldVal) {
			sumCorners += 1
		}

		if _, ok := walkedFields[adjacentToWalk]; ok {
			continue
		}

		if contentToWalk == fieldVal {
			fields, perimeterLength, corners := walkFrom(g, adjacentToWalk, fieldVal, walkedFields)

			sumFields += fields
			sumPerimeterLength += perimeterLength
			sumCorners += corners
		} else {
			sumPerimeterLength++
		}
	}

	return sumFields, sumPerimeterLength, sumCorners
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
	sumPart1 := 0
	sumPart2 := 0

	g := input.Grid()
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			pos := aoc.Coordinate{X: x, Y: y}
			plantType := g.Get(pos, '~')
			if plantType != '.' {
				walkedFields := make(map[aoc.Coordinate]struct{})

				numFields, perimeterLength, numCorners := walkFrom(g, pos, plantType, walkedFields)

				sumPart1 += numFields * perimeterLength
				sumPart2 += numFields * numCorners

				for fieldPos := range walkedFields {
					g.Set(fieldPos, '.')
				}
			}
		}
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
