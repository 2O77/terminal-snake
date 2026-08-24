package model

type Snake struct {
	Apple *Apple
	Body  []Box
}

type SnakeDeps struct {
	Apple *Apple
}

func NewSnake(deps SnakeDeps) *Snake {
	var snake *Snake = &Snake{
		Body: []Box{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
			{4, 0},
			{5, 0},
		},
	}

	snake.Apple = deps.Apple
	return snake
}

func (s *Snake) MoveSnakeRight() {
	head := s.Head()

	newHead := Box{
		X: head.X + 1,
		Y: head.Y,
	}

	if s.IsHeadOn(Box(*s.Apple)) {
		s.Apple.SummonApple(s.Body)
	} else {
		s.Body = s.Body[1:]
	}

	s.Body = append(s.Body, newHead)
}

func (s *Snake) MoveSnakeLeft() {
	head := s.Head()

	newHead := Box{
		X: head.X - 1,
		Y: head.Y,
	}

	if s.IsHeadOn(Box(*s.Apple)) {
		s.Apple.SummonApple(s.Body)
	} else {
		s.Body = s.Body[1:]
	}

	s.Body = append(s.Body, newHead)
}

func (s *Snake) MoveSnakeUp() {
	head := s.Head()

	newHead := Box{
		X: head.X,
		Y: head.Y - 1,
	}

	if s.IsHeadOn(Box(*s.Apple)) {
		s.Apple.SummonApple(s.Body)
	} else {
		s.Body = s.Body[1:]
	}

	s.Body = append(s.Body, newHead)
}

func (s *Snake) MoveSnakeDown() {
	head := s.Head()

	newHead := Box{
		X: head.X,
		Y: head.Y + 1,
	}

	if s.IsHeadOn(Box(*s.Apple)) {
		s.Apple.SummonApple(s.Body)
	} else {
		s.Body = s.Body[1:]
	}

	s.Body = append(s.Body, newHead)
}

func (s *Snake) IsHeadOn(box Box) bool {
	head := s.Head()

	if box.X == head.X && box.Y == head.Y {
		return true
	}

	return false
}

func (s *Snake) Head() Box {
	head := s.Body[len(s.Body)-1]
	return head
}

func (s *Snake) Tail() Box {
	tail := s.Body[0]
	return tail
}
