package funcquote

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func TestQuote() {
	add := func(x, y int) {
		println(&x, " + ", &y, " = ", x + y)
	}

	calc := func(x, y int) int {
		return x + y
	}

	var h1 int = 10
	var h2 int = 20
	println("引用传递值之前 h1：", h1)
	println("引用传递值之前 h2：", h2)

	swap(&h1, &h2)

	println("引用传递值之后 h1：", h1)
	println("引用传递值之后 h2：", h2)

	add(1, 2)

	println("x + y = ", calc(222, 312))

	firstNum := BiBao()
	println("匿名函数firstNum第一次执行:", firstNum())
	println("匿名函数firstNum第二次执行:", firstNum())
	println("匿名函数firstNum第三次执行:", firstNum())

	secondNum := BiBao()
	println("匿名函数secondNum第一次执行:", secondNum())
	println("匿名函数secondNum第二次执行:", secondNum())
	println("匿名函数secondNum第三次执行:", secondNum())
}

func BiBao() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}