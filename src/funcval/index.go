package funcval
// 值传递
func TestVal() {
	var a int = 4
	var b int = 3

	println("值传递值 a:", a)
	println("值传递值 b:", b)

	swap(a, b)

	println("交换之后的 a:", a)
	println("交换之后的 b:", b)
}

func swap(x, y int) int {
	var res int
	res = x
	x = y
	y = res
	return res
}
