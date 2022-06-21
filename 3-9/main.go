package main

import (
	"fmt"
)

//Mydata is structure
type Mydata struct {
	Name string
	Data []int
}

func main() {
	ccc := Mydata{
		"CCC",
		[]int{10, 20, 30},
	}
	fmt.Println(ccc)
	rev(&ccc)
	fmt.Println(ccc)
}

func rev(md *Mydata) {
	od := (*md).Data //これをイメージして使いこなせるか
	nd := []int{}
	for i := len(od) - 1; i >= 0; i-- {
		nd = append(nd, od[i])
		fmt.Println("処理中", nd)
	}
	md.Data = nd
}
