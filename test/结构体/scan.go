package main


import (
	"fmt"
)

// sacn 获取终端输入
func main(){

	var (
		name string
		age int
	)
	//fmt.Scan(&name,&age)
	fmt.Scanf("name:%s,age:%d\n",&name,&age)
	fmt.Println(name,age)
}

