package loop

func Loop1() {
	numbers := [6] int {1, 2, 3, 4, 5, 6}
	c := 5
	println("循环方式1---------")
	for a := 0; a < c; a++ {
		println("a的值为: ", a)
	}

	println("循环方式2---------")
	b := 0
	for b < c {
		println("b的值为: ", b)
		b++
	}

	println("循环方式3---------")
	for index, x := range numbers {
		println("index值为: ", index, "数组的值为: ", x)
	}
}