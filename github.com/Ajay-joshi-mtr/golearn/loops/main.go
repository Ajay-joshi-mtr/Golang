package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}

	}
	//two vars
	for i, j := 1, 1; i <= 5; i, j = i+1, j+1 {
		fmt.Println(i, j)

	}
	//can be redeclare in for loop scope
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	k := 0
	for k <= 5 {
		fmt.Println(k)
		k++
	}
	//infinite break
	x := 0
	for {
		x++
		if x%2 == 0 {
			continue
		}
		fmt.Println(x)

		if x == 15 {
			break
		}
	}
	//Label to break out
Loop:
	for i, j := 1, 1; i <= 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
		if i*j == 9 {
			break Loop
		}

	}
	// work with slice
	s := []int{10, 20, 30, 40, 50}
	for k, v := range s {
		fmt.Println(k, v)
	}
	employee := map[string]string{
		"name":   "Ajay jsohi",
		"class":  "BCA",
		"city":   "Mathura",
		"mobile": "451236985",
	}
	for k, v := range employee {
		fmt.Println(k, v)
	}

	z := "hello world"
	for k, v := range z {
		fmt.Println(k, string(v))
	}
	employee1 := map[string]string{
		"name":   "Ajay jsohi",
		"class":  "BCA",
		"city":   "Mathura",
		"mobile": "451236985",
	}
	// _ ignore key by underscore
	for _, v := range employee1 {
		fmt.Println(v)
	}

}
