package main

import "os"

func main() {
	/*
	常量：在编译期就已知且不可改变的值。可以是数值类型、布尔型、字符串类型
	*/
	///定义，通过关键字const定义
	const Pi float64 = 3.14
	const zero = 0
	const (
		size int64 = 1024
		eof = -1
	)
	const u, v float32 = 0, 3
	const a, b, c = 3, 4, "string"

	//常量的右值可以是编译期运算的常量表达式，比如：
	const mask = 1 << 3
	//不能是任何需要运行期才能知道的结果表达式，比如：
	const Home = os.Getenv("Home")

	//预定义常量true/false/iota
	//iota每使用一次+1,const把iota置为0
	const (
		c0 = iota //0
		c1 //1
		c2 //2
	)

	//枚举，一系列常量，通过const的方式定义枚举值
	const (
		sun = iota
		mon
		tue
		wed
		thu
		fri
		sat
	)
}