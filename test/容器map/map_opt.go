package main


import (
	"fmt"
	"sort"
)


func statement_map(){
	var a=make(map[string]int,10)
	fmt.Println(a)
	a["age"]=33
	a["hig"]=24
	fmt.Println(a) //map[age:33 hig:24]
	fmt.Printf("a type is %T\n",a) //a type is map[string]int

	//申明时初始化元素
	userInfo := map[string]string{
		"username": "per",
		"password": "123456",
	}
	fmt.Println(userInfo)  //map[password:123456 username:per]
}

func key_is_exist(){
	user:=map[string]int{"peter":33,"jine":122}
	v,ok:=user["peterr"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("not exist")
	}
}

func for_map(){
	test_map:=map[string]string{"age":"444","name":"mark","address":"上海"}
	for k,v := range test_map {
		fmt.Println(k,v)
	}
	// 只遍历key值
	for k,_ := range test_map {
		fmt.Println(k)
	}
}

func delete_key(){
	test_map:=map[string]string{"age":"444","name":"mark","address":"上海"}
	delete(test_map,"name")
	fmt.Println(test_map) //map[address:上海 age:444]
}

func sort_for_map(){
	// 流程：将map 里面的key值添加到切片中，对切片进行排序，遍历切片，通过key取到对应的value
	test_map:=map[int]string{33:"444",1:"mark",77:"上海"}
	keys:=[]int{}
	for k,_ :=range test_map {
		keys=append(keys,k)
	}
	fmt.Printf("keys的值为%v\n",keys)
	sort.Ints(keys)
	for _,v:= range keys {
		fmt.Println(v,test_map[v])
	}
}

func slice_map(){
	var test_slice=make([]map[string]string,5)
	fmt.Println(test_slice) //[map[] map[] map[] map[] map[]]
	// 对切片中的map进行初始化
	test_slice[0]=make(map[string]string,5) //初始化
	test_slice[2]=make(map[string]string,5)
	test_slice[0]["name"]="gdj"
	test_slice[2]["age"]="223"
	fmt.Println(test_slice) //[map[name:gdj] map[] map[age:223] map[] map[]]
}

func map_slice(){
	var test_map=make(map[string][]string,3)
	fmt.Println(test_map)
	var test_slice=[]string{"33","55","66"}
	test_map["nums"]=test_slice
	fmt.Println(test_map)  // map[nums:[33 55 66]]
}

func main(){
	//map是一种无序的基于key-value的数据结构
	// 申明map, make(map[KeyType]ValueType, [cap])  
	// key值类型，value值类型，cap 容量

	//申明
	statement_map()

	// 判断某个key是否存在
	key_is_exist()

	// 遍历
	for_map()

	//删除key
	delete_key()

	//按顺序遍历map
	sort_for_map()

	//元素为map的切片
	slice_map()
	
	//值为切片类型的map
	map_slice()

}
