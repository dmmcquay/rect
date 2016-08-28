package rect

import "math"

// Returns the size of the triangle formed by points p1, p2, and p3
func sizeTriangle(p1, p2, p3 Point) float64 {
	return math.Abs((p1.X*p2.Y + p2.X*p3.Y + p3.X*p1.Y - p1.Y*p2.X - p2.Y*p3.X - p3.Y*p1.X) / 2)
}
