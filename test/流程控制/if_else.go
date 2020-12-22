package main

import (
	"fmt"
)

func main(){
	var age int = 8
	if age > 18 {
		fmt.Println("成年了")
	} else if age == 18 {
		fmt.Println("刚好18")
	} else {
		fmt.Println("未成年")
	}

}