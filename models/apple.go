package model

import "math/rand"

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

func (a *Apple) SummonApple() {
	x := rand.Intn(BoardBoxColumns)
	y := rand.Intn(BoardBoxRows)

	a.X = x
	a.Y = y
}
