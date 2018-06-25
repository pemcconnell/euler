package main

// this feels like a really slow solution to the problem, but a solution none-
// the-less.

import "fmt"

type Fifteen struct {
	Width  int
	Height int
	EndX   int
	EndY   int
	Routes int
}

var A Fifteen

func main() {
	// set parameters
	A.Height = 20
	A.Width = 20
	A.EndX = A.Height
	A.EndY = A.Width
	A.Width = 2
	A.Routes = 0

	// start to move, specifying start x & y
	A.move(0, 0)

	fmt.Printf("I found %+v routes", A.Routes)
}

func (f *Fifteen) move(x, y int) {
	if x == f.EndX && y == f.EndY {
		// arrived at the destination
		f.Routes++
	} else {
		// if we can step right, step right
		if x < f.EndX {
			f.stepRight(x, y)
		}
		// if we can step down, step down
		if y < f.EndY {
			f.stepDown(x, y)
		}
	}
}

func (f *Fifteen) stepRight(x, y int) {
	f.move((x + 1), y)
}

func (f *Fifteen) stepDown(x, y int) {
	f.move(x, (y + 1))
}
