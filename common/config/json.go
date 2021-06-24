package main

import (
	// "encoding/json"
	"fmt"
	"os"
)

func Get(x string) {
	// cmap := make(map[string]interface{}, 12)
	json, err := os.ReadFile(x)
	if err != nil {
		fmt.Println("err: " + err.Error())
	}
	fmt.Println(string(json))
}

func main() {
	Get("/Users/wangjiawei/go/src/go-epccpayment/service/config.json")
}
