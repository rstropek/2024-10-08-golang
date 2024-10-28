package main

type Clock struct {
	hour   int
	minute int
}

func NewClock(hour int, minute int) Clock {
	return Clock{hour: hour, minute: minute}
}

func (c Clock) AddMinutesReturningClock(minutes int) Clock {
	return NewClock(c.hour, c.minute+minutes)
}

func (c *Clock) AddMinutes(minutes int) {
	c.minute += minutes
}
