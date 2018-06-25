package main

import "fmt"

type Max struct {
	Sq     int
	Number int
}

type Fourteen struct {
	Answer Max
}

var Q Fourteen

func main() {
	Q.calc(999999)
	fmt.Printf("The highest sequence is %+v by %+v", Q.Answer.Sq, Q.Answer.Number)
}

func (f *Fourteen) calc(startingNumber int) {
	if startingNumber < 1 {
		fmt.Println("over\n")
		return
	} else {
		fmt.Printf("working on %+v\n", startingNumber)
	}
	n := startingNumber
	sq := 1
	for {
		if n == 1 {
			if sq > f.Answer.Sq {
				f.Answer.Sq = sq
				f.Answer.Number = startingNumber
			}
			break
		} else if n%2 == 0 {
			n = f.even(n)
			sq++
		} else {
			sq++
			n = f.odd(n)
		}
	}
	startingNumber--
	f.calc(startingNumber)
}

func (f *Fourteen) odd(n int) int {
	return (n * 3) + 1
}
func (f *Fourteen) even(n int) int {
	return n / 2
}
