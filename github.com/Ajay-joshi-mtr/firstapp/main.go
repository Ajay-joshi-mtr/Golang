package main

import "fmt"

var (
	name   string  = "ajay"
	mobile int     = 45454545
	info   string  = "info"
	j      float32 = 25.5
)
var I int = 50 // exported if capitalize in whole application

func main() {
	i := 55
	var j float32 = 35.5
	k := float32(i) //explicit type conversion
	fmt.Printf("%v, %T\n", name, name)
	fmt.Printf("%v, %T\n", mobile, mobile)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", I, I)
}
