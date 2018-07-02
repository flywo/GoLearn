package main

import (
	"os"
	"fmt"
)

/*
文件目录操作大多数在os包里：
Mkdir:创建目录
MkdirAll:创建多级子目录
Remove:删除目录，目录有文件或目录时会出错
RemoveAll:删除目录及子目录
*/
//func main() {
//	os.Mkdir("testFile", 0777)
//	os.MkdirAll("testFile/file1/file2", 0777)
//	err := os.Remove("testFile")
//	if err!=nil {
//		fmt.Println(err)
//	}
//	os.Remove("testFile")
//}


/*
建立与打开文件：
Create:创建新的文件，返回文件对象
NewFile:根据文件描述符创建相应文件，返回文件对象
Open:打开文件，只读
OpenFile:打开文件，flag打开方式，只读读写等，perm是权限

写文件：
Write:写入byte类型的信息到文件
WriteAt:在指定位置开始写入byte类型的信息
WriteString:写入string信息到文件

读文件：
Read:读取数据到B中
ReadAt:从off开始读取数据到b中

删除文件：
Remove:同删除文件夹
*/
func main() {
	userFile := "test.txt"
	fout, err := os.Create(userFile)
	if err!=nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i:=0; i<10; i++ {
		fout.WriteString("Just a test!\n")
		fout.Write([]byte("just a test!\n"))
	}

	fl, err := os.Open(userFile)
	if err!=nil {
		fmt.Println(userFile, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := fl.Read(buf)
		if n==0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}