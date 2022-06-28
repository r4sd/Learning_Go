package main

import (
	"fmt"
	"learning/01/1-7/hello"
)

func main() {
	name := hello.Input().Input("type your name")
	fmt.Println("Hello, " + name + "!!")
}
