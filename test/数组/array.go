package main

import (
	"fmt"
)

func num_tool(){
	bgys := []int{1, 3, 5, 7, 8, 0}
	for k,v := range bgys {
		for i,d := range bgys[k+1:]{
			if v+d==8{
				fmt.Printf("(%d,%d)",k,i+k+1)
			}
		}
	}
}

func main(){
	// 数组申明  var 数组变量名 [元素数量] 类型
	// 数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变
	var array_a [3]int  
	var numArray = [3]int{1, 2} // 使用指定的初始值进行初始化
	fmt.Println(array_a,numArray) //[0 0 0] [1 2 0] 未赋值的默认0

	// 有初始化值可不指定数组长度
	var array_city = [...]string{"南京","东京","海口"}
	fmt.Printf("array_city类型为%T",array_city) //array_city类型为[3]string

	// 通过索引:值的方式初始化数组
	var array_area = [...]int{1:333,3:665}
	fmt.Println("\n",array_area) //[0 333 0 665]

	// 数组遍历
	for _,v := range array_area {
		fmt.Println(v)
	}
	for i :=0; i<len(array_area);i++ {
		fmt.Println(array_area[i])
	}

	// 二维数组
	vfh := [...][2]string{
		{"杭州","苏州"},{"南京","北京"},{"江苏","安徽"},
	}
	for _,v := range vfh {
		for _,s := range v {
			fmt.Printf("%s\t",s)
		}
		fmt.Println()
	}
	
	// 数值取值、变更
	fmt.Println(vfh[1][0]) // 南京
	vfh[1][0]="更新的值"
	fmt.Println(vfh[1][0])  //更新的值

	// 数组[1, 3, 5, 7, 8]所有元素的和
	bgys := []int{1, 3, 5, 7, 8}
	totals:=0
	for _,v := range bgys {
		totals += v
	}
	fmt.Println(totals)
	fmt.Println(bgys[1:])
	
	// 和为8的两个元素的下标
	num_tool()  //(0,3)(1,2)(4,5)
}