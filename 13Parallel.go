package main


type Vector []float64

//分配给每个CPU的计算任务
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i<n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 //发信号告诉任务管理者我已经计算完成
}

const NCPU = 16//假设有16核

func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU)//接收CPU完成后的信号

	for i:=0; i<NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}

	//等待CPU完成任务
	for i:=0; i<NCPU; i++ {
		<-c //获取到一个数据，表示一个CPU计算完成
	}
	//所有的都已经计算完成
}