package main


func main() {
	/*
	变量：根本上来说，相当于对一块数据存储空间的命名，可以通过定义一个变量来申请一块数据存储空间，之后通过引用变量名来使用这块存储空间
	*/
	///声明，前面需要使用关键字var表明，不需要使用分号结尾
	var v1 int //整形
	var v2 string //字符串
	var v3 [10]int //数组
	var v4 []int //数组切片
	var v5 struct{ //结构体
		f int
	}
	var v6 *int //指针
	var v7 map[string]int //字典
	var v8 func(a int)int //函数
	//同时声明多个变量
	var (
		v9 int
		v10 string
	)

	///初始化
	var v11 int = 10 //很正规的使用方式
	var v12 = 10 //自动推导出类型
	v13 := 10 //自动推导出类型+v13是变量

	///赋值
	v1 = 1
	//多重赋值
	v1, v11 = v11, v1

	///匿名变量
	_, b := GetName()


}

func GetName() (first, second string) {
	return "a", "b"
}
