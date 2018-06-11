package main


func main() {
	//参数类型相同
	add(a, b int) (ret int, err error)
	//返回值类型相同也可
	//只有一个返回值时
	add(a, b int) int

	//不定参数类型
	myFunc(args ...int)
	//...type只能用在函数参数，并且必须是最后一个参数，本质上市一个数组切片，也就是[]type，使用range获取
	for _,arg := range args {

	}

	//不定参数的传递
	func my(args ...int) {
		myfunc(args...)//原样传递
		myfunc(args[1:]...)//任意的切片都可以传进去
	}

	//任意类型的不定参数
	func my(format string, args ...interface{})

	//可以给返回值命名，可以在函数内使用该命名，自动初始化为nil，在遇到不到任何参数的return语句后，会返回对应的变量的值

	//匿名函数：函数可以像变量一样被传递或使用。
	func(a,b int, z float) bool {
		return a*b < int(z)
	}
	//赋值给变量
	f := func(x,y int) int {
		return x+y
	}
	//直接执行
	func(ch chan int) {
		ch <- ACK
	}(reply_chan)




}