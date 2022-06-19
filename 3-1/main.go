package main

import "fmt"

func main() {
	n := 123 //var n int = 123
	p := &n  //var p *int = &n
	fmt.Println("number: ", n)
	fmt.Println("pointer: ", p)
	fmt.Println("value:", *p)
}
