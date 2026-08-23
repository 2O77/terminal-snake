package board_view

import (
	"log"
	model "terminal-snake/models"

	"github.com/gdamore/tcell/v2"
)

type ViewDeps struct {
	Apple           *model.Apple
	Snake           *model.Snake
	Screen          tcell.Screen
	BoardBoxColumns int
	BoardBoxRows    int
}

type View struct {
	apple            *model.Apple
	snake            *model.Snake
	screen           tcell.Screen
	boardCoordinates ViewBoardCoordinates
	rune             rune
	combining        []rune
}

type ViewBoardCoordinates struct {
	boardLeftX   int
	boardRightX  int
	boardTopY    int
	boardBottomY int
}

const (
	boardWidthViewCells  = model.BoardBoxColumns * 2
	boardHeightViewCells = model.BoardBoxRows
)

func NewView(deps ViewDeps) View {
	boardCoordinates := ViewBoardCoordinates{}
	const runeValue rune = 0
	var combining []rune = nil

	return View{
		deps.Apple,
		deps.Snake,
		deps.Screen,
		boardCoordinates,
		runeValue,
		combining,
	}
}

func (v View) SetView() {
	v.screen.Clear()

	viewBoardCoordinates, ok := v.viewBoardCoordinates()
	if !ok {
		log.Fatal()
	}

	v.boardCoordinates = viewBoardCoordinates

	v.setBackground()
	v.setSnake()
	v.setApple()
	// setTopbar(screen)
	v.screen.Show()
}

func (v View) setBackground() {
	unoccupiedCellStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)

	backgroundStyle := tcell.StyleDefault.
		Background(tcell.ColorDarkGreen).
		Foreground(tcell.ColorWhite)

	v.screen.SetStyle(unoccupiedCellStyle)

	x, y := v.boardCoordinates.boardLeftX, v.boardCoordinates.boardTopY

	for viewColumn := 0; viewColumn < boardWidthViewCells; viewColumn += 2 {
		for viewRow := 0; viewRow < boardHeightViewCells; viewRow += 1 {
			for index := range 2 {
				v.screen.SetContent(x+viewColumn+index, y+viewRow, v.rune, v.combining, backgroundStyle)
			}
		}
	}
}

func (v View) setSnake() {
	const runeValue rune = 0
	var combining []rune = nil

	snake := v.snake

	snakeStyle := tcell.StyleDefault.
		Background(tcell.ColorRed).
		Foreground(tcell.Color182)

	headStyle := tcell.StyleDefault.
		Background(tcell.ColorPaleVioletRed).
		Foreground(tcell.Color182)

	snakeWithViewCells := model.Snake{}

	snakeWithViewCells.Body = make([]model.Box, len(snake.Body))
	copy(snakeWithViewCells.Body, snake.Body)

	for index, value := range snakeWithViewCells.Body {
		snakeWithViewCells.Body[index].X = value.X * 2
	}

	head := snakeWithViewCells.Body[len(snake.Body)-1]
	if head.X < 0 || head.X > boardWidthViewCells || head.Y < 0 || head.Y > boardHeightViewCells {
		return
	}

	x, y := v.boardCoordinates.boardLeftX, v.boardCoordinates.boardTopY

	for index, value := range snakeWithViewCells.Body {
		if index == len(snake.Body)-1 {
			for index := range 2 {
				v.screen.SetContent(x+value.X+index, y+value.Y, runeValue, combining, headStyle)
			}
		} else {
			for index := range 2 {
				v.screen.SetContent(x+value.X+index, y+value.Y, v.rune, v.combining, snakeStyle)
			}
		}
	}
}

func (v View) setApple() {
	appleStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorDarkRed)

	boardX, boardY := v.boardCoordinates.boardLeftX, v.boardCoordinates.boardTopY
	appleWithViewCells := model.Apple{
		X: v.apple.X * 2,
		Y: v.apple.Y,
	}

	x := boardX + appleWithViewCells.X
	y := boardY + appleWithViewCells.Y

	for index := range 2 {
		v.screen.SetContent(x+index, y, v.rune, v.combining, appleStyle)
	}
}

func (v View) viewBoardCoordinates() (result ViewBoardCoordinates, ok bool) {
	w, h := v.screen.Size()

	if w < boardWidthViewCells || h < boardHeightViewCells {
		ok = false
		return
	}

	ok = true

	result.boardLeftX = (w - boardWidthViewCells) / 2
	result.boardRightX = result.boardLeftX + boardWidthViewCells - 2
	result.boardTopY = (h - boardHeightViewCells) / 2
	result.boardBottomY = result.boardTopY + boardHeightViewCells - 1

	return
}
