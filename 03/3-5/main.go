package main

import "fmt"

func main() {
	ar := []int{10, 20, 30}
	fmt.Println(ar)
	initial(&ar) // ここでarのポインタを引数にしている
	fmt.Println(ar)
}

func initial(ar *[]int) {
	for i := 0; i < len(*ar); i++ {
		// fmt.Println(len(*ar)) //配列の要素
		(*ar)[i] = 0
		// *ar[i]ではエラーとなる
		//「ar[i]nおポインタにある値」を抽出したいと解釈されてしまうため
		// (*ar)[i]とすることでarのポインタにあるスライスの[i]を指定できる
		// ややこしい

	}
}
