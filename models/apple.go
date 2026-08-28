package model

import "math/rand"

const AppleInitialX = 10
const AppleInitialY = 10

type Apple struct {
	X int
	Y int
}

func NewApple() *Apple {
	return &Apple{
		X: AppleInitialX,
		Y: AppleInitialY,
	}
}

// SummonApple moves the apple to a random free cell (not on the snake).
// If the board is full the apple stays where it is.
func (a *Apple) SummonApple(bannedBoxes []Box) {
	totalCells := BoardBoxColumns * BoardBoxRows

	for len(bannedBoxes) < totalCells {
		box := a.randomBox()
		if !OneToMany(box, bannedBoxes) {
			a.X = box.X
			a.Y = box.Y
			return
		}
	}
}

func (a *Apple) randomBox() Box {
	return Box{
		X: rand.Intn(BoardBoxColumns),
		Y: rand.Intn(BoardBoxRows),
	}
}

func (a *Apple) Reset() {
	a.X = AppleInitialX
	a.Y = AppleInitialY
}
