package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Server struct {
	ServerName	string `json:"serverName"`//输出小写
	ServerIP 	string
}

type Serverslice struct {
	Servers 	[]Server
}

func main() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	//利用interface{}解析JSON
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err := json.Unmarshal(b, &f)

	if err!=nil {
		fmt.Println("解析发生错误：", err)
	}

	var ss Serverslice
	ss.Servers = append(ss.Servers, Server{"s", "11"})
	ss.Servers = append(ss.Servers, Server{"ss", "22"})
	bb, err := json.Marshal(ss)
	if err!=nil {
		fmt.Println("发生了错误：", err)
	}
	fmt.Println(string(bb))

	sss := Ser{
		ID: 3,
		ServerName: `Go "1.0"`,
		ServerName2: 100,
		ServerIP: `0`,
	}
	ba, _ := json.Marshal(sss)
	os.Stdout.Write(ba)
}


type Ser struct {
	ID int `json:"-"` //不会导出到JSON
	ServerName string `json:"serverName"` //ServerName2的值会进行二次JSON编码
	ServerName2 int `json:"serverName2,string"`
	ServerIP string `json:"serverIP,omitempty"`//如果serverIP为空，则不输出到JSON串中
}