package ex7

import (
	"testing"
	"time"
)

func TestGreetingWithClock(t *testing.T) {
	tests := []struct {
		name  string
		clock Clock
		want  string
	}{
		{
			name:  "Утро",
			clock: &mockClock{time.Date(2025, 1, 1, 9, 0, 0, 0, time.UTC)},
			want:  "Доброе утро!",
		},
		{
			name:  "День",
			clock: &mockClock{time.Date(2025, 1, 1, 15, 0, 0, 0, time.UTC)},
			want:  "Добрый день!",
		},
		{
			name:  "Вечер",
			clock: &mockClock{time.Date(2025, 1, 1, 20, 0, 0, 0, time.UTC)},
			want:  "Добрый вечер!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			answer := GreetingWithClock(test.clock)
			if answer != test.want {
				t.Errorf("got %v, want %v", answer, test.want)
			}
		})
	}
}
