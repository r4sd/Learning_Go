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
func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

/*
MyDataをDataとして扱う
 Mydataの値を作成しInitalで初期化しPrintDataで出力は変わらず
 実行の仕方が3-14と違っている
 new関数を使うと代入する変数の方に合わせて値が使われる
 今回はob Dataとして定義されているためnewで代入した値はData型として判断される
*/
func main() {
	var ob Data = new(Mydata)
	ob.Initial("EEE", []int{55, 66, 77})
	ob.PrintData()
}
