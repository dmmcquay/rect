package main

import (
	"fmt"

	"s.mcquay.me/dm/rect"
)

func main() {
	r := rect.Rectangle{
		P1: rect.Point{1, 1},
		P2: rect.Point{1, 2},
		P3: rect.Point{2, 1},
		P4: rect.Point{2, 2},
	}

	fmt.Println(r.IsRect())

	r1 := rect.Rectangle{
		P1: rect.Point{1, 1},
		P2: rect.Point{1, 2},
		P3: rect.Point{2, 1},
		P4: rect.Point{2, 2},
	}
	r2 := rect.Rectangle{
		P1: rect.Point{2, 2},
		P2: rect.Point{2, 3},
		P3: rect.Point{3, 2},
		P4: rect.Point{3, 3},
	}

	fmt.Println(rect.Intersection(r1, r2))

	r3 := rect.Rectangle{
		P1: rect.Point{0, 0},
		P2: rect.Point{0, 3},
		P3: rect.Point{3, 0},
		P4: rect.Point{3, 3},
	}
	r4 := rect.Rectangle{
		P1: rect.Point{2, 2},
		P2: rect.Point{2, 3},
		P3: rect.Point{3, 2},
		P4: rect.Point{3, 3},
	}

	fmt.Println(rect.Adjacency(r3, r4))

	r5 := rect.Rectangle{
		P1: rect.Point{0, 0},
		P2: rect.Point{0, 3},
		P3: rect.Point{3, 0},
		P4: rect.Point{3, 3},
	}
	r6 := rect.Rectangle{
		P1: rect.Point{1, 1},
		P2: rect.Point{1, 2},
		P3: rect.Point{2, 1},
		P4: rect.Point{2, 2},
	}

	fmt.Println(rect.Containment(r5, r6))

}
