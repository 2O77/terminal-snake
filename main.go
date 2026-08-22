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

func setContent(snake *Snake, screen tcell.Screen) {
	backgroundStyle := tcell.StyleDefault.
		Background(tcell.ColorDarkGreen).
		Foreground(tcell.ColorWhite)

	snakeStyle := tcell.StyleDefault.
		Background(tcell.ColorGrey).
		Foreground(tcell.Color182)

	headStyle := tcell.StyleDefault.
		Background(tcell.ColorHotPink).
		Foreground(tcell.Color182)

	const runeValue rune = 0
	var combining []rune = nil

	screen.SetStyle(backgroundStyle)

	for index, value := range snake.Body {
		if index == len(snake.Body)-1 {
			screen.SetContent(value.X, value.Y, runeValue, combining, headStyle)
		} else {
			screen.SetContent(value.X, value.Y, runeValue, combining, snakeStyle)
		}
	}

}

func setTopbar(screen tcell.Screen) {
	backgroundStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)

	const runeValue rune = 0
	var combining []rune = nil

	screen.SetContent(170, 0, runeValue, combining, backgroundStyle)

}
func SetSnake(screen tcell.Screen, snake *Snake, snakeDirection *LastDirection, key tcell.Key) {
	head := snake.Body[len(snake.Body)-1]
	neck := snake.Body[len(snake.Body)-2]

	if key == tcell.KeyRight {
		if head.Y == neck.Y {
			if head.X < neck.X {
				return
			}
		}
		snake.Body = snake.Body[1:]

		head := snake.Body[len(snake.Body)-1]
		newHead := Point{
			X: head.X + 1,
			Y: head.Y,
		}

		snake.Body = append(snake.Body, newHead)
		*snakeDirection = DirectionRight

	}
	if key == tcell.KeyLeft {
		if head.Y == neck.Y {
			if head.X > neck.X {
				return
			}
		}

		snake.Body = snake.Body[1:]
		newHead := Point{
			X: head.X - 1,
			Y: head.Y,
		}

		snake.Body = append(snake.Body, newHead)
		*snakeDirection = DirectionLeft
	}
	if key == tcell.KeyUp {
		if head.X == neck.X {
			if head.Y > neck.Y {
				return
			}
		}
		snake.Body = snake.Body[1:]

		head := snake.Body[len(snake.Body)-1]
		newHead := Point{
			X: head.X,
			Y: head.Y - 1,
		}

		snake.Body = append(snake.Body, newHead)
		*snakeDirection = DirectionUp
	}
	if key == tcell.KeyDown {
		if head.X == neck.X {
			if head.Y < neck.Y {
				return
			}
		}
		snake.Body = snake.Body[1:]

		head := snake.Body[len(snake.Body)-1]
		newHead := Point{
			X: head.X,
			Y: head.Y + 1,
		}

		snake.Body = append(snake.Body, newHead)
		*snakeDirection = DirectionDown
	}
}

func main() {
	var snake *Snake = &Snake{
		Body: []Point{
			{5, 10},
			{6, 10},
			{7, 10},
			{8, 10},
			{9, 10},
			{10, 10},
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

	screen.SetSize(50, 10)
	defer quit(screen)

	for {
		screen.Clear()
		setContent(snake, screen)
		setTopbar(screen)
		screen.Show()

		timeChannel := time.Tick(1 * time.Second)
		for time := range timeChannel {
			log.Print("timer", time)
		}

		switch ev := screen.PollEvent().(type) {
		case *tcell.EventResize:
			screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				screen.Fini()
				os.Exit(0)
			} else {
				SetSnake(screen, snake, lastDirection, ev.Key())
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
