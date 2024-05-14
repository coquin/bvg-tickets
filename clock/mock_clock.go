package clock

import "time"

type mockClock struct {
	t time.Time
}

func Mock(t time.Time) mockClock {
	return mockClock{t}
}

func (m mockClock) Now() time.Time {
	return m.t
}
