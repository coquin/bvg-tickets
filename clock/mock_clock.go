package clock

import "time"

type MockClock struct {
	t time.Time
}

func Mock(t time.Time) MockClock {
	return MockClock{t}
}

func (m MockClock) Now() time.Time {
	return m.t
}
