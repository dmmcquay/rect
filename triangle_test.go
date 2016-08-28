package rect

import "testing"

func TestSizeTriangle(t *testing.T) {
	var sizeTriangleTests = []struct {
		pts      []Point
		expected float64
	}{
		{[]Point{{-2, 3}, {-3, -1}, {3, -2}}, 12.5},
		{[]Point{{0, 0}, {1, 1}, {0, 1}}, 0.5},
		{[]Point{{10, 14}, {20, 15}, {12, 52}}, 189},
	}
	for _, rt := range sizeTriangleTests {
		actual := sizeTriangle(rt.pts[0], rt.pts[1], rt.pts[2])
		if actual != rt.expected {
			t.Errorf(
				"failed sizeTriangle:\n\texpected: %f\n\t  actual: %f",
				rt.expected,
				actual,
			)
		}

	}
}
