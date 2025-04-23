package ex1

import (
	"testing"
)

func TestDiv(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
		err  error
	}{
		{"ok", 10, 2, 5, nil},
		{"div by zero", 5, 0, 0, ErrDivByZero},
		{"neg and pos nums", 6, -2, -3, nil},
		{"first num less then second", 2, 5, 0, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := Div(test.a, test.b)
			if err != nil {
				if err != test.err {
					t.Errorf("got %v, want %v", err, test.err)
				}
			} else if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})

	}
}
