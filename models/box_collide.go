package model

func OneToOne(Box1 Box, Box2 Box) bool {
	if Box1.X == Box2.X && Box1.Y == Box2.Y {
		return true
	}

	return false
}

func OneToMany(Box1 Box, Boxes []Box) bool {
	for _, elem := range Boxes {
		isCollide := OneToOne(Box1, elem)
		if isCollide {
			return true
		}
	}

	return false
}

func IsOutside(Box Box, x int, y int) bool {
	if Box.X < 0 || Box.X > x {
		return true
	}

	if Box.Y < 0 || Box.Y > y {
		return true
	}

	return false
}
