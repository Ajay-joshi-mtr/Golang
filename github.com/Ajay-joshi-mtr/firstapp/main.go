package main

import (
	"fmt"
)

func main() {
	//all variables inilise with zero value

	i := 5
	j := 5
	//arithmetic
	fmt.Println(i + j)
	fmt.Println(i - j)
	fmt.Println(i * j)
	fmt.Println(i / j)
	i = 10 //1010
	j = 3  //0011
	//bitwise
	fmt.Println(i & j)  //0010
	fmt.Println(i | j)  //1011
	fmt.Println(i ^ j)  //1001
	fmt.Println(i &^ j) //1000
	//bitshft
	a := 8              //2^3
	fmt.Println(a << 2) //0010 2^3 * 2^2 = 2^5
	fmt.Println(a >> 2) //0010 2^3 / 2^2 = 2^1
	//strings UTF-8
	s1 := "hello world in go!" //aliases for bytes and immutable
	s2 := " Ajay joshi!"
	//can't be modify s2[0] = "A"
	fmt.Println(s1 + s2)
	b := []byte(s1) //uint8
	fmt.Println(b)
	//rune UTF-32
	c := 'a'         //int32 true type alias rune
	var d rune = 'a' //int32 true type alias rune
	fmt.Printf("%v , %T\n", c, c)
	fmt.Printf("%v , %T\n", d, d)

}
