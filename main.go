package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/mainPage.html")
	if err != nil {
		panic(err.Error())
	}

	t.ExecuteTemplate(w, "mainPage", nil)
}

func hnd() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", mainPage)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Started")
	hnd()
}
