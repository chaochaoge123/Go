package main


import (
	"fmt"
)
// 结构体
// 创建新的类型要使用type关键字
type student struct {
	name string
	age  int
	gender string
	hobby []string
}

func main() {
	//实例化1
	var peter = student{
		name : "peter",
		age: 19,
		gender: "nan",
		hobby: []string{"篮球","睡觉"},
	}
	fmt.Println(peter)

	var p1 = student{}
	fmt.Println(p1.age)

	//实例化2
	var p2 = new(student)
	fmt.Println(p2.hobby)

	//实例化3
	var p3 = &student{}
	fmt.Println(p3.age)


	//结构体初始化1
	

}