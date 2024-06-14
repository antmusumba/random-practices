package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
	// usage: go run . --color=blue green
	if len(os.Args[1:]) < 1 {
		fmt.Println("provide argument")
		return
	}
	var ColorPtr =flag.String("color","blue","my color")
	flag.Parse()
	fmt.Println("the color parsed from the command line is : ",*ColorPtr)

	arg := flag.Args()
	fmt.Println("the argument exactly after the flag argument is : ",arg[0])
	

}
