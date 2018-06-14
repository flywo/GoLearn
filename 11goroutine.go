package main

import (
	"fmt"
	"time"
)

/*
goroutine是轻量级线程实现，由Go运行时管理。
*/

func Add(x, y int) {
	fmt.Println(x+y)
}

//使用go关键字，让代码在新的goroutine中并发执行。当调用结束，goroutine也就结束，如果函数有返回值，该返回值会被丢弃。
func main() {
	for i:=0; i<10; i++ {
		go Add(i, i)
	}
	//最终没有打印10次，由于main函数结束后，程序就退出，不会等待goroutine结束。

	//并发单元间最常见的通信模型：共享数据和消息
	//共享数据：多个并发单元保持对同一个数据的引用，实现对该数据的共享。
	//消息机制：go语言采用的通信模型，称为channel。

	//channel只能传递一种类型的值，需要在声明channel时指定。
	//定义10个channel数组
	chs := make([]chan int, 10)
	//将么给channel分配给10个goroutine
	for i:=0; i<10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	//所有goruntine启动完成后，<-ch语句从10个channel中依次读取数据，该操作也是阻塞的
	for _,ch := range(chs) {
		<- ch
	}

	//超时等待函数
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()
	select {
	case <- ch://从ch中读取到数据
	case <-timeout://一直没有从ch中读取到数据，但从timeout中读取到了
	}

	//单向channel
	var ch1 chan int
	var ch2 chan<- int//只写
	var ch3 <-chan int//只读
	//转换
	ch4 := make(chan int)
	ch5 := <-chan int(ch4)
	ch6 := chan<- int(ch4)

	//关闭channel
	x, ok := <- ch//判断是否已经被关闭
}

func Parse(ch <-chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}

func Count(ch chan int) {
	ch <- 1//想对应的channel中写入一个数据。被读取前，这个操作是阻塞的
	fmt.Println("继续")
}
