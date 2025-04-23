package ex3

import "testing"

func TestSliceFunc(t *testing.T) {
	tests := []struct {
		name string
		sl   []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{5}, 5},
		{"multiple elements", []int{1, 2, 3, 4, 5}, 15},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			answer := SliceFunc(test.sl)
			if answer != test.want {
				t.Errorf("got %v, want %v", answer, test.want)

			}
		})
	}
}
