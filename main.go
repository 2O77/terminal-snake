package main

import (
	"log"
	"os"
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

type Point struct {
	X int
	Y int
}

type Snake struct {
	Body []Point
}

type ViewBoardCoordinates struct {
	boardLeftX   int
	boardRightX  int
	boardTopY    int
	boardBottomY int
}

const BoardModelColumns = 40
const BoardModelRows = 20

const (
	boardWidthViewCells  = BoardModelColumns * 2
	boardHeightViewCells = BoardModelRows
)

func SetScene(screen tcell.Screen, snake *Snake) {
	screen.Clear()
	setContent(snake, screen)
	// setTopbar(screen)
	screen.Show()

}

func setContent(snake *Snake, screen tcell.Screen) {
	unoccupiedCellStyle := tcell.StyleDefault.Foreground(tcell.ColorBlack).Background(tcell.ColorBlack)

	backgroundStyle := tcell.StyleDefault.
		Background(tcell.ColorDarkGreen).
		Foreground(tcell.ColorWhite)

	const runeValue rune = 0
	var combining []rune = nil

	screen.SetStyle(unoccupiedCellStyle)

	viewBoardCoordinates, ok := viewBoardCoordinates(screen)
	if !ok {
		log.Fatal()
	}

	for viewColumn := 0; viewColumn < boardWidthViewCells; viewColumn += 2 {
		for viewRow := 0; viewRow < boardHeightViewCells; viewRow += 1 {
			for index := range 2 {
				screen.SetContent(viewBoardCoordinates.boardLeftX+viewColumn+index, viewBoardCoordinates.boardTopY+viewRow, runeValue, combining, backgroundStyle)
			}
		}
	}

	setSnake(screen, snake, viewBoardCoordinates)

}

func setSnake(screen tcell.Screen, snake *Snake, board ViewBoardCoordinates) {
	const runeValue rune = 0
	var combining []rune = nil

	snakeStyle := tcell.StyleDefault.
		Background(tcell.ColorRed).
		Foreground(tcell.Color182)

	headStyle := tcell.StyleDefault.
		Background(tcell.ColorHotPink).
		Foreground(tcell.Color182)

	head := snake.Body[len(snake.Body)-1]

	log.Println(head.X, head.Y, board.boardLeftX, board.boardRightX)

	if head.X < 0 || head.X > boardWidthViewCells || head.Y < 0 || head.Y > boardHeightViewCells {
		return
	}

	for index, value := range snake.Body {
		if index == len(snake.Body)-1 {
			for index := range 2 {
				screen.SetContent(board.boardLeftX+value.X+index-1, board.boardTopY+value.Y, runeValue, combining, headStyle)
			}
		} else {
			for index := range 2 {
				screen.SetContent(board.boardLeftX+value.X+index-1, board.boardTopY+value.Y, runeValue, combining, snakeStyle)
			}
		}
	}
}

// func setTopbar(screen tcell.Screen) {
// 	backgroundStyle := tcell.StyleDefault.
// 		Background(tcell.ColorBlack).
// 		Foreground(tcell.ColorWhite)

// 	const runeValue rune = 0
// 	var combining []rune = nil

// 	screen.SetContent(170, 0, runeValue, combining, backgroundStyle)
// }

func MoveSnakeRight(snake *Snake, head Point) {
	snake.Body = snake.Body[1:]

	newHead := Point{
		X: head.X + 2,
		Y: head.Y,
	}

	snake.Body = append(snake.Body, newHead)
}

func MoveSnakeLeft(snake *Snake, head Point) {
	snake.Body = snake.Body[1:]

	newHead := Point{
		X: head.X - 2,
		Y: head.Y,
	}

	snake.Body = append(snake.Body, newHead)
}

func MoveSnakeUp(snake *Snake, head Point) {
	snake.Body = snake.Body[1:]

	newHead := Point{
		X: head.X,
		Y: head.Y - 1,
	}

	snake.Body = append(snake.Body, newHead)
}

func MoveSnakeDown(snake *Snake, head Point) {
	snake.Body = snake.Body[1:]

	newHead := Point{
		X: head.X,
		Y: head.Y + 1,
	}

	snake.Body = append(snake.Body, newHead)
}

func viewBoardCoordinates(screen tcell.Screen) (result ViewBoardCoordinates, ok bool) {
	w, h := screen.Size()

	if w < boardWidthViewCells || h < boardHeightViewCells {
		ok = false
		return
	}

	ok = true

	result.boardLeftX = (w - boardWidthViewCells) / 2

	result.boardRightX = result.boardLeftX + boardWidthViewCells - 2

	result.boardTopY = (h - boardHeightViewCells) / 2

	result.boardBottomY = result.boardTopY + boardHeightViewCells - 1

	return
}

func main() {
	logFile := initLogger()
	defer logFile.Close()

	var snake *Snake = &Snake{
		Body: []Point{
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

	defer quit(screen)

	tickerX := time.NewTicker(180 * time.Millisecond)
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

			SetScene(screen, snake)
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
