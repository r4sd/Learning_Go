package main

import (
	"fmt"
	"time"
)

func hello(s string, n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("<%d %s>\n", i, s)
		//fmt.Println("Time:", time.Duration(n)*time.Millisecond)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
}

func main() {
	go hello("hello", 50)
	hello("bye!", 100)
}
