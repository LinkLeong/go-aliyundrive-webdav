package net

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Post(url, token string, data []byte) []byte {
	method := "POST"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	req.Header.Add("content-type", "application/json;charset=UTF-8")
	req.Header.Add("origin", "https://www.aliyundrive.com")
	req.Header.Add("referer", "https://www.aliyundrive.com/")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}
func Put(url, token string, data []byte) []byte {
	method := "PUT"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return body
}
func Get(w http.ResponseWriter, url, token string, rangeStr string, ifRange string) bool {

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return false
	}
	//req.Header.Add("accept", "application/json, text/plain, */*")
	//req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	//req.Header.Add("content-type", "application/json;charset=UTF-8")
	//req.Header.Add("origin", "https://www.aliyundrive.com")
	req.Header.Add("referer", "https://www.aliyundrive.com/")
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("range", rangeStr)
	req.Header.Add("if-range", ifRange)
	fmt.Println(rangeStr)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	}
	io.Copy(w, res.Body)
	res.Body.Close()
	return true

	//	body, err := ioutil.ReadAll(res.Body)
	//	if len(body) == 0 {
	//		fmt.Println("获取详情报错")
	//	}
	//	if err != nil {
	//		fmt.Println(err)
	//		return nil
	//	}
	//	return body
}
func GetProxy(w http.ResponseWriter, req *http.Request, urlStr, token string) []byte {

	//method := "GET"
	u, _ := url.Parse(urlStr)
	proxy := httputil.ReverseProxy{
		Director: func(request *http.Request) {
			request.URL = u
			request.Header.Add("referer", "https://www.aliyundrive.com/")
			request.Header.Add("Authorization", "Bearer "+token)
		},
	}
	proxy.ServeHTTP(w, req)
	//	client := &http.Client{}
	return []byte{}
	//	req, err := http.NewRequest(method, url, nil)
	//
	//	if err != nil {
	//		fmt.Println(err)
	//		return nil
	//	}
	//	//req.Header.Add("accept", "application/json, text/plain, */*")
	//	//req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 //Safari/537.36")
	//	//req.Header.Add("content-type", "application/json;charset=UTF-8")
	//	//req.Header.Add("origin", "https://www.aliyundrive.com")
	//	req.Header.Add("referer", "https://www.aliyundrive.com/")
	//	req.Header.Add("Authorization", "Bearer "+token)
	//
	//	res, err := client.Do(req)
	//	if err != nil {
	//		fmt.Println(err)
	//		return nil
	//	}
	//	defer res.Body.Close()
	//
	//	body, err := ioutil.ReadAll(res.Body)
	//	if len(body) == 0 {
	//		fmt.Println("获取详情报错")
	//	}
	//	if err != nil {
	//		fmt.Println(err)
	//		return nil
	//	}
	//	return body
}
