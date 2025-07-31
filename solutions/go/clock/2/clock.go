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
	// newMin := c.minutes + m
	//
	// if newMin/60 < 1 {
	// 	c.minutes = newMin
	// 	return c
	// }
	//
	// c.hours += newMin / 60
	//
	// if c.hours >= 24 {
	// 	c.hours = c.hours%24
	// }
	//
	// c.minutes = newMin % 60
	//
	// return c
}

func (c Clock) Subtract(m int) Clock {
	newMin := c.minutes - m
	c.minutes = newMin

	if newMin >= 0 {
		return c
	}

	c.hours -= ((c.minutes / 60) * -1) + 1

	if c.hours < 0 {
		c.hours = (24 + c.hours)%24
	}
	

	c.minutes = 60 + (c.minutes/60 * -60) + c.minutes

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
