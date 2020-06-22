package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	a, b := 22.4, 44.6
	var c float64 = math.Max(a, b)
	fmt.Println(c)
	fmt.Printf("type is %T,size is %d", a, unsafe.Sizeof(a))

	name := "\npeter"
	fmt.Println(name)

}
