package main

import (
	"fmt"
	"net/http"
	"text/template"

	"frontendmasters.com/courses/go-basics/fem-museum/api"
	"frontendmasters.com/courses/go-basics/fem-museum/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")
	if err != nil {
		fmt.Println("Error reading template file", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = html.Execute(w, data.GetAll())
	if err != nil {
		fmt.Println("Error executing html", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibition", api.Get)
	server.HandleFunc("/api/exhibitions", api.GetAll)
	server.HandleFunc("/api/create-exhibition", api.Post)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":9090", server)
	if err != nil {
		fmt.Println("Error while opening the server: ", err)
	}
}
