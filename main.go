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

func SetSnakeRight(snake *Snake, head Point) {
	snake.Body = snake.Body[1:]

	newHead := Point{
		X: head.X + 1,
		Y: head.Y,
	}

	snake.Body = append(snake.Body, newHead)
}

func SetSnakeLeft(snake *Snake, head Point) {
	snake.Body = snake.Body[1:]

	newHead := Point{
		X: head.X - 1,
		Y: head.Y,
	}

	snake.Body = append(snake.Body, newHead)
}

func SetSnakeUp(snake *Snake, head Point) {
	snake.Body = snake.Body[1:]

	newHead := Point{
		X: head.X,
		Y: head.Y - 1,
	}

	snake.Body = append(snake.Body, newHead)
}

func SetSnakeDown(snake *Snake, head Point) {
	snake.Body = snake.Body[1:]

	newHead := Point{
		X: head.X,
		Y: head.Y + 1,
	}

	snake.Body = append(snake.Body, newHead)
}

func SetSnakeOnKeyEvent(screen tcell.Screen, snake *Snake, snakeDirection *LastDirection, key tcell.Key) {
	head := snake.Body[len(snake.Body)-1]
	neck := snake.Body[len(snake.Body)-2]

	if key == tcell.KeyRight {
		if head.Y == neck.Y {
			if head.X < neck.X {
				return
			}
		}

		SetSnakeRight(snake, head)
		*snakeDirection = DirectionRight
	}
	if key == tcell.KeyLeft {
		if head.Y == neck.Y {
			if head.X > neck.X {
				return
			}
		}

		SetSnakeLeft(snake, head)
		*snakeDirection = DirectionLeft
	}
	if key == tcell.KeyUp {
		if head.X == neck.X {
			if head.Y > neck.Y {
				return
			}
		}

		SetSnakeUp(snake, head)
		*snakeDirection = DirectionUp
	}
	if key == tcell.KeyDown {
		if head.X == neck.X {
			if head.Y < neck.Y {
				return
			}
		}

		SetSnakeDown(snake, head)
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

	go func() {
		for {
			screen.Clear()
			setContent(snake, screen)
			setTopbar(screen)
			screen.Show()

			switch ev := screen.PollEvent().(type) {
			case *tcell.EventResize:
				screen.Sync()
			case *tcell.EventKey:
				if ev.Key() == tcell.KeyEscape {
					screen.Fini()
					os.Exit(0)
				} else {
					SetSnakeOnKeyEvent(screen, snake, lastDirection, ev.Key())
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-time.After(300 * time.Millisecond):
				switch *lastDirection {
				case DirectionRight:
					SetSnakeRight(snake, snake.Body[len(snake.Body)-1])
				case DirectionLeft:
					SetSnakeLeft(snake, snake.Body[len(snake.Body)-1])
				case DirectionUp:
					SetSnakeUp(snake, snake.Body[len(snake.Body)-1])
				case DirectionDown:
					SetSnakeDown(snake, snake.Body[len(snake.Body)-1])
				}

				screen.Clear()
				setContent(snake, screen)
				setTopbar(screen)
				screen.Show()
			}
		}
	}()

	for {
	}

}

func quit(s tcell.Screen) {
	maybePanic := recover()
	s.Fini()
	if maybePanic != nil {
		panic(maybePanic)
	}
}
