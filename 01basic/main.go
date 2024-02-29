package main

import (
	"fmt"
	"net/http"
	"os"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("./01basic/hello.txt")
	_, err = fmt.Fprintln(w, string(data))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println("启动错误", err)
		return
	}
}
