package ifelse

var a = 20
func TestInfo() {
	println("-------------if else --------------")
	if a >= 10 {
		a -= 5
		println("a:", a)
		TestInfo()
	}
}

func TestSwitch(a string) {
	switch a {
		case "daiwei", "111":
			println("daiwei，你好啊")
		case "hahaha":
			println("why you smile")
		default:
			println("需要传一个正常的参数")
	}
}

func TestSwitchType() {
	var x interface{}
	switch i := x.(type) {
		case nil:
			println("x 的类型是: %T", i)
		case int:
			println("x 的类型是: int", i)
		case float64:
			println("x 的类型是: float64", i)
		case string:
			println("x 的类型是: string", i)
		case func(int):
			println("x 的类型是: func(int)型", i)
	}
}