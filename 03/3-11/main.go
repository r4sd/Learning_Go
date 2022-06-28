package main

import "fmt"

// struct
type Mydata struct {
	Name string
	Data []int
}

func (md Mydata) PrintData() { // (md Mydata)  はレシーバーという
	fmt.Println("*** Mydata ***")
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
	fmt.Println("*** End ***")
}

func main() {
	ddd := Mydata{
		"DDD",
		[]int{98, 76, 54, 32, 10},
	}
	ddd.PrintData()
}
