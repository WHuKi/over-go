package main

import (
	"log"
	"net/http"
	"over-go/jwt_use"
)

func main() {
	// "Signin"和"Welcome"方法是我们将要实现的处理程序
	http.HandleFunc("/signin", jwt_use.Signin)
	http.HandleFunc("/welcome", jwt_use.Welcome)
	http.HandleFunc("/refresh", jwt_use.Refresh)

	// 在8000端口启动服务
	log.Fatal(http.ListenAndServe(":8000", nil))
}
