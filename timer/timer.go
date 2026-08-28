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

// SetPeriod derives the time per move from the snake's queued directions and
// length. More queued moves means a shorter period (faster movement).
func (t *Timer) SetPeriod() {
	moves := t.snake.DirectionCount()
	lapTime := t.newPeriod(len(t.snake.Body))

	t.Period = lapTime / time.Duration(moves)
}

func (t *Timer) SetTimer() {
	t.Now = time.Now()
}

func (t *Timer) newPeriod(size int) time.Duration {
	switch {
	case size < 40:
		return 220 * time.Millisecond
	case size < 60:
		return 200 * time.Millisecond
	case size < 80:
		return 180 * time.Millisecond
	case size < 100:
		return 160 * time.Millisecond
	case size < 120:
		return 150 * time.Millisecond
	case size < 140:
		return 130 * time.Millisecond
	case size < 160:
		return 110 * time.Millisecond
	case size < 180:
		return 90 * time.Millisecond
	default:
		return 180 * time.Millisecond
	}
}

func (t *Timer) Reset(snake *model.Snake) {
	t.Period = t.newPeriod(len(snake.Body))
	t.Now = time.Now()
}
