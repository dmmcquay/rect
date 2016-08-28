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

func TestOnLine(t *testing.T) {
	var onLineTest = []struct {
		p        []Point
		expected bool
	}{
		{[]Point{Point{1, 1}, Point{3, 3}, Point{2, 2}}, true},
		{[]Point{Point{1, 1}, Point{3, 3}, Point{4, 4}}, false},
		{[]Point{Point{-1, -1}, Point{-3, -3}, Point{-4, -4}}, false},
		{[]Point{Point{-1, -1}, Point{-3, -3}, Point{-2, -2}}, true},
	}
	for _, rt := range onLineTest {
		actual := onLine(rt.p[0], rt.p[1], rt.p[2])
		if actual != rt.expected {
			t.Errorf(
				"failed onLine:\n\texpected: %t\n\t  actual: %t",
				rt.expected,
				actual,
			)
		}

	}
}

func TestLineIntersection(t *testing.T) {
	var lineIntersectionTest = []struct {
		l        []line
		expected Point
	}{
		{[]line{line{Point{0, 0}, Point{2, 2}}, line{Point{2, 0}, Point{0, 2}}}, Point{1, 1}},
		{[]line{line{Point{0, 0}, Point{-2, -2}}, line{Point{-2, 0}, Point{0, -2}}}, Point{-1, -1}},
		{[]line{line{Point{0, 0}, Point{-2, -2}}, line{Point{1, 0}, Point{1, -7}}}, Point{0, 0}},
		{[]line{line{Point{5, 8}, Point{8, 5}}, line{Point{3, 7}, Point{7, 3}}}, Point{0, 0}},
	}
	for _, rt := range lineIntersectionTest {
		actual, _ := lineIntersection(rt.l[0], rt.l[1])
		if actual != rt.expected {
			t.Errorf(
				"failed lineIntersection:\n\texpected: %v\n\t  actual: %v",
				rt.expected,
				actual,
			)
		}

	}
}
