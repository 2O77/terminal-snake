package main

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

type Snake struct {
	ColStart int
	ColEnd   int
	RowStart int
	RowEnd   int
}

// 2 column wide 5 row long
type SnakePiece struct {
	RowEnd int
	ColEnd int
}

func moveSnakeUp(screen tcell.Screen, headStyle tcell.Style, bodyStyle tcell.Style, snake *[]SnakePiece) {
	resetScreen(screen)
	if (*snake)[0].ColEnd <= (*snake)[1].ColEnd {
		for i := len(*snake) - 1; i > 0; i-- {
			(*snake)[i].ColEnd = (*snake)[i-1].ColEnd
			(*snake)[i].RowEnd = (*snake)[i-1].RowEnd
		}
		(*snake)[0].ColEnd -= 2
	}
	renderSnake(screen, headStyle, bodyStyle, snake)
}

func moveSnakeDown(screen tcell.Screen, headStyle tcell.Style, bodyStyle tcell.Style, snake *[]SnakePiece) {
	resetScreen(screen)
	if (*snake)[0].ColEnd >= (*snake)[1].ColEnd {
		for i := len(*snake) - 1; i > 0; i-- {
			(*snake)[i].ColEnd = (*snake)[i-1].ColEnd
			(*snake)[i].RowEnd = (*snake)[i-1].RowEnd
		}
		(*snake)[0].ColEnd += 2
	}
	renderSnake(screen, headStyle, bodyStyle, snake)
}

func moveSnakeLeft(screen tcell.Screen, headStyle tcell.Style, bodyStyle tcell.Style, snake *[]SnakePiece) {
	resetScreen(screen)
	if (*snake)[0].RowEnd <= (*snake)[1].RowEnd {
		for i := len(*snake) - 1; i > 0; i-- {
			(*snake)[i].ColEnd = (*snake)[i-1].ColEnd
			(*snake)[i].RowEnd = (*snake)[i-1].RowEnd
		}
		(*snake)[0].RowEnd -= 5
	}
	renderSnake(screen, headStyle, bodyStyle, snake)
}

func moveSnakeRight(screen tcell.Screen, headStyle tcell.Style, bodyStyle tcell.Style, snake *[]SnakePiece) {
	resetScreen(screen)
	if (*snake)[0].RowEnd >= (*snake)[1].RowEnd {
		for i := len(*snake) - 1; i > 0; i-- {
			(*snake)[i].ColEnd = (*snake)[i-1].ColEnd
			(*snake)[i].RowEnd = (*snake)[i-1].RowEnd
		}
		(*snake)[0].RowEnd += 5
	}
	renderSnake(screen, headStyle, bodyStyle, snake)
}

func resetScreen(screen tcell.Screen) {
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.Color100)

	for row := 5; row <= 40; row++ {
		for col := 10; col <= 200; col++ {
			screen.SetContent(col, row, ' ', nil, boxStyle)
		}
	}

}

func renderSnake(screen tcell.Screen, headStyle tcell.Style, bodyStyle tcell.Style, snake *[]SnakePiece) {
	for i := 0; i < len(*snake); i++ {
		p := (*snake)[i]

		if i == 0 {
			for row := p.ColEnd - 1; row <= p.ColEnd; row++ {
				for col := p.RowEnd - 4; col <= p.RowEnd; col++ {
					screen.SetContent(col, row, ' ', nil, headStyle)
				}
			}
		} else {
			for row := p.ColEnd - 1; row <= p.ColEnd; row++ {
				for col := p.RowEnd - 4; col <= p.RowEnd; col++ {
					screen.SetContent(col, row, ' ', nil, bodyStyle)
				}
			}
		}
	}
}

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	snakeBodyStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	snakeHeadStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorDarkGray)
	appleStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)

	pieces := []SnakePiece{
		{29, 6}, {24, 6}, {19, 6}, {14, 6},
	}

	renderSnake(s, snakeHeadStyle, snakeBodyStyle, &pieces)
	resetScreen(s)

	for row := 11; row <= 12; row++ {
		for col := 15; col <= 19; col++ {
			s.SetContent(col, row, ' ', nil, appleStyle)
		}
	}

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
				moveSnakeUp(s, snakeHeadStyle, snakeBodyStyle, &pieces)
			}
			if ev.Key() == tcell.KeyDown {
				moveSnakeDown(s, snakeHeadStyle, snakeBodyStyle, &pieces)
			}
			if ev.Key() == tcell.KeyLeft {
				moveSnakeLeft(s, snakeHeadStyle, snakeBodyStyle, &pieces)
			}
			if ev.Key() == tcell.KeyRight {
				moveSnakeRight(s, snakeHeadStyle, snakeBodyStyle, &pieces)
			}

		}
	}

}
