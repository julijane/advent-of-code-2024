package aoc

type Coordinate struct {
	X, Y int
}

var (
	DirUL = Coordinate{X: -1, Y: -1}
	DirU  = Coordinate{X: 0, Y: -1}
	DirUR = Coordinate{X: 1, Y: -1}
	DirR  = Coordinate{X: 1, Y: 0}
	DirDR = Coordinate{X: 1, Y: 1}
	DirD  = Coordinate{X: 0, Y: 1}
	DirDL = Coordinate{X: -1, Y: 1}
	DirL  = Coordinate{X: -1, Y: 0}

	DirsStraight = []Coordinate{DirU, DirR, DirD, DirL}
	DirsDiagonal = []Coordinate{DirUL, DirUR, DirDR, DirDL}
	DirsAll      = []Coordinate{DirU, DirUR, DirR, DirDR, DirD, DirDL, DirL, DirUL}
)

func (c Coordinate) Same(other Coordinate) bool {
	return c.X == other.X && c.Y == other.Y
}

func (c Coordinate) Add(other Coordinate) Coordinate {
	return Coordinate{
		X: c.X + other.X,
		Y: c.Y + other.Y,
	}
}

func (c Coordinate) Subtract(other Coordinate) Coordinate {
	return Coordinate{
		X: c.X - other.X,
		Y: c.Y - other.Y,
	}
}

func (c Coordinate) AddXY(x, y int) Coordinate {
	return Coordinate{
		X: c.X + x,
		Y: c.Y + y,
	}
}

func (c Coordinate) Copy() Coordinate {
	return Coordinate{
		X: c.X,
		Y: c.Y,
	}
}

func (c Coordinate) Up() Coordinate {
	return c.Add(DirU)
}

func (c Coordinate) Down() Coordinate {
	return c.Add(DirD)
}

func (c Coordinate) Left() Coordinate {
	return c.Add(DirL)
}

func (c Coordinate) Right() Coordinate {
	return c.Add(DirR)
}

func (c Coordinate) UpLeft() Coordinate {
	return c.Add(DirUL)
}

func (c Coordinate) UpRight() Coordinate {
	return c.Add(DirUR)
}

func (c Coordinate) DownLeft() Coordinate {
	return c.Add(DirDL)
}

func (c Coordinate) DownRight() Coordinate {
	return c.Add(DirDR)
}

func (c Coordinate) Move(direction int) Coordinate {
	switch direction {
	case 0:
		return c.Up()
	case 1:
		return c.Right()
	case 2:
		return c.Down()
	case 3:
		return c.Left()
	}
	return c
}

func (c Coordinate) MoveBy(direction int, amount int) Coordinate {
	switch direction {
	case 0:
		return c.AddXY(0, -amount)
	case 1:
		return c.AddXY(amount, 0)
	case 2:
		return c.AddXY(0, amount)
	case 3:
		return c.AddXY(-amount, 0)
	}
	return c
}

type Coordinates []Coordinate

func (cs Coordinates) Includes(c Coordinate) bool {
	for _, test := range cs {
		if test.X == c.X && test.Y == c.Y {
			return true
		}
	}
	return false
}
