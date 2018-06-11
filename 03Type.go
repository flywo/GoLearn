package main

import "math"

func main()  {
	/*
	内置了基础类型：
	bool
	int,byte,uint,uintptr
	float
	complex
	string
	rune
	error
	pointer
	array
	slice
	map
	chan
	struct
	interface
	*/

	//bool
	var v1 bool
	v1 = true
	v2 := (1==2)
	//不支持其他类型赋值，不支持自动或强制类型转换

	//int  int和int32是不同类型，需要类型转换
	var v3 int32 = 1
	var v4 int
	v4 = int(v3)

	//float 表示包含小数的数据，float32等价于c中float，float64等价于double
	//浮点数比较，最好不用==来判断，可能导致结果不稳定

	//string C中，不存在原生字符串类型，通常使用字符数组来表示，并以字符指针来传递
	var str string
	str = "Hello"
	ch := str[2]
	"hello" + "world"
	len("hello")
	//遍历字符串，如果以字节数组的方式遍历
	str = "Hello"
	n := len(str)
	for i:=0; i<n; i++ {
		ch := str[i]//获得的是btyte类型
	}
	//以Unicode字符遍历
	str := "你好"
	for i,ch := range str {
		//获得的是rune类型
	}


	//数组
	//声明
	[32]byte
	[2*N] struct {x,y int}
	[1000]*float32
	[3][5]int
	//数组声明后，长度就不可改变，数组是值传递


	//数组切片
	//基于数组创建
	arr[:]
	arr[:5]
	arr[5:]
	//直接创建
	make([]int, 5)
	make([]int, 5, 10)
	[]int{1,2,3,4,5}
	//cap返回切片分配的空间大小，len返回所存储的元素个数

	//map
	//声明
	//var my map[string] PersonInfo
	//键：字符串 值：自定义
	make(map[string] int)
	make(map[string] int, 100)//指定大小
	map[string] int {
		"1": 1,
		"2": 2,
	}
	value, ok := myMap["123"]
	if ok {
		//有该键
	}
}

