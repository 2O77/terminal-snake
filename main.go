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

func renderSnake(screen tcell.Screen, style tcell.Style, snake *Snake) {
	for row := snake.RowStart; row <= snake.RowEnd; row++ {
		for col := snake.ColStart; col <= snake.ColEnd; col++ {
			screen.SetContent(col, row, ' ', nil, style)
		}
	}
}

func moveSnakeUp(screen tcell.Screen, style tcell.Style, snake *Snake) {
	resetScreen(screen)
	snake.RowStart -= 2
	snake.RowEnd -= 2

	for row := snake.RowStart; row <= snake.RowEnd; row++ {
		for col := snake.ColStart; col <= snake.ColEnd; col++ {
			screen.SetContent(col, row, ' ', nil, style)
		}
	}
}

func moveSnakeDown(screen tcell.Screen, style tcell.Style, snake *Snake) {
	resetScreen(screen)
	snake.RowStart += 2
	snake.RowEnd += 2

	for row := snake.RowStart; row <= snake.RowEnd; row++ {
		for col := snake.ColStart; col <= snake.ColEnd; col++ {
			screen.SetContent(col, row, ' ', nil, style)
		}
	}
}

func moveSnakeLeft(screen tcell.Screen, style tcell.Style, snake *Snake) {
	resetScreen(screen)

	snake.ColStart -= 5
	snake.ColEnd -= 5

	for row := snake.RowStart; row <= snake.RowEnd; row++ {
		for col := snake.ColStart; col <= snake.ColEnd; col++ {
			screen.SetContent(col, row, ' ', nil, style)
		}
	}
}

func moveSnakeRight(screen tcell.Screen, style tcell.Style, snake *Snake) {
	resetScreen(screen)

	snake.ColStart += 5
	snake.ColEnd += 5

	for row := snake.RowStart; row <= snake.RowEnd; row++ {
		for col := snake.ColStart; col <= snake.ColEnd; col++ {
			screen.SetContent(col, row, ' ', nil, style)
		}
	}
}

func getSnakeHead(snake *Snake) {

}

func resetScreen(screen tcell.Screen) {
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.Color100)

	for row := 5; row <= 40; row++ {
		for col := 10; col <= 200; col++ {
			screen.SetContent(col, row, ' ', nil, boxStyle)
		}
	}
}

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.Color100)
	snakeStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)
	appleStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorRed)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	snake := &Snake{
		ColStart: 10,
		ColEnd:   25,
		RowStart: 5,
		RowEnd:   6,
	}

	for row := 5; row <= 40; row++ {
		for col := 10; col <= 200; col++ {
			s.SetContent(col, row, ' ', nil, boxStyle)
		}
	}

	renderSnake(s, snakeStyle, snake)

	for row := 10; row <= 11; row++ {
		for col := 16; col <= 20; col++ {
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
				moveSnakeUp(s, snakeStyle, snake)
			}
			if ev.Key() == tcell.KeyDown {
				moveSnakeDown(s, snakeStyle, snake)
			}
			if ev.Key() == tcell.KeyLeft {
				moveSnakeLeft(s, snakeStyle, snake)
			}
			if ev.Key() == tcell.KeyRight {
				moveSnakeRight(s, snakeStyle, snake)
			}

		}
	}

}
