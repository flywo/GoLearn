package main

import (
	"html/template"
	"os"
)

type Person struct {
	UserName string
	Email []string
	Friends []*Friend
}

type Friend struct {
	Fname string
}

func main() {
	t := template.New("new")
	t, _ = t.Parse("Hello, {{.UserName}}")//{{.}}表示当前对象
	p := Person{UserName: "Alex"}
	t.Execute(os.Stdout, p)

	//嵌套字段{{range}}{{end}}同go中的range，{{with}}{{end}}指当前对象的值，类似上下文的概念
	f1 := Friend{"aa"}
	f2 := Friend{"bb"}
	tt := template.New("firend")
	tt, _ = tt.Parse(`hello {{.UserName}}!
			{{range .Email}}
				an email {{.}}
			{{end}}
			{{with .Friends}}
			{{range .}}
				my friend name is {{.Fname}}
			{{end}}
			{{end}}
			`)
	pp := Person{UserName:"alex", Email:[]string{"11", "22"}, Friends:[]*Friend{&f1, &f2}}
	tt.Execute(os.Stdout, pp)

	//条件处理
	tEmpty := template.New("test")
	tEmpty = template.Must(tEmpty.Parse("{{if ``}} 不会输出. {{end}}\n"))
	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("value")
	tWithValue = template.Must(tWithValue.Parse("{{if `anything`}}会输出。{{end}}\n"))
	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("else")
	tIfElse = template.Must(tIfElse.Parse("{{if `anything`}} if输出 {{else}} else输出 {{end}}\n"))
	tIfElse.Execute(os.Stdout, nil)
}

//func handler(w http.ResponseWriter, r *http.Request) {
//	t := template.New("some template")//创建模板
//	t, _ = t.ParseFiles("tmpl/welcom.html")//解析模板文件
//	user := GetUser() //获取当前用户信息
//	t.Execute(w, user)//执行模板的merger操作
//}
