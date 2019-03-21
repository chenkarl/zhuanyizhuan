package main

import (
	"net/http"

	"./controller"
)

func main() {
	RunServer()
}

func RunServer() {
	handler := controller.MyHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/get_ploygon_weather", &handler)
	server.ListenAndServe()
}
