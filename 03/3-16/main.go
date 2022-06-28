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

//Inital method
func (md *Mydata) Initial(name string, data []int) {
	md.Name = name
	md.Data = data
}

// Check Method
func (md *Mydata) Check() {
	fmt.Printf("Check! Name [%s]\n", md.Name)
}

// println all data
func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

/*
3-15のようにnew関数を使ってもエラーが発生する
インターフェイス方として変数を用意する場合は、その値はインターフェイスに用意された機能しか使えない
とのことだが。。。今はピンときていない
*/
func main() {
	var ob Mydata = Mydata{}
	ob.Initial("EEE", []int{55, 66, 77})
	ob.Check()
}
