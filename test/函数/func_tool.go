package main



import (
	"fmt"
)

func addnum(a int,b int)int{
	dfg:=a+b
	return dfg

}

func test_one(x int,y ... int) int {  //接收两个参数，返回一个值
	num:=x
	for _,v := range y {
		num+=v
	}
	return num
}

func test_two(x int,y int) (int,int) {  //接收两个参数，返回两个值
	var cfy=x+y
	var vbo=x*y
	return cfy,vbo
}

// 定义一个函数，返回值也是函数
func axf() func(){
	name := "peter"
	return func(){
		fmt.Println("BBBBBBBBB",name)
	}
}

func name_mony(x string)(int){
	// 输入用户名返回应得的金币
	name:=x
	mony_count:=0
	for _,v := range name {
		if string(v)=="e" || string(v)=="E"{
			mony_count += 1
		} else if string(v)=="i" || string(v)=="I"{
			mony_count += 2
		} else if string(v)=="o" || string(v)=="O"{
			mony_count += 3
		} else if string(v)=="u" || string(v)=="U"{
			mony_count += 4
		}else {
			fmt.Println("无符合条件字母")
		}
	}
	return mony_count
}

func main(){
	fvh:=addnum(55,66)  //调用函数addnum
	fmt.Println(fvh)

	// 可变参数(固定参数放在前面，可变参数是通过切片来实现的)
	vfp:=test_one(50,2,3,6)
	fmt.Println(vfp)

	// 多个返回值
	bvd,bga:=test_two(66,22)
	fmt.Println(bvd,bga) //88 1452

	// 匿名函数
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)


	// 闭包函数=函数+外层变量的引用
	r:=axf()
	r()  //执行a函数中的匿名函数

	// 异常捕获


	// defer 语句，先被defer的语句最后被执行，最后被defer的语句，最先被执行
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)

	var (
		coins = 50
		users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",}
		user_mony = make(map[string]int, len(users))
	)

	for _,v := range users {
		fmt.Println(v)
		// 获取用户对应的金币后
		mony_count:=name_mony(v)
		coins -= mony_count
		user_mony[v]=mony_count
	}
	fmt.Println(user_mony)
	fmt.Printf("剩余金币数量%d",coins)
	

}


// 函数格式
// func 函数名(参数)(返回值){
//     函数体
// }

//函数名：由字母、数字、下划线组成。但函数名的第一个字母不能是数字。
//在同一个包内，函数名也称不能重名。
//参数：参数由参数变量和参数变量的类型组成，多个参数之间使用,分隔。
//返回值：返回值由返回值变量和其变量类型组成，也可以只写返回值的类型，
//多个返回值必须用()包裹，并用,分隔。
//函数体：实现指定功能的代码块。