package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Sixteen struct {
}

var A Sixteen

func main() {
	A.calc()
}

func (f *Sixteen) calc() {
	n := math.Pow(2, 1000)
	s := strings.Split(strconv.FormatFloat(n, 'f', 6, 64), ".")[0]
	t := 0
	fmt.Println(s)
	for _, r := range s {
		i, _ := strconv.Atoi(string(r))
		t += i
	}
	fmt.Println(t)
}
