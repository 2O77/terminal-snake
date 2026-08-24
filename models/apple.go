package model

import (
	"math/rand"
)

const AppleInitialX = 15
const AppleInitialY = 15

type Apple Box

func NewApple() *Apple {
	var apple *Apple

	apple = &Apple{
		X: AppleInitialX,
		Y: AppleInitialY,
	}

	return apple
}

func (a *Apple) SummonApple(bannedBoxes []Box) {
	apple := a.RandomBox()

	for _, box := range bannedBoxes {
		if OneToOne(*apple, box) {
			apple = a.RandomBox()
			break
		}
	}

	a.X = apple.X
	a.Y = apple.Y
}

func (a *Apple) RandomBox() *Box {
	x := rand.Intn(BoardBoxColumns)
	y := rand.Intn(BoardBoxRows)

	return &Box{
		X: x,
		Y: y,
	}
}
