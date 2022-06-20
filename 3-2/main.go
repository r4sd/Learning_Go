package main

import "fmt"

func main() {
	n := 123
	p := &n
	m := 1000
	p2 := &m
	fmt.Printf("p   value:%d, address:%p\n", *p, p)
	fmt.Printf("p2  value:%d, address:%p\n", *p2, p2)

	pb := p
	p = p2
	p2 = pb
	fmt.Printf("p   value:%d, address:%p\n", *p, p)
	fmt.Printf("p2  value:%d, address:%p\n", *p2, p2)
}
