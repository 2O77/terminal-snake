package board_view

import (
	"fmt"
	model "terminal-snake/models"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

type ViewDeps struct {
	BestScore       int
	Apple           *model.Apple
	Snake           *model.Snake
	Screen          tcell.Screen
	BoardBoxColumns int
	BoardBoxRows    int
}

type View struct {
	bestScore        int
	paused           bool
	countdown        int
	apple            *model.Apple
	snake            *model.Snake
	screen           tcell.Screen
	boardCoordinates ViewBoardCoordinates
}

type ViewBoardCoordinates struct {
	boardLeftX int
	boardTopY  int
}

const (
	boardWidthViewCells  = model.BoardBoxColumns * 2
	boardHeightViewCells = model.BoardBoxRows
)

func NewView(deps ViewDeps) View {
	return View{
		bestScore: deps.BestScore,
		apple:     deps.Apple,
		snake:     deps.Snake,
		screen:    deps.Screen,
	}
}

func (v *View) SetBestScore(score int) {
	v.bestScore = score
}

func (v *View) SetPaused(paused bool) {
	v.paused = paused
}

func (v *View) SetCountdown(countdown int) {
	v.countdown = countdown
}

func (v *View) SetView() {
	v.screen.Clear()

	viewBoardCoordinates, ok := v.viewBoardCoordinates()
	if !ok {
		// terminal too small to draw the board; skip this frame
		return
	}

	v.boardCoordinates = viewBoardCoordinates

	v.setBackground()
	v.setApple()
	v.setSnake()
	v.setStatusBar()

	v.screen.Show()
}

func (v *View) setStatusBar() {
	textStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite)

	boardCenterX := v.boardCoordinates.boardLeftX + (boardWidthViewCells / 2)

	scoreText := fmt.Sprintf("SCORE: %d", len(v.snake.Body)-model.InitialSnakeSize)

	statusText := scoreText
	bottomText := "Press space to pause"

	switch {
	case v.countdown > 0:
		statusText = scoreText + " - PAUSED"
		bottomText = "Press space to cancel, esc to exit"
	case v.paused:
		statusText = scoreText + " - PAUSED"
		bottomText = "Press space to resume, esc to exit"
	case v.snake.IsGameOver():
		statusText = fmt.Sprintf("SCORE: %d - BEST SCORE: %d - GAME OVER", len(v.snake.Body)-model.InitialSnakeSize, v.bestScore)
		bottomText = "Press 'r' to restart, esc to exit"
	}

	v.emitStr(boardCenterX-(len(statusText)/2), v.boardCoordinates.boardTopY-2, textStyle, statusText)
	v.emitStr(boardCenterX-(len(bottomText)/2), v.boardCoordinates.boardTopY+boardHeightViewCells+2, textStyle, bottomText)

	if v.countdown > 0 {
		countdownText := fmt.Sprintf("%d", v.countdown)
		countdownStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
		v.emitStr(boardCenterX-(len(countdownText)/2), v.boardCoordinates.boardTopY+boardHeightViewCells/2, countdownStyle, countdownText)
	}
}

func (v *View) setBackground() {
	backgroundStyle := tcell.StyleDefault.
		Background(tcell.ColorDarkGreen).
		Foreground(tcell.ColorWhite)

	x, y := v.boardCoordinates.boardLeftX, v.boardCoordinates.boardTopY

	for viewColumn := 0; viewColumn < boardWidthViewCells; viewColumn += 2 {
		for viewRow := 0; viewRow < boardHeightViewCells; viewRow++ {
			for index := range 2 {
				v.screen.SetContent(x+viewColumn+index, y+viewRow, 0, nil, backgroundStyle)
			}
		}
	}
}

func (v *View) setSnake() {
	snakeStyle := tcell.StyleDefault.
		Background(tcell.ColorRed).
		Foreground(tcell.Color182)

	headStyle := tcell.StyleDefault.
		Background(tcell.ColorPaleVioletRed).
		Foreground(tcell.Color182)

	x, y := v.boardCoordinates.boardLeftX, v.boardCoordinates.boardTopY
	body := v.snake.Body
	headIndex := len(body) - 1

	for index, box := range body {
		style := snakeStyle
		if index == headIndex {
			style = headStyle
		}

		for cell := range 2 {
			v.screen.SetContent(x+box.X*2+cell, y+box.Y, 0, nil, style)
		}
	}
}

func (v *View) setApple() {
	appleStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorDarkRed)

	x := v.boardCoordinates.boardLeftX + v.apple.X*2
	y := v.boardCoordinates.boardTopY + v.apple.Y

	for cell := range 2 {
		v.screen.SetContent(x+cell, y, 0, nil, appleStyle)
	}
}

func (v *View) viewBoardCoordinates() (result ViewBoardCoordinates, ok bool) {
	w, h := v.screen.Size()

	if w < boardWidthViewCells || h < boardHeightViewCells {
		return ViewBoardCoordinates{}, false
	}

	result.boardLeftX = (w - boardWidthViewCells) / 2
	result.boardTopY = (h - boardHeightViewCells) / 2

	return result, true
}

func (v *View) emitStr(x, y int, style tcell.Style, str string) {
	s := v.screen

	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		s.SetContent(x, y, c, comb, style)
		x += w
	}
}
