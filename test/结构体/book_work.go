package main


import (
	"fmt"
)


type book struct {
	name string
	auth  string
	price int
}

func get_book(){
	var b1=book{}
	fmt.Println(b1)
	input_num()
}
func add_book(){
	var b2=book{}
	var name string
	var auth string
	var price int
	fmt.Scan(&name,&auth,&price)
	b2.name=name
	b2.auth=auth
	b2.price=price
	fmt.Println(b2)

	input_num()
}
func update_book(){

}

func input_num(){
	fmt.Println("1:查看书籍;\n2:增加书籍;\n3:修改书籍;\n4:退出\n")
	var num int
	fmt.Scan(&num)
	if num==1{
		get_book()
	}else if num==2{
		add_book()
	}else if num==3{
		update_book()
	}else{
		fmt.Println("已退出")
	}

}

func main(){
	input_num()

}