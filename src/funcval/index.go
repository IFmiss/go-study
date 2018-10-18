package funcval
// 值传递
func TestVal() {
	// var 
}

func swap(x, y int) int {
	var res int
	res = x
	x = y
	y = res
	return res
}
