package main

import "fmt"

var mydata struct {
	Live string
	Name string
	Data []int
}

func main() {
	mydata.Name = "XXX"
	mydata.Data = []int{10, 20, 30}
	mydata.Live = "Tokyo"
	fmt.Println(mydata) //表示される順番は構造体で設定した順番
}
