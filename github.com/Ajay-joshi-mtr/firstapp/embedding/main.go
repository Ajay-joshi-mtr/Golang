package main

import (
	"fmt"
	"reflect" //reflection
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	Speed  float32
	CanFly bool
}

type Student struct {
	Nme  string `required max:"100"`
	Orgn string
}

//composition relationship
func main() {

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.Speed = 115.5
	b.CanFly = true
	fmt.Println(b)

	a := Bird{
		Animal: Animal{Name: "Emu1", Origin: "Australia1"},
		Speed:  45.5,
		CanFly: true,
	}
	fmt.Println(a)
	//Tagging
	t := reflect.TypeOf(Student{})
	field, _ := t.FieldByName("Nme")
	fmt.Println(field.Tag)

}
