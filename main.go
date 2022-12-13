package main

import (
	fonctions "HANGMAN_WEB_BURLE_CRAYSTON_LIERE/Hangman/fonctions"
	"html/template"
	"net/http"
)

type Hangman struct {
	Word    string
	ToFind  string
	Attemps int
}

func main() {
	fs := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	tmpl := template.Must(template.ParseFiles("index.html"))
	data := Hangman{
		ToFind:  fonctions.RevealLetter(fonctions.RandomWord()),
		Attemps: 10,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
