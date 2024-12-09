package main

import (
	"github.com/julijane/advent-of-code-2024/aoc"
)

func calcPart1(line string) int {
	fileblocks := []int{}
	gaps := []int{}

	for fileid, pos, i := 0, 0, 0; i < len(line); i, fileid = i+2, fileid+1 {
		lengthFile := int(line[i] - '0')

		for j := 0; j < lengthFile; j++ {
			fileblocks = append(fileblocks, fileid)
		}
		pos += lengthFile

		if i < len(line)-1 {
			lengthGap := int(line[i+1] - '0')
			if lengthGap >= 0 {
				for j := 0; j < lengthGap; j, pos = j+1, pos+1 {
					gaps = append(gaps, pos)
				}
			}
		}
	}

	for _, gapPos := range gaps {
		if gapPos > len(fileblocks) {
			break
		}
		fileId := fileblocks[len(fileblocks)-1]
		for pos := len(fileblocks) - 1; pos > gapPos; pos-- {
			fileblocks[pos] = fileblocks[pos-1]
		}
		fileblocks[gapPos] = fileId
	}

	sum := 0
	for pos, fileid := range fileblocks {
		sum += pos * fileid
	}

	return sum
}

type File struct {
	ID     int
	Pos    int
	Length int
}

func calcPart2(line string) int {
	files := []File{}
	gaps := []File{}

	for fileid, pos, i := 0, 0, 0; i < len(line); i, fileid = i+2, fileid+1 {
		lengthFile := int(line[i] - '0')
		files = append(files, File{ID: fileid, Pos: pos, Length: lengthFile})
		pos += lengthFile

		if i < len(line)-1 {
			lengthGap := int(line[i+1] - '0')
			gaps = append(gaps, File{ID: -1, Pos: pos, Length: lengthGap})
			pos += lengthGap
		}
	}

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
			sum += pos * file.ID
		}
	}

	return sum
}

func calc(input *aoc.Input, doPart1, doPart2 bool) (int, int) {
	line := input.PlainLines()[0]

	return calcPart1(line), calcPart2(line)
}

func main() {
	aoc.Run("sample1.txt", calc, true, true)
	aoc.Run("input.txt", calc, true, true)
}
