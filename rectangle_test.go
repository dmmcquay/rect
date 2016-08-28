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
		actual := rt.r.isRect()
		if actual != rt.expected {
			t.Errorf(
				"failed isRect:\n\texpected: %t\n\t  actual: %t",
				rt.expected,
				actual,
			)
		}

	}
}

func TestSize(t *testing.T) {
	var isRectTests = []struct {
		r        Rectangle
		expected float64
	}{
		{Rectangle{P1: Point{1, 1}, P2: Point{1, 2}, P3: Point{2, 1}, P4: Point{2, 2}}, 1},
		{Rectangle{P1: Point{0, 0}, P2: Point{0, 1}, P3: Point{1, 0}, P4: Point{1, 1}}, 1},
		{Rectangle{P1: Point{0, 0}, P2: Point{0, -1}, P3: Point{-1, 0}, P4: Point{-1, -1}}, 1},
		{Rectangle{P1: Point{1.5, 1.5}, P2: Point{1.5, 3.5}, P3: Point{3.5, 1.5}, P4: Point{3.5, 3.5}}, 4},
		{Rectangle{P1: Point{0, 0}, P2: Point{0, 3}, P3: Point{2, 0}, P4: Point{2, 3}}, 6},
		{Rectangle{P1: Point{0, 0}, P2: Point{0, 100}, P3: Point{100, 0}, P4: Point{100, 100}}, 10000},
	}
	for _, rt := range isRectTests {
		actual := rt.r.size()
		if actual != rt.expected {
			t.Errorf(
				"failed size:\n\texpected: %f\n\t  actual: %f",
				rt.expected,
				actual,
			)
		}

	}
}

func TestContainment(t *testing.T) {
	var containmentTest = []struct {
		r        []Rectangle
		expected bool
	}{
		{
			[]Rectangle{
				Rectangle{Point{0, 0}, Point{0, 3}, Point{3, 0}, Point{3, 3}},
				Rectangle{Point{1, 1}, Point{1, 2}, Point{2, 1}, Point{2, 2}}},
			true,
		},
		{
			[]Rectangle{
				Rectangle{Point{0, 0}, Point{0, 3}, Point{3, 0}, Point{3, 3}},
				Rectangle{Point{1, 1}, Point{1, 2}, Point{2, 1}, Point{2, 2}}},
			true,
		},
		{
			[]Rectangle{
				Rectangle{Point{0, 0}, Point{0, 3}, Point{3, 0}, Point{3, 3}},
				Rectangle{Point{0, 0}, Point{0, 3}, Point{3, 0}, Point{3, 3}}},
			false,
		},
		{
			[]Rectangle{
				Rectangle{Point{0, 0}, Point{0, 3}, Point{3, 0}, Point{3, 3}},
				Rectangle{Point{3, 3}, Point{3, 4}, Point{4, 3}, Point{4, 4}}},
			false,
		},
	}
	for _, rt := range containmentTest {
		actual := Containment(rt.r[0], rt.r[1])
		if actual != rt.expected {
			t.Errorf(
				"failed Containment:\n\texpected: %t\n\t  actual: %t",
				rt.expected,
				actual,
			)
		}

	}
}

func TestAdjacency(t *testing.T) {
	var adjacencyTests = []struct {
		r        []Rectangle
		expected bool
	}{
		{
			[]Rectangle{
				Rectangle{Point{0, 0}, Point{0, 3}, Point{3, 0}, Point{3, 3}},
				Rectangle{Point{2, 2}, Point{2, 3}, Point{3, 2}, Point{3, 3}}},
			true,
		},
		{
			[]Rectangle{
				Rectangle{Point{0, 0}, Point{0, 3}, Point{3, 0}, Point{3, 3}},
				Rectangle{Point{1, 3}, Point{2, 3}, Point{1, 4}, Point{2, 4}}},
			true,
		},
		{
			[]Rectangle{
				Rectangle{Point{2, 2}, Point{2, 3}, Point{3, 2}, Point{3, 3}},
				Rectangle{Point{1, 3}, Point{2, 3}, Point{1, 4}, Point{2, 4}}},
			false,
		},
		{
			[]Rectangle{
				Rectangle{Point{-2, -2}, Point{-2, -3}, Point{-3, -2}, Point{-3, -3}},
				Rectangle{Point{-1, -3}, Point{-2, -3}, Point{-1, -4}, Point{-2, -4}}},
			false,
		},
		{
			[]Rectangle{
				Rectangle{Point{0, 0}, Point{0, -3}, Point{-3, 0}, Point{-3, -3}},
				Rectangle{Point{-2, -2}, Point{-2, -3}, Point{-3, -2}, Point{-3, -3}}},
			true,
		},
	}
	for _, rt := range adjacencyTests {
		actual := Adjacency(rt.r[0], rt.r[1])
		if actual != rt.expected {
			t.Errorf(
				"failed adjacency:\n\texpected: %t\n\t  actual: %t",
				rt.expected,
				actual,
			)
		}

	}
}

func TestIntersection(t *testing.T) {
	var intersectionTests = []struct {
		r        []Rectangle
		expected []Point
	}{
		{
			[]Rectangle{
				Rectangle{Point{0, 0}, Point{0, 3}, Point{3, 0}, Point{3, 3}},
				Rectangle{Point{2, 1}, Point{2, 2}, Point{4, 1}, Point{4, 2}}},
			[]Point{Point{3, 1}, Point{3, 2}},
		},
		{
			[]Rectangle{
				Rectangle{Point{0, 0}, Point{0, 3}, Point{3, 0}, Point{3, 3}},
				Rectangle{Point{0, 0}, Point{0, -1}, Point{-1, 0}, Point{-1, -1}}},
			[]Point{Point{0, 0}, Point{0, 0}},
		},
		{
			[]Rectangle{
				Rectangle{Point{2, 1}, Point{2, 2}, Point{4, 1}, Point{4, 2}},
				Rectangle{Point{0, 0}, Point{0, -1}, Point{-1, 0}, Point{-1, -1}}},
			[]Point{},
		},
	}
	for _, rt := range intersectionTests {
		actual := Intersection(rt.r[0], rt.r[1])
		if len(actual) != len(rt.expected) {
			t.Errorf(
				"failed intersection:\n\texpected: %d\n\t  actual: %d",
				len(rt.expected),
				len(actual),
			)
			continue
		}
		for i := range actual {
			if actual[i] != rt.expected[i] {
				t.Errorf(
					"failed intersection:\n\texpected: %v\n\t  actual: %v",
					rt.expected[i],
					actual[i],
				)
			}
		}

	}
}
