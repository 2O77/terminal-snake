package main

import (
	"log"
	"os"
	"terminal-snake/config"
	model "terminal-snake/models"
	"terminal-snake/timer"
	board_view "terminal-snake/views"
	"time"

	"github.com/gdamore/tcell/v2"
)

func main() {
	logFile := initLogger()
	defer logFile.Close()

	cfg, err := config.Load()
	if err != nil {
		log.Print("couldn't load config: ", err)
	}

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	apple := model.NewApple()
	snake := model.NewSnake(model.SnakeDeps{Apple: apple})
	bestScore := cfg.BestScore

	view := board_view.NewView(board_view.ViewDeps{
		BestScore:       bestScore,
		Apple:           apple,
		Snake:           snake,
		Screen:          screen,
		BoardBoxColumns: model.BoardBoxColumns,
		BoardBoxRows:    model.BoardBoxRows,
	})

	gameTimer := timer.NewTimer(timer.TimerDeps{Snake: snake})
	gameTimer.SetTimer()
	gameTimer.SetPeriod()

	savedBestScore := false
	paused := false
	countdown := 0
	countdownDeadline := time.Time{}

	defer quit(screen)

	// draw the initial frame
	view.SetView()

	for {
		needsRedraw := false

		if screen.HasPendingEvent() {
			switch ev := screen.PollEvent().(type) {
			case *tcell.EventResize:
				screen.Sync()
				needsRedraw = true
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					screen.Fini()
					os.Exit(0)
				}
				if ev.Key() == tcell.KeyRune {
					switch ev.Rune() {
					case 'r':
						reset(apple, snake, &gameTimer)
						savedBestScore = false
						paused = false
						countdown = 0
						view.SetPaused(false)
						view.SetCountdown(0)
						needsRedraw = true
					case ' ', 'p':
						if !snake.IsGameOver() {
							if countdown > 0 {
								// cancel countdown, stay paused
								countdown = 0
								view.SetCountdown(0)
								needsRedraw = true
							} else if paused {
								// start countdown to resume
								countdown = 3
								countdownDeadline = time.Now().Add(time.Second)
								view.SetCountdown(countdown)
								needsRedraw = true
							} else {
								paused = true
								view.SetPaused(true)
								needsRedraw = true
							}
						}
					}
				}
				if !snake.IsGameOver() && !paused {
					switch ev.Key() {
					case tcell.KeyUp:
						snake.EnqueueDirections(model.SnakeDirectionUp)
					case tcell.KeyDown:
						snake.EnqueueDirections(model.SnakeDirectionDown)
					case tcell.KeyRight:
						snake.EnqueueDirections(model.SnakeDirectionRight)
					case tcell.KeyLeft:
						snake.EnqueueDirections(model.SnakeDirectionLeft)
					}
				}
			}

			gameTimer.SetPeriod()
		}

		if countdown > 0 && time.Now().After(countdownDeadline) {
			countdown--
			view.SetCountdown(countdown)
			needsRedraw = true

			if countdown == 0 {
				// resume the game
				paused = false
				view.SetPaused(false)
				gameTimer.SetTimer()
			} else {
				countdownDeadline = time.Now().Add(time.Second)
			}
		}

		if !paused && time.Since(gameTimer.Now) > gameTimer.Period {
			if !snake.IsGameOver() {
				direction := snake.LastDirections[0]
				if direction == "" {
					direction = *snake.Direction
				}

				switch direction {
				case model.SnakeDirectionRight:
					snake.MoveSnakeRight()
				case model.SnakeDirectionLeft:
					snake.MoveSnakeLeft()
				case model.SnakeDirectionUp:
					snake.MoveSnakeUp()
				case model.SnakeDirectionDown:
					snake.MoveSnakeDown()
				}

				needsRedraw = true
			}

			snake.DequeueDirections()
			gameTimer.SetTimer()
			gameTimer.SetPeriod()
		}

		if snake.IsGameOver() && !savedBestScore {
			savedBestScore = true
			if score := len(snake.Body) - model.InitialSnakeSize; score > bestScore {
				bestScore = score
				view.SetBestScore(bestScore)
				cfg.BestScore = bestScore
				if err := cfg.Save(); err != nil {
					log.Print("couldn't save best score: ", err)
				}

				needsRedraw = true
			}
		}

		if needsRedraw {
			view.SetView()
		}

		time.Sleep(1 * time.Millisecond)
	}
}

func quit(s tcell.Screen) {
	maybePanic := recover()
	s.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}

func reset(apple *model.Apple, snake *model.Snake, gameTimer *timer.Timer) {
	snake.Reset()
	apple.Reset()
	gameTimer.Reset(snake)
}

func initLogger() *os.File {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Log file cannot open: %v", err)
	}

	log.SetOutput(file)
	return file
}
