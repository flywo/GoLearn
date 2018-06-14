package main

import "fmt"

/*
声明：var chanName chan ElementType
例子：var ch chan int   类型为int的channel
var m map[string] chan bool    类型为map，map的元素是bool的channel
ch := make(chan int)
将数据写入channel:ch <- value，会导致程序阻塞，直到有其他goroutine从这个channel中读取数据
读取数据：value := <-ch，如果之前没有写入数据，从channel读取数据也会阻塞程序，直到channel中被写入数据位置。
*/

/*
select:由select开始一个新的选择块，每个选择条件由case语句描述。每个case语句必须是一个IO操作
select {
case <- chan1://如果成功从chan1中读取到数据，执行该case
case chen2 <- 1://如果成功向chen2写入数据
defalut:都没有成功
}
*/

func main() {
	//ch := make(chan int, 1)
	//for {
	//	//随机的选择可以执行的case
	//	select {
	//	case ch <- 0:
	//	case ch <- 1:
	//	}
	//	i := <- ch
	//	fmt.Println("接收到值", i)
	//}

	//缓冲,1024为缓冲区大小,即使无读取方，写入方也可以一直写入数据，缓冲区被填充之前都不会阻塞
	c := make(chan int, 10)
	num := 0
	//读取方式同非缓冲区读取一致，也可以使用range关键字来更为简便循环读取
	for {
		if num >= 10 {
			break
		}
		select {
		case c <- 0:
		case c <- 1:
		case c <- 2:
		case c <- 3:
		case c <- 4:
		case c <- 5:
		}
		num++
	}
	for i := range c {
		fmt.Println(i)
	}
}