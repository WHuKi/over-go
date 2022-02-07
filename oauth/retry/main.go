package main

// auth2.0 use

import (
	"fmt"
	"html/template"
	"net/http"
	"over-go/oauth"
)

type Conf struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

var conf = Conf{
	ClientId:     "43f2dff544a002b3d28b",
	ClientSecret: "560d4a75043f62304f6bde5faa7190af23fc574d",
	RedirectUrl:  "http://localhost:9090/Oauth/redirect",
}

func Hello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	var temp *template.Template
	var err error
	if temp, err = template.ParseFiles("oauth/views/hello.html"); err != nil {
		fmt.Println("读取文件失败，错误信息为:", err)
		return
	}

	// 利用给定数据渲染模板(html页面)，并将结果写入w，返回给前端
	if err = temp.Execute(w, conf); err != nil {
		fmt.Println("读取渲染html页面失败，错误信息为:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/Oauth/redirect", oauth.Oauth) // 这个和 Authorization callback URL 有关
	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println("监听失败，错误信息为:", err)
		return
	}
}
