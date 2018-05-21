package main

import (
	"net/http"
	"strings"
	"time"
	"html/template"
	"log"
)

type PageVariables struct {
	Date string
	Time string
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func renderHomePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	HomePageVars := PageVariables{
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:02:05"),
	}

	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Print("Tempalte parsing error", err)
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("Could not give back the respone", err)
	}
}

func main() {
	http.Handle("/ftp", http.FileServer(http.Dir("~/workspace")))
	http.HandleFunc("/", renderHomePage)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
