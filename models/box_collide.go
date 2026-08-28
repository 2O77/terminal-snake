package model

func OneToOne(first Box, second Box) bool {
	return first.X == second.X && first.Y == second.Y
}

func OneToMany(target Box, boxes []Box) bool {
	for _, box := range boxes {
		if OneToOne(target, box) {
			return true
		}
	}

	return false
}

// IsOutside reports whether box is outside a board of columns x rows.
// Valid coordinates are 0 <= X < columns and 0 <= Y < rows.
func IsOutside(box Box, columns int, rows int) bool {
	return box.X < 0 || box.X >= columns || box.Y < 0 || box.Y >= rows
}
