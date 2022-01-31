package main

import "fmt"

type Student struct {
	Id      int
	Name    string
	Cources []string
	School  []string
}

func main() {

	teacher := struct{ name string }{name: "Ajay joshi"}
	fmt.Println(teacher)
	teacher1 := teacher // pointer or copy data from struct type
	teacher1.name = "Ajay"
	fmt.Println(teacher1)
	student := Student{
		Id:      1,
		Name:    "Ajay",
		Cources: []string{"BCA", "MCA"},
	}
	fmt.Println(student)
	fmt.Println(student.Cources[1])
}
