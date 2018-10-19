package funcquote

func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func TestQuote() {
	var h1 int = 10
	var h2 int = 20
	println("引用传递值之前 h1：", h1)
	println("引用传递值之前 h2：", h2)

	swap(&h1, &h2)

	println("引用传递值之后 h1：", h1)
	println("引用传递值之后 h2：", h2)
}