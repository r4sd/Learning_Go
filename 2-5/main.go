package main

import (
	"fmt"
	"learning/1-7/hello"
	"strconv"
)

func main() {
	x := hello.Input("type a number")
	fmt.Print(x + "は、")
	if n, err := strconv.Atoi(x); err == nil {
		if n%2 == 0 {
			fmt.Println(" 偶数です")
		} else {
			fmt.Println(" 奇数です")
		}
	} else {
		fmt.Println("整数ではありません")
	}
}
