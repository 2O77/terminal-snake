package main

import (
	"log"
	"os"
	model "terminal-snake/models"
	board_view "terminal-snake/views"
	"time"

	"github.com/gdamore/tcell/v2"
)

type LastDirection string

const (
	DirectionUp    LastDirection = "UP"
	DirectionDown  LastDirection = "DOWN"
	DirectionLeft  LastDirection = "LEFT"
	DirectionRight LastDirection = "RIGHT"
)

// func setTopbar(screen tcell.Screen) {
// 	backgroundStyle := tcell.StyleDefault.
// 		Background(tcell.ColorBlack).
// 		Foreground(tcell.ColorWhite)

// 	const runeValue rune = 0
// 	var combining []rune = nil

// 	screen.SetContent(170, 0, runeValue, combining, backgroundStyle)
// }

func main() {
	logFile := initLogger()
	defer logFile.Close()

	initialDir := DirectionRight
	var snakeLastDirection *LastDirection = &initialDir

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	apple := model.NewApple()
	snake := model.NewSnake(model.SnakeDeps{
		apple,
	})

	view := board_view.NewView(
		board_view.ViewDeps{
			apple,
			snake,
			screen,
			model.BoardBoxColumns,
			model.BoardBoxRows,
		},
	)

	defer quit(screen)

	tickerX := time.NewTicker(200 * time.Millisecond)
	go func() {
		for range tickerX.C {
			switch *snakeLastDirection {
			case DirectionRight:
				snake.MoveSnakeRight()
			case DirectionLeft:
				snake.MoveSnakeLeft()
			case DirectionUp:
				snake.MoveSnakeUp()
			case DirectionDown:
				snake.MoveSnakeDown()
			}

			view.SetView()
		}
	}()

	for {
		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				screen.Fini()
				os.Exit(0)
			}
			if ev.Key() == tcell.KeyUp {
				if *snakeLastDirection != DirectionDown {
					*snakeLastDirection = DirectionUp
				}
			}
			if ev.Key() == tcell.KeyDown {
				if *snakeLastDirection != DirectionUp {
					*snakeLastDirection = DirectionDown
				}
			}
			if ev.Key() == tcell.KeyLeft {
				if *snakeLastDirection != DirectionRight {
					*snakeLastDirection = DirectionLeft
				}
			}
			if ev.Key() == tcell.KeyRight {
				if *snakeLastDirection != DirectionLeft {
					*snakeLastDirection = DirectionRight
				}
			}
		}
	}
}

func quit(s tcell.Screen) {
	maybePanic := recover()
	s.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}

func initLogger() *os.File {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Log dosyası açılamadı: %v", err)
	}

	log.SetOutput(file)
	return file
}
