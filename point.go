package rect

import "math"

// Point is a struct defining a coordinates position on a 2D plane
type Point struct {
	X, Y float64
}

// distance calculates the distance between two points
func distance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow((p1.X-p2.X), 2) + math.Pow((p1.Y-p2.Y), 2))
}

// onLine determines if point p lies on the line formed by points rp1 and rp2
func onLine(rp1, rp2, p Point) bool {
	if math.Abs(distance(rp1, p)+distance(rp2, p)-distance(rp1, rp2)) < ɛ {
		minx, maxx := math.Min(rp1.X, rp2.X), math.Max(rp1.X, rp2.X)
		miny, maxy := math.Min(rp1.Y, rp2.Y), math.Max(rp1.Y, rp2.Y)
		if minx <= p.X && p.X <= maxx && miny <= p.Y && p.Y <= maxy {
			return true
		}
	}
	return false
}
