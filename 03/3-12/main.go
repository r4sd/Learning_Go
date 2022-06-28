package main

import (
	"fmt"
	"learning/01/1-7/hello"
	"strconv"
)

// 整数を入力すると素数かどうかチェックし素因数分解した内容を出力
// そこから数字を２倍し1を足して再度チェックする

type intp int

func (num intp) IsPrime() bool {
	n := int(num)
	for i := 2; i <= (n / 2); i++ {
		fmt.Println("n%i: ", n%i)
		if n%i == 0 {
			fmt.Println("n/2: ", n/2) //割り算４
			return false
		}
	}
	return true
}
func (num intp) PrintFactor() []int {
	ar := []int{}
	x := int(num)
	n := 2
	for x > n {
		if x%n == 0 {
			x /= n
			ar = append(ar, n)
		} else {
			if n == 2 {
				n++
			} else {
				n += 2
			}
		}
	}
	ar = append(ar, x)
	return ar
}

func main() {
	s := hello.Input("type a number")
	n, _ := strconv.Atoi(s)
	x := intp(n)
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrintFactor())
	x *= 2
	x++
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrintFactor())
}
