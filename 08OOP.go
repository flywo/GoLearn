package main

import (
	"fmt"
)

func main()  {
	/*
	类型系统：
	1.基础类型，byte,int,bool,float等
	2.复合类型，数组，结构体，指针等
	3.指向任意对象的类型，any
	4.值语义和引用语义
	5.面向对象
	6.接口
	*/

	//使用
	var a Integer = 1
	if a.Less(2) {
		fmt.Println(a, "less 2")
	}
	a.Add(2)
	fmt.Println(a)

	/*数组，结构体，指针都是值语义*/
	var a1 = [3]int{1,2,3}
	var b1 = a1//b1类型为[3]int
	b1[1]++//对b的修改，不会影响a的值
	fmt.Println(a1,b1)

	/*引用语义*/
	var b2 = &a1
	b2[1]++//b2类型为*[3]int
	fmt.Println(a1,*b2)

	//数组切片是一个区间，可将[]T表示为，数组切片类型本身的赋值仍然是值语义
	//type slice struct {
	//	first *T//指向数组的指针
	//	len int
	//	cap int
	//}

	//map本质上是一个字典指针，map[K]V表示：
	//type Map_K_V struct {
	//
	//}
	//type map[K]V struct {
	//	impl *Map_K_V
	//}

	//引用类型：数组切片、map、channel、接口

}

//为类型添加方法
//可以给任意类型添加相应方法
type Integer int//新定义了一个类型Integer
func (a Integer) Less(b Integer) bool {//新增一个Less方法
/*
只有在a需要被修改时，才传入指针
(a *Integer)
后续可以对a进行修改
*a += 1
*/
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}
