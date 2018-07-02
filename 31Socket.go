package main


/*
net包中有一个类型：TCPConn，可以作为客户端和服务端交互的通道，包含两个函数：
Write
Read
TCPConn可以用在客户端和服务端来读写数据。

TCPAddr:表示一个TCP地址信息，包含：
IP PORT Zone

通过ResolveTCPAddr获取一个TCPAddr
*/
func main() {
	/*
	通过net包中的DialTCP函数来建立一个TCP连接，并返回一个TCPConn类型的对象，连接建立时，服务端也会创建出同一个对象。
	*/
}
