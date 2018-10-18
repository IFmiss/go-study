package main

import (
	f "fmt"
	"../hello"
	datatype "../type"
	// calc "../calc"
	// utils "../mianji"
	// condition "../ifelse"
	loopinfo "../loop"
)

func main() {
	hello.TestHello()
	f.Println("你好")
	// utils.GetMianji()
	datatype.PrintInfo()

	// calc.TestCalc()
	// calc.BaseInfo()
	// calc.SetInfo()
	// condition.TestInfo()
	// condition.TestSwitch("daiwei")
	// condition.TestSwitch("111")
	// condition.TestSwitchType()
	loopinfo.Loop1()
}