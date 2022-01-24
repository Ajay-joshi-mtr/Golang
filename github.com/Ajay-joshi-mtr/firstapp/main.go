package main

import (
	"fmt"
)

const Const int16 = 30

func main() {

	const Const int = 32
	// const Const1 int = math.Sin(15) can't be init
	// Const = 5 can't be assigned to constant
	a := 15 + Const
	fmt.Printf("%v , %T\n", a, a)
	const b = iota // init by 0

	const (
		_ = iota + 5
		j
	)
	const ( //reset to 0
		l = iota
	)

	fmt.Printf("%v , %T\n", j, j)
	fmt.Printf("%v , %T\n", l, l)
}
