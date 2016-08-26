package rect

import "testing"

func TestIsRect(t *testing.T) {
	var isRectTests = []struct {
		r        Rectangle
		expected bool
	}{
		{Rectangle{P1: Point{1, 1}, P2: Point{1, 2}, P3: Point{2, 1}, P4: Point{2, 2}}, true},
		{Rectangle{P1: Point{0, 0}, P2: Point{0, 1}, P3: Point{1, 0}, P4: Point{1, 1}}, true},
		{Rectangle{P1: Point{0, 0}, P2: Point{0, -1}, P3: Point{-1, 0}, P4: Point{-1, -1}}, true},
		{Rectangle{P1: Point{1.5, 1.5}, P2: Point{1.5, 3.5}, P3: Point{3.5, 1.5}, P4: Point{3.5, 3.5}}, true},
		{Rectangle{P1: Point{0, 0}, P2: Point{0, 3}, P3: Point{2, 0}, P4: Point{2, 3}}, true},
		{Rectangle{P1: Point{0, 0}, P2: Point{0, 100}, P3: Point{-1, 0}, P4: Point{0, 23}}, false},
		{Rectangle{P1: Point{0, 0}, P2: Point{0, 0}, P3: Point{0, 0}, P4: Point{0, 0}}, false},
	}
	for _, rt := range isRectTests {
		actual := rt.r.IsRect()
		if actual != rt.expected {
			t.Errorf(
				"failed spiral:\n\texpected: %d\n\t  actual: %d",
				rt.expected,
				actual,
			)
		}

	}
}
