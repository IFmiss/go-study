package main

import (
	f "fmt"
	"../hello"
	utils "../mianji"
)

func main() {
	hello.TestHello()
	f.Println("你好")
	utils.GetMianji()
}