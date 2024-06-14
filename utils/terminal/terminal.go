package main

import (
	"fmt"
	"strings"
)

func main(){
	s := "hello dope bwoy"
	width := 60
	padding :=(width - len(s)) / 2
	
	fmt.Printf("%-60s\n",s)
	fmt.Printf("%s%s%s\n",strings.Repeat(" ",padding),s,strings.Repeat(" ", width - len(s) - padding))
	fmt.Printf("%60s\n",s)





}