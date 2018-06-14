package main

import (
	"net/http"
	"fmt"
	"log"
	"time"
)

/*
处理HTTP请求：net/http提供http.ListenAndServe()，在指定的地址进行监听，开启一个HTTP。
func ListenAndServe(addr string, handler Handler) error
用于在指定的TCP网络地址进行监听，然后调用服务器端处理程序来处理传入的链接请求。
*/
func main() {
	//处理请求
	http.Handle("/foo", fooHandler)
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("处理/bar请求")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	//自定义server
	s := &http.Server{
		Addr: ":8080",
		Handler: myHandler,
		ReadTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
		MaxHeaderBytes: 1<<20,
	}
	log.Fatal(s.ListenAndServe())


	/*
	处理https请求：http.ListenAndServeTLS()，在该方法内设置好证书与路径
	*/
	http.Handle("/foo", fooHandler)
	http.ListenAndServeTLS(":10443", "cert.pem","key.pem", nil)
}
