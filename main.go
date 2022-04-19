package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"text/template"
)

var host = "127.0.0.1"
var port = "4500"

const (
	dirTmpl   = "templates"
	dirStatic = "static"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	mux := http.NewServeMux()
	getWd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(filepath.Join(getWd, dirTmpl, dirStatic))
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(getWd, dirTmpl, dirStatic))))

	mux.HandleFunc("/", index)
	mux.Handle("/static/", fs)

	fmt.Printf("start: http://%s\n", host+":"+port)
	go func() {
		log.Fatal(http.ListenAndServe(":"+port, mux))
	}()

	switch <-c {
	case os.Interrupt:
		log.Print("Got SIGINT...")
	case syscall.SIGTERM:
		log.Print("Got SIGTERM...")
	case syscall.SIGINT:
		log.Print("Got SIGINT...")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	userAgent := r.UserAgent()
	query := r.URL.Query().Get("q")

	fmt.Printf("[%s]\n", r.Method)

	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	tmpl, err := template.ParseFiles(filepath.Join(dirTmpl, "index.html"))
	if err != nil {
		return
	}
	v := ViewList{AddrIP: r.RemoteAddr, UserAgent: userAgent, Msg: "Hello from K8S", Query: query}

	err = tmpl.Execute(w, v)
	if err != nil {
		log.Println(err)
		return
	}
}

type ViewList struct {
	UserAgent string
	AddrIP    string
	Msg       string
	Query     string
}
