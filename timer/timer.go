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
}

func NewTimer(deps TimerDeps) Timer {
	return Timer{
		snake: deps.Snake,
	}
}

func (t *Timer) SetPeriod() {
	moves := t.snake.DirectionCount()
	snake := t.snake

	for _, value := range snake.LastDirections {
		if value != "" {
			moves++
		}
	}
	if moves == 0 {
		moves = 1
	}

	lapTime := t.newPeriod(len(snake.Body))

	t.Period = lapTime / time.Duration(moves) * time.Millisecond
}

func (t *Timer) SetTimer() {
	t.Now = time.Now()
}

// Make it newPeriod
func (t *Timer) newPeriod(size int) time.Duration {
	if size < 15 {
		return (180)
	}
	if size < 35 {
		return 160
	}
	if size < 55 {
		return 140
	}
	if size < 75 {
		return 130
	}
	if size < 95 {
		return 120
	}
	if size < 115 {
		return 110
	}

	return 180
}

func (t *Timer) Reset(snake model.Snake) {
	lapTime := t.newPeriod(len(snake.Body))
	t.Period = lapTime * time.Millisecond
	t.Now = time.Now()
}
