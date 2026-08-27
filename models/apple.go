package model

import (
	"math/rand"
)

const AppleInitialX = 10
const AppleInitialY = 10

type Apple Box

func NewApple() *Apple {
	apple := &Apple{}

	apple = &Apple{
		X: AppleInitialX,
		Y: AppleInitialY,
	}

	return apple
}

func (a *Apple) SummonApple(bannedBoxes []Box) {
	apple := a.randomBox()

	for _, box := range bannedBoxes {
		if OneToOne(*apple, box) {

			// it is recursive, be careful
			a.SummonApple(bannedBoxes)
			return
		}
	}

	a.X = apple.X
	a.Y = apple.Y
}

func (a *Apple) randomBox() *Box {
	x := rand.Intn(BoardBoxColumns)
	y := rand.Intn(BoardBoxRows)

	return &Box{
		X: x,
		Y: y,
	}
}

func (a *Apple) Reset() {
	*a = Apple{
		AppleInitialX,
		AppleInitialY,
	}
}
