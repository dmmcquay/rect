package rect

import "math"

type Rectangle struct {
	P1, P2, P3, P4 Point
}

func (r Rectangle) IsRect() bool {
	// make sure they aren't all just the same point
	if (r.P1.X == r.P2.X && r.P1.X == r.P3.X && r.P1.X == r.P4.X) &&
		(r.P1.Y == r.P2.Y && r.P1.Y == r.P3.Y && r.P1.Y == r.P4.Y) {
		return false
	}

	cx := (r.P1.X + r.P2.X + r.P3.X + r.P4.X) / 4.0
	cy := (r.P1.Y + r.P2.Y + r.P3.Y + r.P4.Y) / 4.0

	dd1 := math.Sqrt(math.Abs(cx-r.P1.X)) + math.Sqrt(math.Abs(cy-r.P1.Y))
	dd2 := math.Sqrt(math.Abs(cx-r.P2.X)) + math.Sqrt(math.Abs(cy-r.P2.Y))
	dd3 := math.Sqrt(math.Abs(cx-r.P3.X)) + math.Sqrt(math.Abs(cy-r.P3.Y))
	dd4 := math.Sqrt(math.Abs(cx-r.P4.X)) + math.Sqrt(math.Abs(cy-r.P4.Y))
	return dd1 == dd2 && dd1 == dd3 && dd1 == dd4
}
