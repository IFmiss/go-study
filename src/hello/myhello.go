package hello

import "fmt"

var a uint8 = 244
var b = "hello world"
var a1, b1, c1 = "戴", "伟", "哈哈哈"
// a1, b1, c1 := "戴", "伟", "哈哈哈"
var e1, d1, f1 int
func TestHello() {
	a2, b3, c4 := "戴", "伟", "哈哈哈"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a2)
	fmt.Println(b3)
	fmt.Println(c4)
}