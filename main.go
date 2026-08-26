package main

import (
	"log"
	"os"
	model "terminal-snake/models"
	"terminal-snake/timer"
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

	timer := timer.NewTimer(timer.TimerDeps{snake})
	timer.SetTimer()
	timer.SetPeriod()

	for {
		if screen.HasPendingEvent() {
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
						if *snake.Direction != model.SnakeDirectionDown && *snake.Direction != model.SnakeDirectionUp {
							snake.EnqueueDirections(model.SnakeDirectionUp)

						}
					}
					if ev.Key() == tcell.KeyDown {
						if *snake.Direction != model.SnakeDirectionDown && *snake.Direction != model.SnakeDirectionUp {
							snake.EnqueueDirections(model.SnakeDirectionDown)
						}
					}
					if ev.Key() == tcell.KeyRight {
						if *snake.Direction != model.SnakeDirectionLeft && *snake.Direction != model.SnakeDirectionRight {
							snake.EnqueueDirections(model.SnakeDirectionRight)
						}
					}
					if ev.Key() == tcell.KeyLeft {
						if *snake.Direction != model.SnakeDirectionLeft && *snake.Direction != model.SnakeDirectionRight {
							snake.EnqueueDirections(model.SnakeDirectionLeft)
						}
					}
				}
			}

			timer.SetPeriod()
		}

		if time.Since(timer.Now) > timer.Period {
			if !isGameOver {
				switch snake.LastDirections[0] {
				case model.SnakeDirectionRight:
					snake.MoveSnakeRight()
				case model.SnakeDirectionLeft:
					snake.MoveSnakeLeft()
				case model.SnakeDirectionUp:
					snake.MoveSnakeUp()
				case model.SnakeDirectionDown:
					snake.MoveSnakeDown()
				default:
					switch *snake.Direction {
					case model.SnakeDirectionRight:
						snake.MoveSnakeRight()
					case model.SnakeDirectionLeft:
						snake.MoveSnakeLeft()
					case model.SnakeDirectionUp:
						snake.MoveSnakeUp()
					case model.SnakeDirectionDown:
						snake.MoveSnakeDown()
					}
				}
			}

			snake.DequeueDirections()
			timer.SetTimer()
			timer.SetPeriod()

		}

		view.SetView()
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
		log.Fatalf("Log file cannot open: %v", err)
	}

	log.SetOutput(file)
	return file
}
