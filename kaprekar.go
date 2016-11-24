package main

import (
	"flag"
	"fmt"
	"math"
)

const (
	base = 10
)

func FindKrat(min, max int, c chan int) {
	control := make(chan bool)
	for i := min; i <= max; i++ {
		go isKrat(i, c, control)
	}

	for i := min; i <= max; i++ {
		select {
		case <-control:
		case value := <-c:
			fmt.Println(value)
			<-control
		}
	}
	//close(c)
}

func isKrat(candidate int, c chan int, control chan bool) {

	squaredCandidate := math.Pow(float64(candidate), 2)
	var i float64
	defer func() {
		control <- true
	}()
	for i = 0; ; i++ {
		if squaredCandidate/math.Pow(base, i) < 1 {
			return
		}
		first, second := getSplitNumber(squaredCandidate, i)
		if first != 0 && second != 0 {
			if first+second == candidate {
				c <- candidate
			}
		}
	}
}

func getSplitNumber(candidate, power float64) (first, second int) {
	firstfloat, _ := math.Modf(candidate / math.Pow(base, power))
	first = int(firstfloat)
	second = int(candidate) - (first * int(math.Pow(base, power)))
	return
}

func main() {
	var min = flag.Int("min", 0, "")
	var max = flag.Int("max", 50000, "")
	flag.Parse()

	c := make(chan int)

	FindKrat(*min, *max, c)
}
