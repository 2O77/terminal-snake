package model

type Snake struct {
	apple          *Apple
	Body           []Box
	LastDirections *LastDirections
	Direction      *SnakeDirection
	gameOver       bool
}

type SnakeDeps struct {
	Apple *Apple
}

type SnakeDirection string

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
	snake.Direction = &initialDir
	return snake
}

func (s *Snake) IsGameOver() bool {
	return s.gameOver
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

	if OneToMany(newHead, s.Body) || IsOutside(newHead, BoardBoxColumns, BoardBoxRows) {
		s.gameOver = true
	}

	s.Body = append(s.Body, newHead)
}

// EnqueueDirections queues a pending move. A direction that would reverse the
// snake (relative to the last queued or executed direction) is ignored.
func (s *Snake) EnqueueDirections(direction SnakeDirection) {
	if s.isOpposite(direction) {
		return
	}

	directions := s.LastDirections

	for index, value := range directions {
		if value == "" {
			directions[index] = direction
			return
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

func (s *Snake) isOpposite(direction SnakeDirection) bool {
	last := s.lastQueuedDirection()
	if last == "" {
		last = *s.Direction
	}

	switch last {
	case SnakeDirectionUp:
		return direction == SnakeDirectionDown
	case SnakeDirectionDown:
		return direction == SnakeDirectionUp
	case SnakeDirectionLeft:
		return direction == SnakeDirectionRight
	case SnakeDirectionRight:
		return direction == SnakeDirectionLeft
	}

	return false
}

func (s *Snake) lastQueuedDirection() SnakeDirection {
	for index := len(s.LastDirections) - 1; index >= 0; index-- {
		if s.LastDirections[index] != "" {
			return s.LastDirections[index]
		}
	}

	return ""
}

func (s *Snake) isHeadOn(box Box) bool {
	head := s.head()

	return box.X == head.X && box.Y == head.Y
}

func (s *Snake) head() Box {
	head := s.Body[len(s.Body)-1]
	return head
}

func (s *Snake) Reset() {
	s.Body = nil

	for index := range InitialSnakeSize {
		value := Box{
			X: index,
			Y: 0,
		}

		s.Body = append(s.Body, value)
	}

	initialDir := SnakeDirectionRight
	s.Direction = &initialDir
	s.LastDirections = &LastDirections{}
	s.gameOver = false
}
