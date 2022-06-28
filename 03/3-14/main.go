package main

import "fmt"

// interface
type Data interface {
	Initial(name string, data []int)
	PrintData()
}

//Struct
type Mydata struct {
	Name string
	Data []int
}

//method
func (md *Mydata) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

// println all data
func (md *Mydata) Printdata() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

func main() {
	var ob Mydata = Mydata{}
	ob.Initial("EEE", []int{55, 66, 77})
	ob.Printdata()
}

//  発生している軽微のエラーは後日対応
