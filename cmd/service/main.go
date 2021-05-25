package main

import (
	// "go-epccpayment/common/config"
	"context"
	"go-epccpayment/common/log"
	"go-epccpayment/common/server"
	"go-epccpayment/service"
)

func main() {
	log.Info("go-epccpayment is running!")
	cxt, err := server.Start(context.Background(), service.Start, "/epcc", "localhost", "9051")
	if err != nil {
		log.Info(err.Error())
	}
	<-cxt.Done()
}
