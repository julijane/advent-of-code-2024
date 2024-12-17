package main

import (
	"slices"

	"github.com/julijane/advent-of-code-2024/aoc"
)

type FileBlocks struct {
	FileID int
	Pos    int
	Length int
}

func defrag(files []FileBlocks, gaps []FileBlocks) int {
	for i := len(files) - 1; i >= 0; i-- {
		if len(gaps) == 0 {
			break
		}

		for j, gap := range gaps {
			if gap.Pos > files[i].Pos {
				break
			}

			if gap.Length >= files[i].Length {
				files[i].Pos = gap.Pos
				if gap.Length > files[i].Length {
					gaps[j].Pos += files[i].Length
					gaps[j].Length -= files[i].Length
				} else {
					gaps = append(gaps[:j], gaps[j+1:]...)
				}
				break
			}
		}
	}

	sum := 0
	for _, file := range files {
		for pos := file.Pos; pos < file.Pos+file.Length; pos++ {
			sum += pos * file.FileID
		}
	}

	return sum
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (any, any) {
	line := input.PlainLines()[0]

	gapsPart2 := []FileBlocks{}
	filesPart2 := []FileBlocks{}

	for fileid, pos, i := 0, 0, 0; i < len(line); i, fileid = i+2, fileid+1 {
		lengthFile := int(line[i] - '0')
		filesPart2 = append(filesPart2, FileBlocks{FileID: fileid, Pos: pos, Length: lengthFile})
		pos += lengthFile

		if i < len(line)-1 {
			lengthGap := int(line[i+1] - '0')
			gapsPart2 = append(gapsPart2, FileBlocks{FileID: -1, Pos: pos, Length: lengthGap})
			pos += lengthGap
		}
	}

	// Part 1 is just a special case of Part 2 where we split the files into blocks of just 1 length

	filesPart1 := []FileBlocks{}
	gapsPart1 := slices.Clone(gapsPart2)

	for _, file := range filesPart2 {
		for j := 0; j < file.Length; j++ {
			filesPart1 = append(filesPart1, FileBlocks{FileID: file.FileID, Pos: file.Pos + j, Length: 1})
		}
	}

	//

	return defrag(filesPart1, gapsPart1), defrag(filesPart2, gapsPart2)
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
