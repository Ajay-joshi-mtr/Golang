package main

import "fmt"

func main() {
	//emp1 := make(map[string]string)
	employee := map[string]string{
		"name":   "Ajay jsohi",
		"class":  "BCA",
		"city":   "Mathura",
		"mobile": "451236985",
	}

	fmt.Println(employee)
	delete(employee, "name")
	fmt.Println(employee)
	fmt.Println(employee["cla"]) //shows 0

	_, ok := employee["class"] // comma ok syntex for interrogate value
	fmt.Println(ok)
	fmt.Println(len(employee))
}
