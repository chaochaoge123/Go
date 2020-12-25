package main

import (
	"fmt"
)

type main_int int //自定义类型
type main_new_string = string //类型别名

func type_name()  {
	var a main_int = 88
	var b main_new_string = "sda"
	fmt.Printf("a类型为%T,b类型为%T\n",a,b) //a类型为main.main_int,b类型为string
}

//结构体
type person struct {
	name string
	city bool
	age  int8
}

type bookinfo struct {
	name string;
	page int;
	is_open bool
}

func main(){
	//类型别名与自定义类型
	type_name()

	// 结构体(面向对象的类，通过struct来实现面向对象)
	// 格式 type 类型名 struct {}
	//只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段
	// 定义结构体并实例化
	var p1 person
	p1.name="ddss"
	fmt.Println(p1)  // {ddss false 0}

	// 匿名结构体
	var user struct {name string; age int} 
	user.name="vda"
	fmt.Printf("user类型为%T,user值为%#v\n",user,user)  
	//user类型为struct { name string; age int },user值为struct { name string; age int }{name:"vda", age:0}

	// new 方法对结构体进行实例化
	var book = new(bookinfo)
	book.name="jine"
	fmt.Printf("类型为%T,值为%#v\n",book,book) 
	//类型为*main.bookinfo,值为&main.bookinfo{name:"jine", page:0, is_open:false}


	// 结构体初始化
	book1:=bookinfo{
		name: "pei",
		page: 223,
		is_open: false,
	}
	fmt.Println(book1)
}