package main

import (
	"fmt"
)

var wig int = 344 // 全局变量

func foo()(int,string){
	return 22,"peter"
}
func main() {
	fmt.Println("Google" + "Runoob")
	var age int = 66  //申明变量并初始化,局部变量
	fmt.Println("my age is ", age,wig) 
	

	// 忽略某个接收的值用_代替, 如： x,_:=foo()
	// 在函数内部可以使用 := 申明并初始化变量，name := "jine"
	x,y:=foo()  
	fmt.Println(x,y)

	// 申明常量
	const e = 2.7182

	// 批量申明
	// var(
	// 	a string = "dd"
	// 	b int
	// 	c bool
	// )

}
