package funcs

import "reflect"

func MaxInfo(a, b int) int {
	var res int

	println(reflect.TypeOf(a))
	println(reflect.TypeOf(b))
	println(reflect.TypeOf(a).Elem)

	if a > b {
		res = a
	}

	if b > a {
		res = b
	}

	if b == a {
		res = b
	}
	return res
}

func Swap(x, y, j string) (string, string, string){
	return x, y, j
}
