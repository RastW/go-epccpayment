package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(ctx context.Context, funcHandler func(), name, host, port string) (context.Context, error) {
	// 执行方法创建接收处理流程
	funcHandler()
	ctx, cancel := context.WithCancel(ctx)

	var srv http.Server
	srv.Addr = ":" + port
	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	go func() {
		fmt.Println(name, "started. Press any key to stop.")
		var s string
		fmt.Scanln(&s)
		// 结束上下文
		srv.Shutdown(ctx)
	}()
	return ctx, nil
}
