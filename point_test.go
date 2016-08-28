package rect

import (
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	var distanceTest = []struct {
		p        []Point
		expected float64
	}{
		{[]Point{Point{1, 1}, Point{4, 5}}, 5},
		{[]Point{Point{-1, -1}, Point{2, 3}}, 5},
		{[]Point{Point{1, 1}, Point{2, 2}}, math.Sqrt(2)},
		{[]Point{Point{1, 1}, Point{40, 20}}, math.Sqrt(1882)},
	}
	for _, rt := range distanceTest {
		actual := distance(rt.p[0], rt.p[1])
		if actual != rt.expected {
			t.Errorf(
				"failed distance:\n\texpected: %f\n\t  actual: %f",
				rt.expected,
				actual,
			)
		}

	}
}
