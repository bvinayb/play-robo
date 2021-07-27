package table

const MAX_X, MAX_Y int = 5, 5

type Table struct {
	MaxX, MaxY int
}

func NewTable() *Table {
	return &Table{MAX_X, MAX_Y}
}

func IsValidPosition(x int, y int) bool {
	table := NewTable()
	if x > table.MaxX-1 || x < 0 || y > table.MaxY-1 || y < 0 {
		return false
	}
	return true
}
