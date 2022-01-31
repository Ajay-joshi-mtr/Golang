package main

import (
	"fmt"
	"math"
)

func main() {
	/* 1. maximum in array*/
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	max := 0
	for i := 0; i < len(x); i++ {
		if x[i] > max {
			max = x[i]
		}
	}
	fmt.Println(x)
	fmt.Println(max)

	/* 2. program to find largest of 8 digit */
	fmt.Println(math.Pow(10, 8) - 1)

	/* 3. accept input and double it */
	fmt.Print("Enter a number: ")
	var input float64
	//	fmt.Scanf("%f", &input) //accept input from KBD
	input = 2
	output := input * 2
	fmt.Println(output)

	/* program to computes average of a series of numbers */
	z := []float64{55, 10, 20, 30, 45, 99, 66, 33}
	total := 0.0
	for _, v := range z {
		total = total + v
	}
	fmt.Println(total / float64(len(z)))

	/* program to return multple values */
	c, d := f() //called
	fmt.Println(c, d)
	// program scope an var function
	t := 5
	increment := func() int {
		t++
		return t
	}
	fmt.Println(increment())
	fmt.Println(increment())
	//
}

// return multiple values demo function f
func f() (int, int) {
	return 5, 6
}
