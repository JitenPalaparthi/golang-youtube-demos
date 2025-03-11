package handlers

import "net/http"

var Roothandler = func(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Hello World!"))
		w.WriteHeader(http.StatusOK)
		return
	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(501)
		return
	}
	w.Write([]byte("Pong"))
	w.WriteHeader(200)
}

func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(501)
		return
	}
	w.Write([]byte("Ok"))
	w.WriteHeader(200)
}
