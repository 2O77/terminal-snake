package model

type Box struct {
	X int
	Y int
}

type Snake struct {
	Body []Box
}

func NewSnake() *Snake {
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

	return snake
}

func (s *Snake) MoveSnakeRight() {
	s.Body = s.Body[1:]

	head := s.GiveMeHead()

	newHead := Box{
		X: head.X + 1,
		Y: head.Y,
	}

	s.Body = append(s.Body, newHead)
}

func (s *Snake) MoveSnakeLeft() {
	s.Body = s.Body[1:]

	head := s.GiveMeHead()

	newHead := Box{
		X: head.X - 1,
		Y: head.Y,
	}

	s.Body = append(s.Body, newHead)
}

func (s *Snake) MoveSnakeUp() {
	s.Body = s.Body[1:]

	head := s.GiveMeHead()

	newHead := Box{
		X: head.X,
		Y: head.Y - 1,
	}

	s.Body = append(s.Body, newHead)
}

func (s *Snake) MoveSnakeDown() {
	s.Body = s.Body[1:]

	head := s.GiveMeHead()

	newHead := Box{
		X: head.X,
		Y: head.Y + 1,
	}

	s.Body = append(s.Body, newHead)
}

func (s *Snake) GiveMeHead() Box {
	head := s.Body[len(s.Body)-1]
	return head
}
