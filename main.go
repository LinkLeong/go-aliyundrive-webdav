package main

import (
	"context"
	"flag"
	"fmt"
	"go-aliyun/aliyun"
	"go-aliyun/aliyun/cache"
	"go-aliyun/aliyun/model"
	"go-aliyun/webdav"
	"net/http"
	"strings"
	"time"
)

func init() {
	cache.Init()
}

func main() {
	var addr *string
	var path *string
	var refreshToken *string
	//
	addr = flag.String("addr", "192.168.2.176:8085", "")
	path = flag.String("path", "./", "")
	//refreshToken = flag.String("rt", "a4d7e58c0f7949cb9c88670d9fb00a30", "refresh_token")
	refreshToken = flag.String("rt", "61e9d623b0f147cb8a6e08add70f2b54", "refresh_token")
	flag.Parse()

	//todo 判断

	refreshResult := aliyun.RefreshToken(*refreshToken)

	config := model.Config{
		RefreshToken: refreshResult.RefreshToken,
		Token:        refreshResult.AccessToken,
		DriveId:      refreshResult.DefaultDriveId,
		ExpireTime:   time.Now().Unix() + refreshResult.ExpiresIn,
	}

	fs := &webdav.Handler{
		Prefix:     "/",
		FileSystem: webdav.Dir(*path),
		LockSystem: webdav.NewMemLS(),
		Config:     config,
	}

	//fmt.p

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// 获取用户名/密码
		//username, password, ok := req.BasicAuth()
		//if !ok {
		//	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		//	w.WriteHeader(http.StatusUnauthorized)
		//	return
		//}
		// 验证用户名/密码
		//if username != "user" || password != "123456" {
		//	http.Error(w, "WebDAV: need authorized!", http.StatusUnauthorized)
		//	return
		//}

		// Add CORS headers before any operation so even on a 401 unauthorized status, CORS will work.

		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")

		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if req.Method == "GET" && strings.HasPrefix(req.URL.Path, fs.Prefix) {
			info, err := fs.FileSystem.Stat(context.TODO(), strings.TrimPrefix(req.URL.Path, fs.Prefix))
			if err == nil && info.IsDir() {
				req.Method = "PROPFIND"

				if req.Header.Get("Depth") == "" {
					req.Header.Add("Depth", "1")
				}
			}
		}

		fmt.Println(req.URL)

		fs.ServeHTTP(w, req)
	})

	http.ListenAndServe(*addr, nil)
}
