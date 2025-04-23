package ex7

import (
	"time"
)

type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (c *RealClock) Now() time.Time {
	return time.Now()
}

func GreetingWithClock(clock Clock) string {
	hour := clock.Now().Hour()
	switch {
	case hour < 12:
		return "Доброе утро!"
	case hour < 18:
		return "Добрый день!"
	default:
		return "Добрый вечер!"
	}
}

type mockClock struct {
	t time.Time
}

func (m *mockClock) Now() time.Time {
	return m.t
}
