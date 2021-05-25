package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Start(cxt context.Context, funcHandler func(), name, host, port string) (context.Context, error) {
	// http.HandleFunc(name, funcHandler)
	funcHandler()
	var srv http.Server
	srv.Addr = ":" + port

	go func() {
		log.Println(srv.ListenAndServe())
	}()
	go func() {
		var s string
		fmt.Scanln(s)
		// 结束上下文
		srv.Shutdown(cxt)
	}()
	return cxt, nil
}
