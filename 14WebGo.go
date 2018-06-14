package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe("localhost:5208", nil)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if len(req.Form["name"])>0 {
		fmt.Fprint(rw,"Hello,",req.Form["name"][0])
	}else {
		fmt.Fprint(rw,"Hello, world!")
	}
}

func logout(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(rw, "logout")
}
