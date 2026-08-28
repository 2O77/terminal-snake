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

// SetPeriod derives the time per move from the snake's length only.
// Speed is score-based and independent of input frequency.
func (t *Timer) SetPeriod() {
	t.Period = t.newPeriod(len(t.snake.Body))
}

func (t *Timer) SetTimer() {
	t.Now = time.Now()
}

func (t *Timer) newPeriod(size int) time.Duration {
	switch {
	case size < 10:
		return 260 * time.Millisecond
	case size < 20:
		return 240 * time.Millisecond
	case size < 30:
		return 220 * time.Millisecond
	case size < 40:
		return 200 * time.Millisecond
	case size < 50:
		return 190 * time.Millisecond
	case size < 60:
		return 180 * time.Millisecond
	default:
		return 170 * time.Millisecond
	}
}

func (t *Timer) Reset(snake *model.Snake) {
	t.Period = t.newPeriod(len(snake.Body))
	t.Now = time.Now()
}
