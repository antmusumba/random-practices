package main

import "fmt"

func main() {
	go Test("antony")
}

func Test(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
