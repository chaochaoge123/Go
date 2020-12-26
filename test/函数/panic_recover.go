package main


import (
	"fmt"
)

func ax (){
	fmt.Println("func ax")
}
func b () {
	defer func() {
	err := recover()
	if err != nil {
		fmt.Println("func b error")
		}
	}()
	panic("func b")
}
func c (){
	fmt.Println("func c")
}

func main(){
	ax()
	b()
	c()

	// recover()必须搭配defer使用
	// defer一定要在可能引发panic的语句之前定义
}