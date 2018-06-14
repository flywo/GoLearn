package main

import "net/http"

/*
net/http包

http.get():请求一个资源
resp, err := http.get("xxxx")
if err!=nil {
	retrun
}
defer resp.body.close()
io.copy(os.stdout, resp.body)

http.post():传入参数：
1.请求的URL
2.要POST的数据资源类型
3.数据的比特流
resp, err := http.post("www./", "image/jpeg", &imageDataBuf)
if err!=nil {
}
if resp.statuscode != http.statusOK {
处理错误
}

http.postForm():实现了标准编码格式的表单提交
resp, err := http.postForm("www", url.Values{"title":{"xxxx"},"content":{"xxxxx"}})

http.head():表明只请求目标url的头部信息。不反悔http body。
resp, err := http.head("www")

(*http.client).Do():如果需要自定义header字段，使用该方法
1.自定义user-agent，而不是默认的go http package
2.传递cookie
req, err := http.NewRequest("get", "www", nil)
req.Header.Add("User-Agent", "Gobook")
client := &http.Client{}
resp, err := client.Do{req}
*/

/*
高级封装
自定义http.Client:前面的方法都是在http.DeaultClient的基础上调用的。
*/
//创建自定义client
func main() {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	resp, err := client.Get("http.www")
	req, err := http.NewRequest("GET", "http", nil)
	req.Header.Add("User-Agent", "out")
	req.Header.Add("If-None-Match", `W/"TheFileEtag"`)
	resp, err := client.Do(req)
}

