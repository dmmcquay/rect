package rect

import "testing"

func TestSizeTriangle(t *testing.T) {
	var sizeTriangleTests = []struct {
		pts      []Point
		expected float64
	}{
		{[]Point{Point{-2, 3}, Point{-3, -1}, Point{3, -2}}, 12.5},
		{[]Point{Point{0, 0}, Point{1, 1}, Point{0, 1}}, 0.5},
		{[]Point{Point{10, 14}, Point{20, 15}, Point{12, 52}}, 189},
	}
	for _, rt := range sizeTriangleTests {
		actual := SizeTriangle(rt.pts[0], rt.pts[1], rt.pts[2])
		if actual != rt.expected {
			t.Errorf(
				"failed spiral:\n\texpected: %d\n\t  actual: %d",
				rt.expected,
				actual,
			)
		}

	}
}
