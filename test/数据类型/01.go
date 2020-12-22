package main
import (
        "fmt"
        "math"
)
func main() {
	// 整型
	var age int = 33
	fmt.Printf("%v,%T\n",age,age)

	// 浮点型
	var hig float32 = 39.7873
	fmt.Printf("%.2f\n",hig)
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.1f\n", math.Pi) //保留一位小数

	//布尔型数据只有true（真）和false（假）两个值
	var a bool = true
	fmt.Println(a)

	//字符串string
	// 多行字符串用反引号
	s1 := `第一行
		   第二行
		   第三行
		`
	s2 := "one,\ntwo,\nthree"
	fmt.Println(s1)
	fmt.Println(s2)

	//byte(代表了ASCII码的一个字符)和rune类型(代表一个 UTF-8字符)
	qw := "hello你好"
	for _,r := range qw {  //遍历包含中文的字符串
		fmt.Printf("%v(%c)\n",r,r)
	}

	// 变更字符串
	var rty string = "sello"
	byterty := []byte(rty)
	byterty[0]='h'
	fmt.Println(string(byterty))

	var yur string = "不好啊ssddd"
	runeyur := []rune(yur)
	fmt.Printf("%v,%T",runeyur,runeyur)
	runeyur[0]='很'
	fmt.Println(string(runeyur))

	// 统计个数
	bbc := "半城烟沙wwe"
	cou := 0
	for num,r := range bbc {
		fmt.Println(num,string(r))
		if len(string(r))>=3{
			cou++
		}	
	}
	fmt.Printf("总长度%d\n",len(bbc))
	fmt.Printf("中文字符个数%d",cou)
}