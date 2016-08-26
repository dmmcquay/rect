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
		P1: rect.Point{0, 0},
		P2: rect.Point{0, 1},
		P3: rect.Point{1, 0},
		P4: rect.Point{1, 1},
	}
	fmt.Println(r1.IsRect())
	r1 := rect.Rectangle{
		P1: rect.Point{0, 0},
		P2: rect.Point{0, 1},
		P3: rect.Point{1, 0},
		P4: rect.Point{1, 1},
	}
	fmt.Println(r1.IsRect())
}
