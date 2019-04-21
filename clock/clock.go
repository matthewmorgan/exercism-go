package clock

import "fmt"

type Clock struct {
	hour int
	min int
}

const MINUTES_IN_DAY int = 1440

func New(hour, min int) Clock {
	// normalize hoursAndMinutesInMinutes to discard whole days
	hoursAndMinutesInMinutes := (hour * 60 + min) % MINUTES_IN_DAY
	if hoursAndMinutesInMinutes < 0 {
		hoursAndMinutesInMinutes += 1440
	}
	h := (hoursAndMinutesInMinutes - hoursAndMinutesInMinutes % 60) / 60
	m := hoursAndMinutesInMinutes - h * 60
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