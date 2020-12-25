package main


import (
	"fmt"
)

func test_new(){
	a:=new(int)
	b:=new(string)
	fmt.Println(a,b)  //0xc000012100 0xc0000321f0
	fmt.Printf("a类型为%T,b类型为%T\n",a,b) //a类型为*int,b类型为*string
	fmt.Println(*a,*b)

	// 申明流程
	var c *int  //申明类型
	c = new(int) //分配内存，初始化
	*c = 10 // 赋值操作
}

func main(){
	// 每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置
	a:=888
	b:=&a
	//a:代表被取地址的变量，类型为T
	//b:用于接收地址的变量，ptr的类型就为*T，称做T的指针类型。*代表指针。
	fmt.Printf("类型为%T,指针地址为%p\n",b,b)  //类型为*int,指针地址为0xc000060090 

	// 根据指针去内存取值
	c:=*b
	fmt.Println(c) // 888

	//&取出地址，*根据地址取出地址指向的值

	// new 和 make
	//二者都是用来做内存分配的。
	//make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	//而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针
	test_new()

}