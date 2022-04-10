package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

var host = "localhost"
var port = "3000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	fmt.Printf("start: http://%s", host+":"+port)
	log.Fatal(http.ListenAndServe(net.JoinHostPort(host, port), mux))
}

func index(w http.ResponseWriter, r *http.Request) {
	userAgent := r.UserAgent()
	query := r.URL.Query().Get("q")

	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, err := w.Write([]byte(fmt.Sprintf("user-agent:%s<br>query:%s", userAgent, query)))
	if err != nil {
		fmt.Println(err)
		return
	}
}
