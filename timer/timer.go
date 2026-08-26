package timer

import (
	model "terminal-snake/models"
	"time"
)

type TimerDeps struct {
	Snake *model.Snake
}

type Timer struct {
	Now    time.Time
	Period time.Duration
	snake  *model.Snake
	moves  *int
}

func NewTimer(deps TimerDeps) Timer {
	return Timer{
		snake: deps.Snake,
	}
}

func (t *Timer) SetTimer() {
	moves := 0
	snake := t.snake

	for _, value := range snake.LastDirections {
		if value != "" {
			moves++
		}
	}
	if moves == 0 {
		moves = 1
	}

	lapTime := t.NewTime(len(snake.Body))

	t.moves = &moves
	t.Period = lapTime / time.Duration(moves) * time.Millisecond
	t.Now = time.Now()
}

func (t *Timer) SetMoves() {
	snake := t.snake
	for _, value := range snake.LastDirections {
		if value != "" {
			*t.moves++
		}
	}
	if *t.moves == 0 {
		*t.moves = 1
	}
}

// Make it newPeriod
func (t *Timer) NewTime(size int) time.Duration {
	if size < 15 {
		return (180)
	}
	if size < 25 {
		return 160
	}
	if size < 35 {
		return 140
	}
	if size < 45 {
		return 130
	}
	if size < 55 {
		return 110
	}
	if size < 65 {
		return 100
	}

	return 180
}
