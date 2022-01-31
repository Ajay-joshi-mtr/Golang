package main

import "fmt"

func main() {
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)

	grades := [...]int{11, 22, 33}
	gradesa := &grades
	// gradesa := &grades  refrenced
	gradesa[0] = 500
	fmt.Println(grades)
	fmt.Println(gradesa)
	fmt.Printf("Length: %v\n", len(gradesa))
	fmt.Printf("Capacity: %v\n", cap(gradesa))
	//[] slice [...] array
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := a[:]   //slice all element
	c := a[3:]  //slice from 4th element
	d := a[:4]  //slice first 4 element
	e := a[3:5] //slice 4th 5th element
	//first index including second index excluding
	a[3] = 42
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// slice with make function
	aa := make([]int, 5, 100) // don't have fixed size in slice
	fmt.Println(aa)
	fmt.Printf("Length: %v\n", len(aa))
	fmt.Printf("Capacity: %v\n", cap(aa))
	cc := []int{}
	cc = append(cc, 1, 2, 3, 4, 5) //append capacity^2
	//cc = append(cc, []int{1, 2, 3, 4, 5}) //can't assign array to slice
	cc = append(cc, []int{1, 2, 3, 4, 5}...) //assign array by spread operator
	fmt.Println(cc)
	fmt.Printf("Length: %v\n", len(cc))
	fmt.Printf("Capacity: %v\n", cap(cc))
	// Stack
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	y := append(x[:2], x[3:]...)
	fmt.Println(y)
	fmt.Println(x)

}
