// 运算符
package calc

import(
	f "fmt"
)

var a int = 10
var b int = 20

func TestCalc() {
	f.Print("---------- calc ---------\n")

	f.Println("a + b: ", a + b)
	a = a + b
	f.Println("a: ", a)

	f.Println("a - b: ", a - b)
	a = a - b
	f.Println("a: ", a)

	f.Println("a * b: ", a * b)
	a = a * b
	f.Println("a: ", a)

	f.Println("a / b: ", a / b)
	a = a / b
	f.Println("a: ", a)

	f.Println("a % b: ", a % b)
	a = a % b
	f.Println("a: ", a)
}

func BaseInfo() {
	var e uint = 60		/* 60 = 0011 1100 */
	var d uint = 13		/* 13 = 0000 1101 */
	var h uint = 0

	f.Println("e: 0011 1100")
	f.Println("d: 0000 1101")

	h = e & d
	f.Println("e & d: ", h)

	h = e | d
	f.Println("e | d: ", h)

	h = e ^ d
	f.Println("e ^ d: ", h)

	h = e << 2
	f.Println("e << 2: ", h)

	h = e >> 2
	f.Println("e >> 2: ", h)
}

func SetInfo() {
	var a int = 21
	var c int

	c = a
	println("c = a: ", c)

	c = a
	println("c = a: ", c)

	c += a
	f.Println("c +=  a", c)

	c -= a
	f.Println("c -= a", c)

	c *= a
	f.Println("c *= a", c)

	c /= a
	f.Println("c /= a", c)

	c %= a
	f.Println("c %= a", c)

	c = 200

	c <<= 2
	f.Println("c <<= 2", c)

	c >>= 2
	f.Println("c >>= 2", c)

	c ^= 2
	f.Println("c ^= 2", c)

	c |= 2
	f.Println("c |= 2", c)

	c &= 2
	f.Println("c &= 2", c)
}