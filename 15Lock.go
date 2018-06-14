package main

import (
	"sync"
	"fmt"
)

/*
sync包提供了两种类型的锁：sync.Mutex和sync.RWMutex。
Mutex最简单的一种锁，也比较暴力，当一个goroutine获取到mutex后，其他的goroutine只能乖乖的等待该goroutine释放该mutex
RWMutex比较友善，只会阻止写，不会阻止读
*/

//任何一个lock和rlock均需要保证对应有unlock或者runlock调用与之对应。否则可能导致等待该锁的所有goroutine处于饥饿状态，导致死锁
/*
经典使用模式：
var 1 sync.Mutex
func foo() {
	1.lock()
	defer 1.unlock()
}
*/

/*
全局唯一性操作：once类型保证全局唯一性的操作
*/
var a string
var once sync.Once

func setup() {
	a = "hello, world"
}

//引入了once后，setup只会被调用一次，所有的goroutine在调用该代码时，都会被阻塞，知道全局唯一的once.do调用结束后才继续。
func doprint() {
	once.Do(setup)
	fmt.Print(a)
}

func twoprint() {
	go doprint()
	go doprint()
}


