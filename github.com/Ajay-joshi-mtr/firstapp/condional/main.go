package main

import (
	"fmt"
)

func main() {
	number := 50
	guess := 105
	if guess < 1 || returnTrue() || guess > 100 {
		fmt.Println("Guess must be between 1 and 100")
	} else {
		if guess < number {
			fmt.Println("To low")
		}
		if guess > number {
			fmt.Println("Too High ")
		}
		if guess == number {
			fmt.Println("You got it")
		}
		fmt.Println(guess < number, guess > number, guess == number)

	}

	//switch case with tag
	switch i := 2 + 6; i {
	case 2, 4, 6:
		fmt.Println("2,4,6")
	case 8:
		fmt.Println("8")
	case 10:
		fmt.Println("10")
	default:
		fmt.Println("Another no.")
	}

	//switch case tag less
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than eqaul to 10")
	case i <= 20:
		fmt.Println("greter then eqal to 20")
	default:
		fmt.Println("Another no.")
	}
	//type switch interface
	var j interface{} = [3]int{}
	switch j.(type) {
	case int:
		fmt.Println("integer")
	case float32:
		fmt.Println("float32")
	case string:
		fmt.Println("string")
	case [3]int:
		fmt.Println("array of three int")
		break // break to exit
		fmt.Println("Break")
	default:
		fmt.Println("Another type")
	}

}

//custom function
func returnTrue() bool {
	fmt.Println("Returning True")
	return true
}
