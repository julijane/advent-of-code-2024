package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

type Field [4]bool

func walkFrom(g *aoc.Grid, c aoc.Coordinate, fieldVal byte, walkedFields map[aoc.Coordinate]Field) (int, int) {
	sumFields := 1
	sumPerimeter := 0

	field := Field{false, false, false, false}
	walkedFields[c] = field

	for dir := 0; dir < 4; dir++ {
		dirXY := aoc.DirsStraight[dir]
		adjacent := c.Add(dirXY)
		if _, ok := walkedFields[adjacent]; ok {
			continue
		}

		if g.Get(adjacent, '~') == fieldVal {
			fields, perimeter := walkFrom(g, adjacent, fieldVal, walkedFields)
			sumFields += fields
			sumPerimeter += perimeter
			continue
		}

		field[dir] = true

		sumPerimeter++
	}

	walkedFields[c] = field

	return sumFields, sumPerimeter
}

func countSides(walkedFields map[aoc.Coordinate]Field, width, height int) int {
	numSides := 0

	for x := 0; x < width; x++ {
		isSideLeft := false
		isSideRight := false

		for y := 0; y < height; y++ {
			pos := aoc.Coordinate{X: x, Y: y}
			if field, ok := walkedFields[pos]; ok {
				if field[1] { // border on the right
					if !isSideRight {
						numSides++
					}
				}
				isSideRight = field[1]

				if field[3] { // border on the left
					if !isSideLeft {
						numSides++
					}
				}
				isSideLeft = field[3]
			} else {
				isSideLeft = false
				isSideRight = false
			}
		}
	}

	for y := 0; y < height; y++ {
		isSideTop := false
		isSideBottom := false

		for x := 0; x < width; x++ {
			pos := aoc.Coordinate{X: x, Y: y}
			if field, ok := walkedFields[pos]; ok {
				if field[0] { // border on the top
					if !isSideTop {
						numSides++
					}
				}
				isSideTop = field[0]

				if field[2] { // border on the bottom
					if !isSideBottom {
						numSides++
					}
				}
				isSideBottom = field[2]
			} else {
				isSideTop = false
				isSideBottom = false
			}
		}
	}

	return numSides
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
				walkedFields := make(map[aoc.Coordinate]Field)
				numFields, perimeterLength := walkFrom(g, pos, plantType, walkedFields)
				sumPart1 += numFields * perimeterLength

				numSides := countSides(walkedFields, g.Width, g.Height)
				sumPart2 += numFields * numSides

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
