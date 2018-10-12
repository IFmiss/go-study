package mianji

import f "fmt"

var l int
var r int

func GetMianji() {
	f.Println("请输入矩形的长：")
	f.Scanf("%d", &l)
	f.Println("请输入矩形的宽：")
	f.Scanf("%d", &r)
	f.Println("矩形的面积为:", l * r)
}