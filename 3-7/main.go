package main

import "fmt"

type Mydata struct {
	Name string
	Data []int
}

func main() {
	aaa := Mydata{"AAA", []int{10, 20, 30}} //書籍では},} だが消える
	bbb := Mydata{
		Name: "BBB",
		Data: []int{40, 50, 60},
	}
	fmt.Println(aaa)
	fmt.Println(bbb)
}
