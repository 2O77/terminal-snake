package main

import (
	"log"
	"os"
	"terminal-snake/models"
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

type ViewBoardCoordinates struct {
	boardLeftX   int
	boardRightX  int
	boardTopY    int
	boardBottomY int
}

const BoardBoxColumns = 40
const BoardBoxRows = 20

const (
	boardWidthViewCells  = BoardBoxColumns * 2
	boardHeightViewCells = BoardBoxRows
)

// func setTopbar(screen tcell.Screen) {
// 	backgroundStyle := tcell.StyleDefault.
// 		Background(tcell.ColorBlack).
// 		Foreground(tcell.ColorWhite)

// 	const runeValue rune = 0
// 	var combining []rune = nil

// 	screen.SetContent(170, 0, runeValue, combining, backgroundStyle)
// }

func MoveSnakeRight(snake *models.Snake, head models.Box) {
	snake.Body = snake.Body[1:]

	newHead := models.Box{
		X: head.X + 1,
		Y: head.Y,
	}

	snake.Body = append(snake.Body, newHead)
}

func MoveSnakeLeft(snake *models.Snake, head models.Box) {
	snake.Body = snake.Body[1:]

	newHead := models.Box{
		X: head.X - 1,
		Y: head.Y,
	}

	snake.Body = append(snake.Body, newHead)
}

func MoveSnakeUp(snake *models.Snake, head models.Box) {
	snake.Body = snake.Body[1:]

	newHead := models.Box{
		X: head.X,
		Y: head.Y - 1,
	}

	snake.Body = append(snake.Body, newHead)
}

func MoveSnakeDown(snake *models.Snake, head models.Box) {
	snake.Body = snake.Body[1:]

	newHead := models.Box{
		X: head.X,
		Y: head.Y + 1,
	}

	snake.Body = append(snake.Body, newHead)
}

func main() {
	logFile := initLogger()
	defer logFile.Close()

	var snake *models.Snake = &models.Snake{
		Body: []models.Box{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
			{4, 0},
			{5, 0},
		},
	}

	initialDir := DirectionRight
	var lastDirection *LastDirection = &initialDir

	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := screen.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	view := board_view.NewView(
		board_view.ViewDeps{
			snake, screen,
		},
	)

	defer quit(screen)

	tickerX := time.NewTicker(500 * time.Millisecond)
	go func() {
		for range tickerX.C {
			switch *lastDirection {
			case DirectionRight:
				MoveSnakeRight(snake, snake.Body[len(snake.Body)-1])
			case DirectionLeft:
				MoveSnakeLeft(snake, snake.Body[len(snake.Body)-1])
			case DirectionUp:
				MoveSnakeUp(snake, snake.Body[len(snake.Body)-1])
			case DirectionDown:
				MoveSnakeDown(snake, snake.Body[len(snake.Body)-1])
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
				if *lastDirection != DirectionDown {
					*lastDirection = DirectionUp
				}
			}
			if ev.Key() == tcell.KeyDown {
				if *lastDirection != DirectionUp {
					*lastDirection = DirectionDown
				}
			}
			if ev.Key() == tcell.KeyLeft {
				if *lastDirection != DirectionRight {
					*lastDirection = DirectionLeft
				}
			}
			if ev.Key() == tcell.KeyRight {
				if *lastDirection != DirectionLeft {
					*lastDirection = DirectionRight
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
