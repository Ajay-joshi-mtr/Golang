package main

import (
	"fmt"
	"log"
)

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			//panic(err) rethrow
		}
	}()
	panic("Something bad Happened!")
}
func main() {

	fmt.Println("Start")
	panicker()
	fmt.Println("end")
	//defer example
	/*
		res, err := http.Get("http://www.google.co.in/robots.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close() //should be defer to prevent exit and prevent to open
		robots, err := ioutil.ReadAll((res.Body))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", robots)
	*/

}
