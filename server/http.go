package main

import "net/http"

func server() {
	http.HandleFunc("/epcc", excute)
	http.ListenAndServe("localhost:9051", nil)
}

func excute(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {
	server()
}
