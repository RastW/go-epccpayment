package main

import (
	"go-epccpayment/common/log"
	"go-epccpayment/common/server"
)

func main() {
	log.Info("go-epccpayment is running!")
	server.Server()
}
