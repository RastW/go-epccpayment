package server

import "net/http"

func Server() {
	http.HandleFunc("/epcc", excute)
	http.ListenAndServe("localhost:9051", nil)
}

func excute(w http.ResponseWriter, r *http.Request) {
	return
}
