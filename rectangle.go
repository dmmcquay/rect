package rect

import (
	"math"
	"sort"
)

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

func (r Rectangle) Neighbors(p Point) (Point, Point) {
	keys := []float64{distance(r.P1, p), distance(r.P2, p), distance(r.P3, p), distance(r.P4, p)}
	sort.Float64s(keys)
	n := []Point{}
	d := distance(r.P1, p)
	if keys[1] == d || keys[2] == d {
		n = append(n, r.P1)
	}
	d = distance(r.P2, p)
	if keys[1] == d || keys[2] == d {
		n = append(n, r.P2)
	}
	d = distance(r.P3, p)
	if keys[1] == d || keys[2] == d {
		n = append(n, r.P3)
	}
	d = distance(r.P4, p)
	if keys[1] == d || keys[2] == d {
		n = append(n, r.P4)
	}
	return n[0], n[1]
}

func (r Rectangle) Size() float64 {
	n1, n2 := r.Neighbors(r.P1)
	return distance(r.P1, n1) * distance(r.P1, n2)
}

func (r Rectangle) inOrder() []Point {
	n1, n2 := r.Neighbors(r.P1)
	accross := &Point{}
	if r.P2 != n1 || r.P2 != n2 {
		accross = &r.P2
	}
	if r.P3 != n1 || r.P3 != n2 {
		accross = &r.P3
	}
	if r.P4 != n1 || r.P4 != n2 {
		accross = &r.P4
	}
	return []Point{r.P1, n1, n2, *accross}
}

func Adjacency(r1, r2 Rectangle) bool {
	order1 := r1.inOrder()
	order2 := r2.inOrder()

	sides1 := []line{
		line{order1[0], order1[1]},
		line{order1[0], order1[2]},
		line{order1[3], order1[1]},
		line{order1[3], order1[2]},
	}
	sides2 := []line{
		line{order2[0], order2[1]},
		line{order2[0], order2[2]},
		line{order2[3], order2[1]},
		line{order2[3], order2[2]},
	}

	for _, i := range sides1 {
		for _, j := range sides2 {
			if lineOnLine(i, j) {
				return true
			}
		}
	}
	return false
}

func Intersection(r1, r2 Rectangle) []Point {
	order1 := r1.inOrder()
	order2 := r2.inOrder()

	sides1 := []line{
		line{order1[0], order1[1]},
		line{order1[0], order1[2]},
		line{order1[3], order1[1]},
		line{order1[3], order1[2]},
	}
	sides2 := []line{
		line{order2[0], order2[1]},
		line{order2[0], order2[2]},
		line{order2[3], order2[1]},
		line{order2[3], order2[2]},
	}

	pts := []Point{}
	for _, i := range sides1 {
		for _, j := range sides2 {
			p, err := lineIntersection(i, j)
			if err == nil {
				pts = append(pts, p)
			}
		}
	}
	return pts
}

func sumOfTri(r Rectangle, p Point) bool {
	n1, n2 := r.Neighbors(r.P1)
	x1, x2 := Point{}, Point{}
	accross := &Point{}
	if r.P2 != n1 || r.P2 != n2 {
		accross = &r.P2
		x1, x2 = r.Neighbors(r.P2)
	}
	if r.P3 != n1 || r.P3 != n2 {
		accross = &r.P3
		x1, x2 = r.Neighbors(r.P3)
	}
	if r.P4 != n1 || r.P4 != n2 {
		accross = &r.P4
		x1, x2 = r.Neighbors(r.P4)
	}
	sumTri := SizeTriangle(r.P1, n1, p) +
		SizeTriangle(r.P1, n2, p) +
		SizeTriangle(*accross, x1, p) +
		SizeTriangle(*accross, x2, p)
	if r.Size() == sumTri {
		return true
	}
	return false
}

// Containment returns whether r2 is contained inside of r1
func Containment(r1, r2 Rectangle) bool {
	if r1.Size() <= r2.Size() {
		return false
	}
	if sumOfTri(r1, r2.P1) && sumOfTri(r1, r2.P2) && sumOfTri(r1, r2.P3) && sumOfTri(r1, r2.P4) {
		return true
	}
	return false
}
