package main

import (
	"log"
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

// 2 column wide 5 row long
type Piece struct {
	RowEnd int
	ColEnd int
}

type Game struct {
	screen        tcell.Screen
	snake         *[]Piece
	apple         *Piece
	isFirstRender *bool
}

func (g Game) moveSnakeUp(headStyle tcell.Style, bodyStyle tcell.Style) {
	g.renderScreen()
	if (*(g.snake))[0].ColEnd <= (*(g.snake))[1].ColEnd {
		for i := len(*(g.snake)) - 1; i > 0; i-- {
			(*(g.snake))[i].ColEnd = (*(g.snake))[i-1].ColEnd
			(*(g.snake))[i].RowEnd = (*(g.snake))[i-1].RowEnd
		}
		(*(g.snake))[0].ColEnd -= 2
	}
	g.renderSnake(headStyle, bodyStyle)
}

func (g Game) moveSnakeDown(headStyle tcell.Style, bodyStyle tcell.Style) {
	g.renderScreen()
	if (*(g.snake))[0].ColEnd >= (*(g.snake))[1].ColEnd {
		for i := len(*(g.snake)) - 1; i > 0; i-- {
			(*(g.snake))[i].ColEnd = (*(g.snake))[i-1].ColEnd
			(*(g.snake))[i].RowEnd = (*(g.snake))[i-1].RowEnd
		}
		(*(g.snake))[0].ColEnd += 2
	}
	g.renderSnake(headStyle, bodyStyle)
}

func (g Game) moveSnakeLeft(headStyle tcell.Style, bodyStyle tcell.Style) {
	g.renderScreen()
	if (*(g.snake))[0].RowEnd <= (*(g.snake))[1].RowEnd {
		for i := len(*(g.snake)) - 1; i > 0; i-- {
			(*(g.snake))[i].ColEnd = (*(g.snake))[i-1].ColEnd
			(*(g.snake))[i].RowEnd = (*(g.snake))[i-1].RowEnd
		}
		(*(g.snake))[0].RowEnd -= 5
	}
	g.renderSnake(headStyle, bodyStyle)
}

func (g Game) moveSnakeRight(headStyle tcell.Style, bodyStyle tcell.Style) {
	g.renderScreen()
	if (*(g.snake))[0].RowEnd >= (*(g.snake))[1].RowEnd {
		for i := len(*(g.snake)) - 1; i > 0; i-- {
			(*(g.snake))[i].ColEnd = (*(g.snake))[i-1].ColEnd
			(*(g.snake))[i].RowEnd = (*(g.snake))[i-1].RowEnd
		}
		(*(g.snake))[0].RowEnd += 5
	}
	g.renderSnake(headStyle, bodyStyle)
}

func (g Game) renderScreen() {
	snake := g.snake
	apple := g.apple
	isFirstRender := g.isFirstRender
	screen := g.screen

	if ((*snake)[0].ColEnd == apple.ColEnd && (*snake)[0].RowEnd == apple.RowEnd) || *isFirstRender {
		g.generateApple()
	}
	appleStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.Color100)

	for row := 6; row <= 40; row++ {
		for col := 10; col <= 200; col++ {
			screen.SetContent(col, row, ' ', nil, boxStyle)
		}
	}

	for row := apple.RowEnd - 1; row <= apple.RowEnd; row++ {
		for col := apple.ColEnd - 4; col <= apple.ColEnd; col++ {
			screen.SetContent(col, row, ' ', nil, appleStyle)
		}
	}

	*isFirstRender = false
}

func (g Game) renderSnake(headStyle tcell.Style, bodyStyle tcell.Style) {
	snake := g.snake
	screen := g.screen
	for i := 0; i < len(*snake); i++ {
		p := (*snake)[i]

		if i == 0 {
			for col := p.ColEnd - 1; col <= p.ColEnd; col++ {
				for row := p.RowEnd - 4; row <= p.RowEnd; row++ {
					screen.SetContent(row, col, ' ', nil, headStyle)
				}
			}
		} else {
			for col := p.ColEnd - 1; col <= p.ColEnd; col++ {
				for row := p.RowEnd - 4; row <= p.RowEnd; row++ {
					screen.SetContent(row, col, ' ', nil, bodyStyle)
				}
			}
		}
	}
}

func (g Game) generateApple() {
	colEnd := RandomInt(10, 200)
	rowEnd := RandomInt(5, 40)

	*g.apple = Piece{
		ColEnd: colEnd,
		RowEnd: rowEnd,
	}
}

func NewGame(g Game) *Game {
	return &g
}

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	snakeBodyStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	snakeHeadStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorDarkGray)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)

	snake := []Piece{
		{29, 8}, {24, 8}, {19, 8}, {14, 8},
	}

	isFirstRender := true

	apple := Piece{}

	g := NewGame(Game{
		snake:         &snake,
		isFirstRender: &isFirstRender,
		apple:         &apple,
		screen:        s,
	})

	g.renderScreen()
	g.renderSnake(snakeHeadStyle, snakeBodyStyle)

	quit := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	// render screen
	for {
		s.Show()
		ev := s.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				return
			}

			if ev.Key() == tcell.KeyUp {
				g.moveSnakeUp(snakeHeadStyle, snakeBodyStyle)
			}
			if ev.Key() == tcell.KeyDown {
				g.moveSnakeDown(snakeHeadStyle, snakeBodyStyle)
			}
			if ev.Key() == tcell.KeyLeft {
				g.moveSnakeLeft(snakeHeadStyle, snakeBodyStyle)
			}
			if ev.Key() == tcell.KeyRight {
				g.moveSnakeRight(snakeHeadStyle, snakeBodyStyle)
			}

		}
	}

}

func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
