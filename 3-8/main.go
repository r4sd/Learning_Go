package main

import "fmt"

//Mydata is structure
type Mydata struct {
	Name string
	Data []int
}

// これは値渡しをするもの次に描く参照渡しとの違いを知るためのもの
func main() {
	ccc := Mydata{
		"CCC",
		[]int{10, 20, 30},
	}
	fmt.Println(ccc)
	ccc = rev(ccc)
	fmt.Println(ccc)
}
func rev(md Mydata) Mydata {
	od := md.Data
	nd := []int{}
	for i := len(od) - 1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
	return md

}
