package main

import (
	"HANGMAN_WEB_BURLE_CRAYSTON_LIERE/Hangman"
	"html/template"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Hangman.Hangman()
		data:=Hangman.HangManData{
			ToFind
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
