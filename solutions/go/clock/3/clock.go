package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hours   int
	minutes int
}

func New(h, m int) Clock {

	if m < 0 {
		h -= 1 - m/60
		m = m%60 + 60
	}

	if h < 0 {
		h = h%24 + 24
	}

	if m/60 > 0 {
		h += m / 60
		m = m % 60
	}

	return Clock{hours: h % 24, minutes: m}
}

func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes  + m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
