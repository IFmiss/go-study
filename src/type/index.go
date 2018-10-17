package datatype

var a int = 3;
var b int = a;
var c string = "daiwei"
const (
	a1 = "abc"
	b1 = len(a1)
)

const (
	a2 = iota
	b2
	c2
	d2 = "ha"
	e2
	h2 = 100
	i2
	j2 = iota
	k2
	l2
)

func PrintInfo() {
	b = 4
	println("a的值为%d:", a);
	a, b = b, a
	println("a的值为%d:", a);
	println(a1, b1);
	println(a2, b2, c2, d2, e2, h2, i2, j2, k2, l2);
}

