package main

func main() {
	//error接口
	func foo(param int)(n int, err error)

	//调用
	n, err := foo(0)
	if err != nil {

	}else {

	}

	//自定义error
	type PathError struct {
		Op string
		Pth string
		Err error
	}
	func (e *PathError) Error() string {
		return "err"
	}

	//defer语句，保证资源被释放，最后一个defer语句会被最先执行

	//panic：调用该函数，正常的执行流程会被终止，但函数中之前使用defer关键字延迟执行的语句将正常执行，之后将返回。

	//recover：捕获异常

}
