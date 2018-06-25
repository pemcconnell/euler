package main

import "fmt"

func main() {
	calc(1.0)
}

func calc(a float64) {
	// a + b + c = 1000; a < b < c
	// aSq + bSq = cSq
	b := a + 1.0
	for {
		// if a + b + c = 1000, then c = 1000 - a + b
		c := 1000.0 - (a + b)
		if b < c { // a < b < c
			// aSq + bSq = cSq
			sab := (a * a) + (b * b)
			sc := (c * c)
			if sab == sc {
				fmt.Printf("FOUND IT! a:%+v, b:%+v, c:%+v\n", a, b, c)
				fmt.Printf("Answer: %f", a*b*c)
				return
			}
			b = b + 1.0
			continue
		} else {
			if a > 333 { // a < b < c, c == 1000, a < (c/3)
				fmt.Println("something is wrong\n")
				return
			}
			calc(a + 1.0)
			break
		}
	}
}
