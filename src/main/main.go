package main

import (
	f "fmt"
	"../hello"
	datatype "../type"
	// calc "../calc"
	// utils "../mianji"
	// condition "../ifelse"
	loopinfo "../loop"
	funcinfo "../funcs"
	funcval "../funcval"
	funcquote "../funcquote"
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
	println("最大值为：", funcinfo.MaxInfo(22, 111))

	// 返回多个值
	a, b, c := funcinfo.Swap("Mahesh", "Kumar", "daiwei")
	println(a, b, c)

	funcval.TestVal()
	funcquote.TestQuote()
}