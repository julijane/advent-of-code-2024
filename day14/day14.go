package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

type Robot struct {
	StartPos aoc.Coordinate
	Velocity aoc.Coordinate
}

func (r Robot) PositionAfter(seconds, gridWidth, gridHeight int) aoc.Coordinate {
	return aoc.Coordinate{
		X: (r.StartPos.X + r.Velocity.X*seconds) % gridWidth,
		Y: (r.StartPos.Y + r.Velocity.Y*seconds) % gridHeight,
	}
}

func calc(input *aoc.Input, _, doPart2 bool, param ...any) (any, any) {
	sumPart1 := 0
	sumPart2 := 0

	params := param[0].([]interface{})

	gridWidth := params[0].(int)
	gridHeight := params[1].(int)

	robots := []Robot{}
	for _, line := range input.PlainLines() {
		values := aoc.ExtractNumbers(line)

		robots = append(robots, Robot{
			StartPos: aoc.Coordinate{
				X: values[0],
				Y: values[1],
			},
			Velocity: aoc.Coordinate{
				X: (values[2] + gridWidth) % gridWidth,
				Y: (values[3] + gridHeight) % gridHeight,
			},
		})
	}

	quadrantCount := [4]int{}

	for _, robot := range robots {
		finalPos := robot.PositionAfter(100, gridWidth, gridHeight)

		if finalPos.X < gridWidth/2 {
			if finalPos.Y < gridHeight/2 {
				quadrantCount[0]++
			} else if finalPos.Y > gridHeight/2 {
				quadrantCount[2]++
			}
		} else if finalPos.X > gridWidth/2 {
			if finalPos.Y < gridHeight/2 {
				quadrantCount[1]++
			} else if finalPos.Y > gridHeight/2 {
				quadrantCount[3]++
			}
		}

	}

	sumPart1 = quadrantCount[0] * quadrantCount[1] * quadrantCount[2] * quadrantCount[3]

	if doPart2 {
		for {
			sumPart2++
			g := aoc.NewGrid(gridWidth, gridHeight, ' ')
			for _, robot := range robots {
				pos := robot.PositionAfter(sumPart2, gridWidth, gridHeight)
				g.Set(pos, '#')
			}

			connected := 0
			unconnected := 0

			for y := 0; y < gridHeight; y++ {
			xloop:
				for x := 0; x < gridWidth; x++ {
					pos := aoc.Coordinate{X: x, Y: y}
					if g.Get(pos, '.') != '#' {
						continue
					}

					for dir := 0; dir < 4; dir++ {
						if g.Get(pos.Add(aoc.DirsStraight[dir]), ' ') == '#' {
							connected++
							continue xloop
						}
					}

					unconnected++
				}
			}

			if float64(connected) > float64(connected+unconnected)*0.7 {
				g.Print()
				break
			}

		}
	}

	return sumPart1, sumPart2
}

func main() {
	aoc.Run("sample1.txt", calc, true, false, 11, 7)
	aoc.Run("input.txt", calc, true, true, 101, 103)
}
