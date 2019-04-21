package clock

import "fmt"

type Clock struct {
	hour int
	min int
}

const MINUTES_PER_DAY int = 1440
const MINUTES_PER_HOUR int = 1440

func New(hour, min int) Clock {
	// normalize totalMinutes to discard whole days
	totalMinutes := (hour * MINUTES_PER_HOUR + min) % MINUTES_PER_DAY
	if totalMinutes < 0 {
		totalMinutes += MINUTES_PER_DAY
	}
	h := (totalMinutes - totalMinutes % MINUTES_PER_HOUR) / MINUTES_PER_HOUR
	m := totalMinutes - h * MINUTES_PER_HOUR
	return Clock{h, m}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.min)
}

func (c Clock) Add(min int) Clock {
	return New(c.hour, c.min + min)
}

func (c Clock) Subtract(min int) Clock {
	return New(c.hour, c.min - min)
}