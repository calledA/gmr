package main

import (
	"fmt"
	"gmr/gmr-blog/pkg/setting"
	"gmr/gmr-blog/routers"
	"net/http"

)

func main() {
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("服务器开始启动")
	s.ListenAndServe()
	fmt.Println("服务器关闭")

}
