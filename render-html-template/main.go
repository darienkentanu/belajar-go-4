package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

// type Data struct {
// 	title string
// 	name  string
// }

func main() {
	http.HandleFunc("/", route)

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func route(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Belajar Golang",
		"name":  "Darien",
	}

	// var data = struct {
	// 	title string
	// 	name  string
	// }{title: "Belajar Golang", name: "darien"}

	// var data = Data{"Belajar Golang", "Darien"}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
