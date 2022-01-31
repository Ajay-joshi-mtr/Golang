package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(&a, b)
	a = 27
	fmt.Println(a, *b)
	*b = 15
	fmt.Println(a, *b)

	x := []int{1, 2, 3}
	y := &x[1]
	fmt.Println(x, *y)

	var ms *mystruct
	fmt.Println(ms)    // default is nil by compiler
	ms = new(mystruct) //return
	(*ms).foo = 40     //derefrencing
	//fmt.Println((*ms).foo) //ugly access
	fmt.Println(ms.foo) // simple acces

	//array & slices
	p := [3]int{1, 2, 3}
	q := p
	fmt.Println(p, q)
	p[1] = 40
	fmt.Println(p, q)
	//slices use pointer internally
	t := []int{1, 2, 3}
	r := t
	fmt.Println(t, r)
	t[1] = 40
	fmt.Println(t, r)
	/*
		[1 2 3] [1 2 3]
		[1 40 3] [1 2 3]
		[1 2 3] [1 2 3]
		[1 40 3] [1 40 3]
	*/

	//with map
	n := map[string]string{"foo": "bar", "baz": "buz"}
	m := n //copied
	fmt.Println(n, m)
	m["foo"] = "qux"
	fmt.Println(n, m)
	/**
	  map[baz:buz foo:bar] map[baz:buz foo:bar]
	  map[baz:buz foo:qux] map[baz:buz foo:qux]
	*/
}

type mystruct struct {
	foo int
}
