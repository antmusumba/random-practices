package main

import (
	"flag"
	"fmt"
)

func main() {
	// declare a pointer variable
	ColorPtr := flag.String("color", "blue", "clouring output")
	flag.Parse()
	fmt.Println(*ColorPtr)
}
