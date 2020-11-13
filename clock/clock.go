package clock

import "fmt"

const minutesInOneHour = 60
const minutesInOneDay = 24 * minutesInOneHour

type Clock struct {
	minutes int
}

func New(h int, m int) Clock {
	return Clock{
		minutes: normalizeMinutes(h*minutesInOneHour + m),
	}
}

func (c Clock) Add(minutes int) Clock {
	c.minutes = normalizeMinutes(c.minutes + minutes)
	return c
}

func (c Clock) Subtract(minutes int) Clock {
	c.minutes = normalizeMinutes(c.minutes - minutes)
	return c
}

func (c Clock) String() string {
	minutes := c.minutes % minutesInOneDay

	if minutes < 0 {
		minutes += minutesInOneDay
	}

	return fmt.Sprintf("%02d:%02d", minutes/minutesInOneHour%24, minutes%minutesInOneHour)
}

func normalizeMinutes(minutes int) int {
	minutes = minutes % minutesInOneDay

	if minutes < 0 {
		minutes += minutesInOneDay
	}

	return minutes
}
