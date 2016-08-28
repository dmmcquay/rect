# rect

[![GoDoc](https://godoc.org/s.mcquay.me/dm/rect?status.svg)](https://godoc.org/s.mcquay.me/dm/rect)
[![Go Report Card](https://goreportcard.com/badge/s.mcquay.me/dm/rect)](https://goreportcard.com/report/s.mcquay.me/dm/rect)

Package rect implements rectangles and can determine when rectangles
intersect, are contained, and when rectangles lie adjacent

## usage

in cmd/recty/main.go is an example of how one could use this package.

``` go
r := rect.Rectangle{     
    P1: rect.Point{1, 1},
    P2: rect.Point{1, 2},
    P3: rect.Point{2, 1},
    P4: rect.Point{2, 2},
}                        
```

defines a rectangle.

``` go
fmt.Println(r.IsRect())
```

would varify that the rectangle provided is a valid rectangle.

``` go
r1 := rect.Rectangle{     
    P1: rect.Point{1, 1},
    P2: rect.Point{1, 2},
    P3: rect.Point{2, 1},
    P4: rect.Point{2, 2},
}                        
r2 := rect.Rectangle{     
    P1: rect.Point{1, 1},
    P2: rect.Point{1, 2},
    P3: rect.Point{2, 1},
    P4: rect.Point{2, 2},
}                        

fmt.Println(rect.Intersection(r1, r2))
```

Running the above code would print out all intersection points of the two 
rectangles.

``` go
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
```

This would return a boolean for whether r3 and r4 have an adjacent side.

``` go
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
```

And finally, running `rect.Containment(r5, r6)` would return true because r6
is contained inside of r5
