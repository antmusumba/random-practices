package main

import "fmt"
func main (){
	const max = 50
	for n := 2 ; n <= max ; n++ {
		isprime := true
		for i := 2 ; i*i <= n ; i++ {
			if n%i == 0 {
				isprime = false
				break
			}
		}
		if !isprime{
			continue
		}
		fmt.Print(n)

	}
}