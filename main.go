package main

import (
	"log"
	"os"
	model "terminal-snake/models"
	board_view "terminal-snake/views"
	"time"

	"github.com/gdamore/tcell/v2"
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

	isGameOver := false

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	apple := model.NewApple()
	snake := model.NewSnake(model.SnakeDeps{
		&isGameOver, apple,
	})

	view := board_view.NewView(
		board_view.ViewDeps{
			&isGameOver,
			apple,
			snake,
			screen,
			model.BoardBoxColumns,
			model.BoardBoxRows,
		},
	)

	defer quit(screen)

	snakeLastDirection := snake.LastDirection

	tickerX := time.NewTicker(100 * time.Millisecond)
	go func() {
		for range tickerX.C {
			if !isGameOver {
				switch *snakeLastDirection {
				case model.SnakeDirectionRight:
					snake.MoveSnakeRight()
				case model.SnakeDirectionLeft:
					snake.MoveSnakeLeft()
				case model.SnakeDirectionUp:
					snake.MoveSnakeUp()
				case model.SnakeDirectionDown:
					snake.MoveSnakeDown()
				}

				view.SetView()
			}
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
			if !isGameOver {
				if ev.Key() == tcell.KeyUp {
					if *snakeLastDirection != model.SnakeDirectionDown {
						*snakeLastDirection = model.SnakeDirectionUp
					}
				}
				if ev.Key() == tcell.KeyDown {
					if *snakeLastDirection != model.SnakeDirectionUp {
						*snakeLastDirection = model.SnakeDirectionDown
					}
				}
				if ev.Key() == tcell.KeyLeft {
					if *snakeLastDirection != model.SnakeDirectionRight {
						*snakeLastDirection = model.SnakeDirectionLeft
					}
				}
				if ev.Key() == tcell.KeyRight {
					if *snakeLastDirection != model.SnakeDirectionLeft {
						*snakeLastDirection = model.SnakeDirectionRight
					}
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
