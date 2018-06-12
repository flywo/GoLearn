package main

func main() {

	//将对象实例赋值给接口，要求改实例实现了接口的所有方法

	var a Integer = 1
	var b LessAdder = &a
}

type Integer int
func (a Integer) Less(b Integer) bool {
	return a < b
}
func (a *Integer) Add(b Integer) {
	*a += b
}
//定义接口
type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}
