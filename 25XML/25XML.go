package main

import (
	"encoding/xml"
	"os"
	"fmt"
	"io/ioutil"
)

type Recurlyservers struct {//首字母必须大写
	XMLName 	xml.Name 	`xml:"servers"` //``内容，用来辅助反射的。解析时，优先读取struct tag，如果没有，就会对应字段名。元素名是大小写敏感的。
	Version 	string 		`xml:"version,attr"`
	Svs 		[]server 	`xml:"server"`
	Description string		`xml:",innerxml"`
}
/*
xml解析到struct，遵循的规则：
1.如果struct的一个字段是string或者[]byte类型且它的tag含有",innerxml"，Unmarshal将会将此字段所对应的元素内所有
内嵌的原始xml累加到此字段上，如上面的Description定义。
2.如果struct中有一个叫做XMLName，且类型为xml.Name字段，解析的时候回保存这个element的名字到该字段，如上面的servers。
*/

type server struct {
	XMLName 	xml.Name	`xml:"server"`
	ServerName	string		`xml:"serverName"`
	ServerIP  	string 		`xml:"serverIP"`
}

type Servers struct {
	XMLName		xml.Name 	`xml:"servers"`
	Version 	string		`xml:"version,attr"`
	Svs			[]ServersServer	`xml:"server"`
}

type ServersServer struct {
	ServerName	string		`xml:"serverName"`
	ServerIP	string		`xml:"serverIP"`
}

func main() {
	//解析XML，通过xml包的Unmarshal函数实现解析xml
	file, err := os.Open("25XML/servers.xml")
	if err!=nil {
		fmt.Println("读文件发生错误:%v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err!=nil {
		fmt.Println("读取发生错误:%v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err!=nil {
		fmt.Println("解析发生错误:%v", err)
		return
	}
	fmt.Println("name:", v.XMLName, "\nversion:", v.Version, "\nsvs:", v.Svs, "\ndescription:", v.Description)

	//生成XML，通过Marshal和MarshalIndent两个函数实现。第二个函数会增加前缀和缩进。
	servers := &Servers{Version: "1"}
	servers.Svs = append(servers.Svs, ServersServer{"1", "192"})
	servers.Svs = append(servers.Svs, ServersServer{"2", "193"})
	output, err := xml.MarshalIndent(servers, "前缀", "缩进")
	if err!=nil {
		fmt.Println("转换发生错误：%v", err)
	}
	fmt.Println("分割线----------")
	os.Stdout.Write([]byte(xml.Header))
	fmt.Println("分割线----------")
	os.Stdout.Write(output)
}
