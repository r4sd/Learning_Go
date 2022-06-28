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
	ccc := new(Mydata)
	fmt.Println(ccc)
	ccc.Name = "CCC"
	ccc.Data = make(
		[]int, //type
		2,     //  長さ
		5,     //　スライス 長さ＜＝スライスであること
	)
	fmt.Println(ccc)
}

// 結果: cccがポインタであることが以下の結果よりわかる&がついている
//
//&{ []}
//&{CCC [0 0 0 0 0]}
