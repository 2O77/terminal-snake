package timer

import "time"

func NewTime(size int) time.Duration {
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
