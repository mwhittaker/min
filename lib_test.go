package min_test

import (
	"testing"

	"github.com/mwhittaker/min"
)

func TestMin(t *testing.T) {
	for _, test := range []struct{ x, y, want int }{
		{0, 0, 0},
		{10, 0, 0},
		{0, 10, 0},
	} {
		if got := min.Min(test.x, test.y); got != test.want {
			t.Errorf("Min(%d, %d): got %d; want %d", test.x, test.y, got, test.want)
		}
	}
}

func TestMax(t *testing.T) {
	for _, test := range []struct{ x, y, want int }{
		{0, 0, 0},
		{10, 0, 10},
		{0, 10, 10},
	} {
		if got := min.Max(test.x, test.y); got != test.want {
			t.Errorf("Max(%d, %d): got %d; want %d", test.x, test.y, got, test.want)
		}
	}
}
