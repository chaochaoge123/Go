package main

import (
	"fmt"
)

func main(){
	for i:=0; i<10; i++ {
		if i==4 {
			break
		}
		fmt.Println(i)
	}

	//for range 键值循环
	var fgh string = "hellogo"
	for k,v := range fgh{
		fmt.Println(k,string(v))
	}

	// switch case
	var age int = 15
	switch  {
	case age < 25:
		fmt.Println("好好学习")
		fallthrough // 满足条件的case的下一个case
	case age >25 && age < 45:
		fmt.Println("好好工作")
	default:
		fmt.Println("好好活着")
		
	}

	// goto 退出标签
	for a :=0; a<10; a++ {
		if a==5 {
			goto outfor
		}
		fmt.Printf("a的值%d\n",a)
	}
	outfor:
		fmt.Println("退出了")

	// continue 跳过当次循环
	for k,v := range fgh{
		if string(v) == "g" {
			continue
		}
		fmt.Println(k,string(v))
	}

	// 9*9乘法口诀
	for a := 1; a<10; a++ {
		for b :=1 ; b<=a; b ++{
			fmt.Printf("%d * %d = %d ",b,a,a*b)
		}
		fmt.Println()
	}


	



}