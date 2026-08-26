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

	permanentLastDirection := snake.PermanentSnakeDirection

	// go func() {
	// 	var moves time.Duration
	// 	moves = 0
	// 	for _, value := range snake.LastDirections {
	// 		if value != "" {
	// 			moves++
	// 		}
	// 	}
	// 	if moves == 0 {
	// 		moves = 1
	// 	}

	// 	ticker := timer.NewTime(len(snake.Body))
	// 	period := ticker / moves * time.Millisecond

	// 	timer := time.NewTicker(period)

	// 	for range timer.C {
	// 		if !isGameOver {
	// 			switch *permanentLastDirection {
	// 			case model.SnakeDirectionRight:
	// 				snake.MoveSnakeRight()
	// 			case model.SnakeDirectionLeft:
	// 				snake.MoveSnakeLeft()
	// 			case model.SnakeDirectionUp:
	// 				snake.MoveSnakeUp()
	// 			case model.SnakeDirectionDown:
	// 				snake.MoveSnakeDown()
	// 			default:
	// 				log.Fatal("permanent last directioin cannot be empty")
	// 			}

	// 			cleanDirectionQueue(*snake)

	// 			period = ticker / moves * time.Millisecond

	// 			timer.Reset(period)

	// 			view.SetView()
	// 		}
	// 	}
	// }()

	var moves int
	moves = 0
	for _, value := range snake.LastDirections {
		if value != "" {
			moves++
		}
	}
	if moves == 0 {
		moves = 1
	}

	lapTime := timer.NewTime(len(snake.Body))
	period := lapTime / time.Duration(moves) * time.Millisecond

	now := time.Now()

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
						if *permanentLastDirection != model.SnakeDirectionDown && *permanentLastDirection != model.SnakeDirectionUp {
							setDirection(model.SnakeDirectionUp, snake.LastDirections, (*model.PermanentSnakeDirection)(permanentLastDirection))
						}
					}
					if ev.Key() == tcell.KeyDown {
						if *permanentLastDirection != model.SnakeDirectionDown && *permanentLastDirection != model.SnakeDirectionUp {
							setDirection(model.SnakeDirectionDown, snake.LastDirections, (*model.PermanentSnakeDirection)(permanentLastDirection))
						}
					}
					if ev.Key() == tcell.KeyRight {
						if *permanentLastDirection != model.SnakeDirectionLeft && *permanentLastDirection != model.SnakeDirectionRight {
							setDirection(model.SnakeDirectionRight, snake.LastDirections, (*model.PermanentSnakeDirection)(permanentLastDirection))
						}
					}
					if ev.Key() == tcell.KeyLeft {
						if *permanentLastDirection != model.SnakeDirectionLeft && *permanentLastDirection != model.SnakeDirectionRight {
							setDirection(model.SnakeDirectionLeft, snake.LastDirections, (*model.PermanentSnakeDirection)(permanentLastDirection))
						}
					}
				}
			}

		}

		if time.Since(now) > period {
			switch *permanentLastDirection {
			case model.SnakeDirectionRight:
				snake.MoveSnakeRight()
			case model.SnakeDirectionLeft:
				snake.MoveSnakeLeft()
			case model.SnakeDirectionUp:
				snake.MoveSnakeUp()
			case model.SnakeDirectionDown:
				snake.MoveSnakeDown()
			default:
				log.Fatal("permanent last directioin cannot be empty")
			}

			cleanDirectionQueue(*snake)

			period = lapTime / time.Duration(moves) * time.Millisecond
			lapTime = timer.NewTime(len(snake.Body))
			now = time.Now()

			for _, value := range snake.LastDirections {
				if value != "" {
					moves++
				}
			}
			if moves == 0 {
				moves = 1
			}

			log.Print(moves)

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

func setDirection(lastMove model.SnakeDirection, directions *model.LastDirections, permanentDirection *model.PermanentSnakeDirection) {
	isEmpty := false
	for index, value := range directions {
		if value == "" {
			directions[index] = lastMove
			isEmpty = true
			break
		}
	}

	if isEmpty {
		*permanentDirection = model.PermanentSnakeDirection(lastMove)
	}
}

func cleanDirectionQueue(snake model.Snake) {
	for index := 0; index < len(snake.LastDirections)-1; index++ {
		snake.LastDirections[index] = snake.LastDirections[index+1]
	}

	snake.LastDirections[len(snake.LastDirections)-1] = ""
}
