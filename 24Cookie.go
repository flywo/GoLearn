package main

import (
	"time"
	"net/http"
	"fmt"
)

func main() {
	//写入cookie
	expiration := time.Now()
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
	http.SetCookie(w, &cookie)//w表示需要写入的response

	//读取cookie
	cookie, _ = r.Cookie("username")
	fmt.Println(w, cookie)

	//或者
	for _, cookie := range r.Cookies() {
		fmt.Println(w, cookie.Name)
	}
}


