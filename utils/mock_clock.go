package utils

import "time"

type MockClock struct {
	t time.Time
}

func NewMockClock(t time.Time) MockClock {
	return MockClock{t}
}

func (m MockClock) Now() time.Time {
	return m.t
}
