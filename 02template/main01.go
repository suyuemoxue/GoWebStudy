package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	//解析模板
	t, err := template.ParseFiles("02template/hello.gohtml")
	if err != nil {
		fmt.Println(err)
		return
	}
	//渲染模板
	user := UserInfo{
		Name:   "墨雪",
		Gender: "男",
		Age:    18,
	}

	err = t.Execute(writer, user)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}
