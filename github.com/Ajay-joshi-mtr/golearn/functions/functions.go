package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello with Go!", i)
	}
	name := "ajay"
	greeting := "hi"
	sayGreeting(&greeting, &name)
	x := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(*x)
	//devide handle
	d, err := devide(5.0, 2.0) //error handle 5.0/0.0
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// for loop
	for i := 0; i < 5; i++ {
		func(i int) { //best practice to pass sync or async
			fmt.Println(i)
		}(i) //ananomous
	}

	//function as types
	var dvd func(float64, float64) (float64, error)
	dvd = func(f1, f2 float64) (float64, error) {
		if f2 == 0.0 {
			return 0.0, fmt.Errorf("Can not devide by Zero")
		} else {
			return f1 / f2, nil
		}
	}
	e, err := dvd(5.0, 2.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(e)

	//methods
	g := greeter{name: "GoLang"}
	g.greet()
	fmt.Println("name:", g.name)

}

func sayMessage(msg string, idx int) {
	fmt.Println("value of index", idx, msg)
}

func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
}

func sum(values ...int) *int {
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	return &result
}

/*
func sum(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}
*/
func devide(x, y float64) (float64, error) {
	if y == 0.0 {
		return 0.0, fmt.Errorf("Can not devide by Zero")
	}
	return x / y, nil
}

//methods

type greeter struct {
	name string
}

//work as unknown context for any type
func (g *greeter) greet() {
	fmt.Println(g.name)
	g.name = ""
}
