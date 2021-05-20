package server

import (
	"go-epccpayment/common/log"
	"io/ioutil"
	"net/http"
)

func Server() {
	http.HandleFunc("/epcc", excute)
	http.ListenAndServe("localhost:9051", nil)
}

func excute(w http.ResponseWriter, r *http.Request) {
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
