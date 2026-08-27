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

func (t *Timer) newPeriod(size int) time.Duration {
	if size < 20 {
		return 220
	}
	if size < 40 {
		return 220
	}
	if size < 60 {
		return 200
	}
	if size < 80 {
		return 180
	}
	if size < 100 {
		return 160
	}
	if size < 120 {
		return 150
	}
	if size < 140 {
		return 130
	}
	if size < 160 {
		return 110
	}
	if size < 180 {
		return 90
	}

	return 180
}

func (t *Timer) Reset(snake model.Snake) {
	lapTime := t.newPeriod(len(snake.Body))
	t.Period = lapTime * time.Millisecond
	t.Now = time.Now()
}
