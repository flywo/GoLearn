package main

import "fmt"

func main() {
	/*
	结构体：只具有组合这个最基础的特性
	*/
	//创建Rect
	rect1 := new(Rect)
	/*
	new背后会执行一个方法
	&Rect{0, 0, 0, 0}
	*/
	rect2 := &Rect{}
	rect3 := &Rect{0,0,100,200}
	rect4 := &Rect{width:100, height:200}
	fmt.Println(rect1.Area(), rect2.Area(), rect3.Area(), rect4.Area())
}

type Rect struct {
	x, y float64
	width,height float64
}

//添加方法
func (r Rect) Area() float64 {
	return r.width*r.height
}

//匿名组合