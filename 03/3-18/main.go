package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Data is inteface for MyData
type Data interface {
	SetValue(vals map[string]string)
	PrintData()
}

// Mydata is structure
type Mydata struct {
	Name string
	Data []int
}

// SetValue Method
func (md *Mydata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	valt := strings.Split(vals["data"], " ")
	vali := []int{}
	for _, i := range valt {
		n, _ := strconv.Atoi(i)
		vali = append(vali, n)
	}
	md.Data = vali
}

// PrintData Method
func (md *Mydata) PrintData() {
	fmt.Println("Name: ", md.Name)
	fmt.Println("Data: ", md.Data)
}

// Youedata
type Yourdata struct {
	Name string
	Mail string
	Age  int
}

// SetValue is Yourdata Method
func (md *Yourdata) SetValue(vals map[string]string) {
	md.Name = vals["name"]
	md.Mail = vals["mail"]
	n, _ := strconv.Atoi(vals["age"])
	md.Age = n
}

// PrintData is Yourdata Method
func (md *Yourdata) PrintData() {
	fmt.Printf("I'm %s. (%d).\n", md.Name, md.Age)
	fmt.Printf("Mail: %s.\n", md.Mail)
}
func main() {
	ob := [2]Data{}
	ob[0] = new(Mydata)
	ob[0].SetValue(map[string]string{
		"name": "FFF",
		"data": "55, 66, 77",
	})
	ob[1] = new(Yourdata)
	ob[1].SetValue(map[string]string{
		"name": "GGG",
		"mail": "abc@def.g",
		"age":  "120",
	})
	for _, d := range ob {
		d.PrintData()
		fmt.Println()
	}
}
