package model

import (
	"log"
)

type Snake struct {
	apple          *Apple
	Body           []Box
	LastDirections *LastDirections
	Direction      *SnakeDirection
	IsGameOver     *bool
}

type SnakeDeps struct {
	IsGameOver *bool
	Apple      *Apple
}

type SnakeDirection string

type PermanentSnakeDirection SnakeDirection
type LastDirections [10]SnakeDirection

const (
	SnakeDirectionUp    SnakeDirection = "UP"
	SnakeDirectionDown  SnakeDirection = "DOWN"
	SnakeDirectionLeft  SnakeDirection = "LEFT"
	SnakeDirectionRight SnakeDirection = "RIGHT"
)

const InitialSnakeSize = 10

func NewSnake(deps SnakeDeps) *Snake {
	snake := &Snake{}

	for index := range InitialSnakeSize {
		value := Box{
			X: index,
			Y: 0,
		}

		snake.Body = append(snake.Body, value)
	}

	initialDir := SnakeDirectionRight
	snake.LastDirections = &LastDirections{}

	snake.apple = deps.Apple
	snake.IsGameOver = deps.IsGameOver
	snake.Direction = &initialDir
	return snake
}

func (s *Snake) MoveSnakeRight() {
	head := s.head()

	newHead := Box{
		X: head.X + 1,
		Y: head.Y,
	}

	*s.Direction = SnakeDirectionRight
	s.move(newHead)
}

func (s *Snake) MoveSnakeLeft() {
	head := s.head()

	newHead := Box{
		X: head.X - 1,
		Y: head.Y,
	}

	*s.Direction = SnakeDirectionLeft
	s.move(newHead)
}

func (s *Snake) MoveSnakeUp() {
	head := s.head()

	newHead := Box{
		X: head.X,
		Y: head.Y - 1,
	}

	*s.Direction = SnakeDirectionUp
	s.move(newHead)
}

func (s *Snake) MoveSnakeDown() {
	head := s.head()

	newHead := Box{
		X: head.X,
		Y: head.Y + 1,
	}

	*s.Direction = SnakeDirectionDown
	s.move(newHead)
}

func (s *Snake) move(newHead Box) {
	if s.isHeadOn(Box(*s.apple)) {
		s.apple.SummonApple(s.Body)
	} else {
		s.Body = s.Body[1:]
	}

	if OneToMany(newHead, s.Body) {
		log.Print("Game Over")
		*s.IsGameOver = true
	}

	if IsOutside(newHead, BoardBoxColumns, BoardBoxRows) {
		log.Print("Game Over")
		*s.IsGameOver = true
	}

	s.Body = append(s.Body, newHead)
}

func (s *Snake) EnqueueDirections(lastMove SnakeDirection) {
	directions := s.LastDirections

	for index, value := range directions {
		if value == "" {
			directions[index] = lastMove
			break
		}
	}
}

func (s *Snake) DequeueDirections() {
	snake := s

	for index := 0; index < len(snake.LastDirections)-1; index++ {
		snake.LastDirections[index] = snake.LastDirections[index+1]
	}

	snake.LastDirections[len(snake.LastDirections)-1] = ""
}

func (s *Snake) DirectionCount() int {
	moves := 0

	for _, value := range s.LastDirections {
		if value != "" {
			moves++
		}
	}
	if moves == 0 {
		moves = 1
	}

	return moves
}

func (s *Snake) isHeadOn(box Box) bool {
	head := s.head()

	if box.X == head.X && box.Y == head.Y {
		return true
	}

	return false
}

func (s *Snake) head() Box {
	head := s.Body[len(s.Body)-1]
	return head
}

func (s *Snake) tail() Box {
	tail := s.Body[0]
	return tail
}
