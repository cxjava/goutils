package httputils

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
)

//添加消息头
func AddRequestHeader(request *http.Request, headers map[string]string) {
	for k, v := range headers {
		request.Header.Set(k, v)
	}
}

func NewGETRequest(targetUrl string) *http.Request {
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		fmt.Println("NewGETRequest - http.NewRequest:", err)
		return nil
	}
	return req
}

//读取响应
func ReadBody(resp *http.Response) string {
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err := gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		defer reader.Close()
		bodyByte, err := ioutil.ReadAll(reader)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return string(bodyByte)
	default:
		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return ""

		}
		return string(bodyByte)
	}
}
