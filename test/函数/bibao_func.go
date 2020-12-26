package main

import (
	"fmt"
	"strings"
)

func a(name string) func() {
	return func() {
		fmt.Println("hello,", name)
	}
}

func makesuffix(suffix string) func(string) string  {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {  //文件后缀名检测
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}


func main() {
	//闭包函数=函数+外层变量的引用
	r := a("peter")
	r()

	one:=makesuffix(".txt")
	res:=one("makerr.txt")
	fmt.Println(res)  //makerr.txt

	x,y:=calc(100)  //calc 函数传一个参数返回两个函数，返回的函数再传参调用
	ret1 :=x(200)
	fmt.Println(ret1)
	ret2 :=y(200)
	fmt.Println(ret2)
}