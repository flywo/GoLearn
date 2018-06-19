package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"
	"html/template"
	"strconv"
	"regexp"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("路劲", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form {
		fmt.Println("键", k)
		fmt.Println("值", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	/*
	request.Form是一个url.Values类型，里面存储的是对应的类似key=value的信息，下面展示了可以对form数据进行的一些操作:
v := url.Values{}
v.Set("name", "Ava")
v.Add("friend", "Jess")
v.Add("friend", "Sarah")
v.Add("friend", "Zoe")
// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
fmt.Println(v.Get("name"))
fmt.Println(v.Get("friend"))
fmt.Println(v["friend"])
	*/
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, err := template.ParseFiles("22FromIN/login.gtpl")
		if err != nil {
			log.Println("解析发生了错误：", err)
		}else {
			log.Println("将文件写入", t.Execute(w, nil))
		}
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断，不会自动解析form，需要手动调用
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		//使用formvalue时无需调用r.parseform，会在内部自动调用
		fmt.Println("username", r.FormValue("username"))
		fmt.Println("password", r.FormValue("password"))

		//转义
		fmt.Println("用户名：", template.HTMLEscapeString(r.Form.Get("username")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))

		//输入判断
		//为空判断
		if len(r.Form["username"][0]) == 0 {
			//为空处理
		}
		//数字
		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			//错误处理
		}
		if getint > 100 {
			//处理
		}
		//正则
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {

		}
		//验证中文
		if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("realname")); !m {

		}
		//英文
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {

		}
		//电子邮箱
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Println("no")
		}else{
			fmt.Println("yes")
		}
		//手机号
		if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {

		}
		//下拉菜单验证
		slice:=[]string{"apple","pear","banana"}
		v := r.Form.Get("fruit")
		for _, item := range slice {
			if item == v {
			}
		}
		//单选按钮
		slice1:=[]string{"1","2"}

		for _, v := range slice1 {
			if v == r.Form.Get("gender") {

			}
		}
		//复选框
		//slice:=[]string{"football","basketball","tennis"}
		//a:=Slice_diff(r.Form["interest"],slice) https://github.com/astaxie/beeku
		//if a == nil{
		//}

		//时间处理
		t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		fmt.Printf("Go launched at %s\n", t.Local())

		//身份证验证
		//验证15位身份证，15位的是全部数字
		if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
		}

		//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
		}
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("监听端口发生错误：",err)
	}
}
