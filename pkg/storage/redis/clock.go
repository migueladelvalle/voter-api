package cache

import "time"

type Clock interface {
	Now() time.Time
}

type realClock struct{}

func (c *realClock) Now() time.Time {
	return time.Now()
}

type mockClock struct {
	now time.Time
}

func (c *mockClock) Now() time.Time {
	return c.now
}
