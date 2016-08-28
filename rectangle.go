package rect

import (
	"math"
	"sort"
)

// Rectangle struct defines a plane figure with four straight sides
// and four right angles, which contains 4 vertixes points, P1 through P4
type Rectangle struct {
	P1, P2, P3, P4 Point
}

// maximum error used for floating point math
var ɛ = 0.00001

// isRect determins if the rectangle provided really is a rectangle, which
// by definition means a plane figure with four straight sides and four
// right angles.
func (r Rectangle) isRect() bool {
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
	return math.Abs(dd1-dd2) < ɛ && math.Abs(dd1-dd3) < ɛ && math.Abs(dd1-dd4) < ɛ
}

// neighbors returns the two points that form the line which creates the right
// angle of the point passed in as a parameter.
func (r Rectangle) neighbors(p Point) (Point, Point) {
	keys := []float64{distance(r.P1, p), distance(r.P2, p), distance(r.P3, p), distance(r.P4, p)}
	sort.Float64s(keys)
	n := []Point{}
	d := distance(r.P1, p)
	if math.Abs(keys[1]-d) < ɛ || math.Abs(keys[2]-d) < ɛ {
		n = append(n, r.P1)
	}
	d = distance(r.P2, p)
	if math.Abs(keys[1]-d) < ɛ || math.Abs(keys[2]-d) < ɛ {
		n = append(n, r.P2)
	}
	d = distance(r.P3, p)
	if math.Abs(keys[1]-d) < ɛ || math.Abs(keys[2]-d) < ɛ {
		n = append(n, r.P3)
	}
	d = distance(r.P4, p)
	if math.Abs(keys[1]-d) < ɛ || math.Abs(keys[2]-d) < ɛ {
		n = append(n, r.P4)
	}
	return n[0], n[1]
}

// size returns the size of the Rectangle
func (r Rectangle) size() float64 {
	n1, n2 := r.neighbors(r.P1)
	return distance(r.P1, n1) * distance(r.P1, n2)
}

// inOrder returns a []Point, let us call pts, in which the four sides of the
// Rectangle can be defined as pts[0] -- pts[1], pts[0] -- pts[2],
// pts[3] -- pts[1], and pts[3] -- pts[2]
func (r Rectangle) inOrder() []Point {
	n1, n2 := r.neighbors(r.P1)
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

// Adjacency detects whether two rectangles, r1, and r2, are adjacent.
// Adjacency is defined as the sharing of a side
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

// Intersection determine whether two rectangles, r1 and r2, have one or more
// intersecting lines and produce a result, []Point, identifying the points
// of intersection
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

// sumOfTri determines if the sum of the areas of the triangles created from
// point p to two neighboring vertices of a side of the Rectangle, r, add up
// to the same size as the Rectangle, r.  This is one way to determine if
// a point is inside of a Rectangle.
func sumOfTri(r Rectangle, p Point) bool {
	n1, n2 := r.neighbors(r.P1)
	x1, x2 := Point{}, Point{}
	accross := &Point{}
	if r.P2 != n1 || r.P2 != n2 {
		accross = &r.P2
		x1, x2 = r.neighbors(r.P2)
	}
	if r.P3 != n1 || r.P3 != n2 {
		accross = &r.P3
		x1, x2 = r.neighbors(r.P3)
	}
	if r.P4 != n1 || r.P4 != n2 {
		accross = &r.P4
		x1, x2 = r.neighbors(r.P4)
	}
	sumTri := sizeTriangle(r.P1, n1, p) +
		sizeTriangle(r.P1, n2, p) +
		sizeTriangle(*accross, x1, p) +
		sizeTriangle(*accross, x2, p)
	if math.Abs(r.size()-sumTri) < ɛ {
		return true
	}
	return false
}

// Containment returns whether r2 is contained inside of r1
func Containment(r1, r2 Rectangle) bool {
	if r1.size() <= r2.size() {
		return false
	}
	if sumOfTri(r1, r2.P1) && sumOfTri(r1, r2.P2) && sumOfTri(r1, r2.P3) && sumOfTri(r1, r2.P4) {
		return true
	}
	return false
}
