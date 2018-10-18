package main

import (
	f "fmt"
	"../hello"
	datatype "../type"
	calc "../calc"
	// utils "../mianji"
)

func main() {
	hello.TestHello()
	f.Println("你好")
	// utils.GetMianji()
	datatype.PrintInfo()

	calc.TestCalc()
	calc.BaseInfo()
	calc.SetInfo()
}