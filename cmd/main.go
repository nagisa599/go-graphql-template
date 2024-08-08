package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // 各リクエストに対してhandlerが呼ばれる
	fmt.Println("Starting server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, Hello()) // Hello関数からの返り値をレスポンスとして書き出す
}

func Hello() string {
	return "Hello"
}
