package main

//　ポインタによるメソッドの話

import (
	"fmt"
	"learning/01/1-7/hello"
	"strconv"
)

type intp int

func (num intp) IsPrime() bool {
	n := int(num)
	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
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

func (num *intp) doPrime() {
	pf := num.PrintFactor()
	// numはポインタなので本来は　*(num).PrintFactor で呼び出していたが
	// ポインタをレシーバーにすることでポインタからメソッドを直接呼び出すことができるようになる
	// 書き方はわかりやすいと思う
	fmt.Println("len", len(pf))
	*num = intp(pf[len(pf)-1]) // 配列の最大要素数を指定
}

func main() {
	s := hello.Input("type a number")
	n, _ := strconv.Atoi(s)
	fmt.Println("n: ", n)
	x := intp(n)
	fmt.Println("============================")
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrintFactor())
	x.doPrime()
	fmt.Println("============================")
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrintFactor())
	x++
	fmt.Println("============================")
	fmt.Printf("%d [%t].\n", x, x.IsPrime())
	fmt.Println(x.PrintFactor())
}
