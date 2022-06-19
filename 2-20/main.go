package main

import "fmt"

func main() {
	//mapは順序を保証しない（袋のようなもの）キーで値を管理している
	m := map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
	}
	for k, v := range m {
		//kはキー
		//vはキーに紐づく値

		fmt.Println(k+":", v)
	}
}
