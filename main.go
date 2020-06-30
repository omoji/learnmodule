package main

import (
	"fmt"
	"learnmodule/simpleinterest"
)

func main() {
	p := 5000.0
	r := 15.0
	t := 20.0
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("The simple interest is", si)
}
