package rect

import "math"

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

func lineOnLine(l1, l2 line) bool {
	if onLine(l1.pt1, l1.pt2, l2.pt1) && onLine(l1.pt1, l1.pt2, l2.pt2) {
		return true
	}
	return false
}
