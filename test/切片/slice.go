package main


import (
	"fmt"
)


func statement_slice(){
	var fty = [4]int{3,5,77,89}
	var sdt = [...]string{"衣服","鞋子","外套"}
	fmt.Println(fty)
	fmt.Println(sdt)
}

func slice_opt(){
	var fav=[4]int{33,44,67,14}
	s:= fav[1:3] // 取头不取尾
	s2:=fav[1:]
	fmt.Printf("s:%v,len(s):%v,cap(s):%v\n",s,len(s),cap(s)) //s:[44 67],len(s):2,cap(s):3
	fmt.Printf("s2:%v,len(s2):%v,cap(s2):%v\n",s2,len(s2),cap(s2)) //s2:[44 67 14],len(s2):3,cap(s2):3
}

func make_slice(){
	a:=make([]int,4,10) //参数：类型初始值，长度，容量
	fmt.Printf("a:%v,len(a):%v,cap(a):%v\n",a,len(a),cap(a))
	//a:[0 0 0 0],len(a):4,cap(a):10
}

func update_slice(){
	var kdp=make([]int,4,10)
	s1:=kdp  //将kdp赋值给s1
	s1[0]=100
	fmt.Printf("kdp:%v,s1:%v",kdp,s1) //kdp:[100 0 0 0],s1:[100 0 0 0]
}

func for_slice(){
	var vpf=[...]int{33,44,55,22}
	for i:=0;i<len(vpf);i++ {
		fmt.Println(vpf[i])
	}

	for _,v := range vpf {
		fmt.Println(v)
	}
}


func append_tool() {
	var a []int
	a=append(a,33,44) //添加元素
	fmt.Println(a) //[33 44]

	b:=[]int{66,333,444} // 添加切片
	a=append(a, b...)
	fmt.Println(a) //[33 44 66 333 444]

}

func copy_slice(){
	a:=[]int{33,44,55,66}
	b:=make([]int,5,10)
	copy(b,a)  //将a切片中的数据复制到b切片
	fmt.Printf("a:%v,b:%v",a,b) //a:[33 44 55 66],b:[33 44 55 66 0]
	a[0]=33114
	b[1]=44522
	fmt.Printf("a:%v,b:%v",a,b) //a:[33114 44 55 66],b:[33 44522 55 66 0]
}

func delete_slice(){
	// 通过索引删除
	gvp:=[]int{444,77,242,66}
	gvp=append(gvp[:2],gvp[3:]...)
	fmt.Println("\n",gvp) //删除索引为2的元素

	// 删除指定的元素
	gfd:=[]int{33,77,22,44}
	for i,v := range gfd {
		if v == 44 {
			gfd=append(gfd[:i],gfd[i+1:]...)
			fmt.Println(gfd)
			break
		}
	}
}

func main(){
	//切片是一个引用类型，它的内部结构包含地址、长度和容量。
	//切片一般用于快速地操作一块数据集合
	//通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量
	// 判断切片为空len(a) == 0,切片之间不能直接比较

	//申明切片
	statement_slice()

	//切片操作
	slice_opt()

	// 构造切片
	make_slice()
	
	// 引用更新,对一个切片的修改会影响另一个切片的内容
	update_slice()
	
	// 遍历
	for_slice()

	//追加元素
	append_tool()

	// 复制切片，开辟新的内存地址，更新互不影响
	copy_slice()

	//删除
	delete_slice()

}