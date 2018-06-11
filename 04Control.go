package main


func main() {
	//条件语句
	if a<5 {
		return 0
	}else {
		return 1
	}
	//在if后，条件语句之前，可以添加变量初始化语句，使用;间隔。
	//有返回值的函数中，不允许将最终的return包含在if else结构中

	//选择语句
	switch 1 {
	case 0:
	case 2:
		fallthrough
	case 3:
	default:
	}
	//switch后表达式可以不需要
	switch {
	case 0<Num && Num<=3:
	default:
	}
	/*
	1.单个case中可以出现多个结果选项
	2.不需要Break
	3.fallthrough会继续执行下一case
	*/

	//循环语句，只支持for
	for i:=0; i<4; i++ {

	}
	//无限循环
	for {
		sum++
		if sum>100 {
			break
		}
	}
	//多重赋值
	for i,j:=0,1; ; {

	}

}
