package main

import (
	"fmt"
	"learning/Go/1-7/hello"
)

func main() {
	name := hello.Input("type your name")
	fmt.Println("Hello, " + name + "!!")
}
