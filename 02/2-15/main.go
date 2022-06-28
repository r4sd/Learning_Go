package main

import (
	"fmt"
	"learning/01/1-7/hello"
	"strconv"
	"strings"
)

func main() {
	x := hello.Input("Input data")
	ar := strings.Split(x, " ")
	t := 0
	for _, v := range ar {
		n, er := strconv.Atoi(v)
		//fmt.Println("n", n)
		//fmt.Println("er", er)
		if er != nil {
			goto err
		}
		t += n
	}
	fmt.Println("Total:", t)
	return
err:
	fmt.Println("Error!!")
}
