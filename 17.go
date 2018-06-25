package main

import (
	"fmt"
	"regexp"

	"github.com/pemcconnell/num2words"
)

type Seventeen struct {
}

var A Seventeen

func main() {
	A.calc()
}

func (f *Seventeen) calc() {
	var cs string
	l := 0
	re := regexp.MustCompile("[^a-zA-Z]")
	for i := 1; i <= 1000; i++ {
		cs = re.ReplaceAllString(num2words.Convert(i), "")
		l += len(cs)
	}
	fmt.Println(l)
	fmt.Println(cs)
	fmt.Println(num2words.Convert(115))
	fmt.Println(re.ReplaceAllString(num2words.Convert(115), ""))
	fmt.Println(len(re.ReplaceAllString(num2words.Convert(115), "")))
}
