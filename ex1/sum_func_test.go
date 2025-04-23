package ex1

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"pos_nums", 2, 3, 5},
		{"neg_nums", -2, -3, -5},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Sum(test.a, test.b)
			if got != test.want {
				t.Errorf("got %v, want %v", got, test.want)
			}
		})

	}
}
