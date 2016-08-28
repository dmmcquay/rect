package rect

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type line struct {
	pt1, pt2 Point
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow((p1.X-p2.X), 2) + math.Pow((p1.Y-p2.Y), 2))
}

func onLine(rp1, rp2, p Point) bool {
	if distance(rp1, p)+distance(rp2, p) == distance(rp1, rp2) {
		minx, maxx := math.Min(rp1.X, rp2.X), math.Max(rp1.X, rp2.X)
		miny, maxy := math.Min(rp1.Y, rp2.Y), math.Max(rp1.Y, rp2.Y)
		if minx <= p.X && p.X <= maxx && miny <= p.Y && p.Y <= maxy {
			return true
		}
	}
	return false
}

func lineHelper(l line) (float64, float64, float64) {
	a := l.pt2.Y - l.pt1.Y
	b := l.pt1.X - l.pt2.X
	c := a*l.pt1.X + b*l.pt1.Y
	return a, b, c
}

func lineIntersection(l1, l2 line) (Point, error) {
	a1, b1, c1 := lineHelper(l1)
	a2, b2, c2 := lineHelper(l2)
	det := a1*b2 - a2*b1

	if det != 0 {
		x := ((b2*c1 - b1*c2) / det)
		y := ((a1*c2 - a2*c1) / det)
		minx1, maxx1 := math.Min(l1.pt1.X, l1.pt2.X), math.Max(l1.pt1.X, l1.pt2.X)
		miny1, maxy1 := math.Min(l1.pt1.Y, l1.pt2.Y), math.Max(l1.pt1.Y, l1.pt2.Y)
		minx2, maxx2 := math.Min(l2.pt1.X, l2.pt2.X), math.Max(l2.pt1.X, l2.pt2.X)
		miny2, maxy2 := math.Min(l2.pt1.Y, l2.pt2.Y), math.Max(l2.pt1.Y, l2.pt2.Y)
		if minx1 <= x && x <= maxx1 && miny1 <= y && y <= maxy1 {
			if minx2 <= x && x <= maxx2 && miny2 <= y && y <= maxy2 {
				return Point{x, y}, nil
			}
		}
	}
	return Point{}, fmt.Errorf("lines do not intersect")
}

func lineOnLine(l1, l2 line) bool {
	if onLine(l1.pt1, l1.pt2, l2.pt1) && onLine(l1.pt1, l1.pt2, l2.pt2) {
		return true
	}
	return false
}
