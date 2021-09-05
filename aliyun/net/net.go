package net

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Post(url, token string, data []byte) []byte {
	method := "POST"
	client := &http.Client{
	}
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
func Get(url, token string) []byte {

	method := "GET"

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	//req.Header.Add("accept", "application/json, text/plain, */*")
	//req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36")
	//req.Header.Add("content-type", "application/json;charset=UTF-8")
	//req.Header.Add("origin", "https://www.aliyundrive.com")
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
