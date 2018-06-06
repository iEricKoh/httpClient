package main

import (
	"fmt"
	"httpClient"
	"os"
)

func main() {
	curPath, _ := os.Getwd()
	captcha := curPath + "/captcha.jpg"

	file, err := os.Open(captcha)
	defer file.Close()

	//form := &httpClient.Form{
	//	"username": "ericstore",
	//	"password": "cc111111",
	//	"typeid":   "2000",
	//	"softid":   "89392",
	//	"softkey":  "acc058a141b24d8a8398b1b1dbaf57da",
	//	"image":    file,
	//}

	resp, err := httpClient.Post("http://api.ruokuai.com/create.json", &httpClient.Options{
		Header: &httpClient.Header{
			"Content-Type": "multipart/form-data",
		},
		//Form: form,
	})

	fmt.Println("err", err)
	fmt.Println(string(resp.Body))
}
