package main

import (
	"fmt"
	"log"
	"net/http"
)

var host = "0.0.0.0"
var port = "8080"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	fmt.Printf("start: http://%s", host+":"+port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
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
