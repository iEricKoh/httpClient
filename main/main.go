package main

import (
	"fmt"
	"httpClient"
	"net/http"
	"net/http/cookiejar"
)

var cookies []*http.Cookie = []*http.Cookie{}
var jar, _ = cookiejar.New(nil)

func main() {
	_, err := httpClient.Get("https://www.douban.com/accounts/login", &httpClient.Options{
		Jar:    jar,
		Cookie: cookies,
	})

	fmt.Println(err)
	//resp, err := httpClient.Post("http://localhost:8090/api/weibo/message/5175429989", &httpClient.Options{
	//	Form: &httpClient.Form{
	//		"text": "几点了",
	//	},
	//})
}
