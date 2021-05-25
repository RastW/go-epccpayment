package service

import (
	"go-epccpayment/common/log"
	"io/ioutil"
	"net/http"
)

func Start() {
	http.HandleFunc("/epcc", Server)
}

func Server(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		message, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Info("server receive err: " + err.Error())
		}
		w.WriteHeader(http.StatusOK)
		log.Info("Receive message: " + string(message))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Info("this request is not post , reject request")
	}
}
