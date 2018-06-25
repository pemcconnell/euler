package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Eighteen struct {
	Row map[int]map[int]int
}

var A Eighteen

func main() {
	// init
	f, err := os.Open("18.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	row := 0
	A.Row = make(map[int]map[int]int, 0)
	for scanner.Scan() {
		A.Row[row] = make(map[int]int, 0)
		//lines = append(lines, scanner.Text())
		cols := strings.Split(scanner.Text(), " ")
		for k, i := range cols {
			if v, err := strconv.Atoi(i); err == nil {
				A.Row[row][k] = v
			}
		}
		row++
		fmt.Println(scanner.Text())
	}
	// calc
	A.calc()
}

type pointRef struct {
	col int
	row int
	val int
}

func (f *Eighteen) calc() {
	lr := len(f.Row) - 1
	// start from the 2nd last row
	for r := lr - 1; r >= 0; r-- {
		for k, v := range f.Row[r] {
			bv := f.Row[r+1][k]
			if f.Row[r+1][k+1] > bv {
				bv = f.Row[r+1][k+1]
			}
			v += bv
			f.Row[r][k] = v
		}
		delete(f.Row, r+1)
	}
	fmt.Println(f.Row)
}
