package main

import (
	"fmt"
)

const (
	_ = iota // ignore init val

	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {

	filesize := 409625163445

	fmt.Printf("%v GB", filesize/GB)
}
