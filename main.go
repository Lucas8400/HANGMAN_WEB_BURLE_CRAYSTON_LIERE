package main

import (
	fonctions "HANGMAN_WEB_BURLE_CRAYSTON_LIERE/Hangman/fonctions"
	"fmt"
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
	tmpl1 := template.Must(template.ParseFiles("victory.html"))
	tmpl2 := template.Must(template.ParseFiles("defeat.html"))
	data := Hangman{
		Word:    fonctions.RandomWord(),
		ToFind:  fonctions.RevealLetter(fonctions.RandomWord()),
		Attemps: 10,
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		input := r.FormValue("letter")
		if fonctions.VerifyLetter(input, data.Word) && len(input) == 1 {
			data.Attemps--
		}
		var indexes []int
		for index, letter := range data.Word {
			if input == string(letter) {
				if fonctions.VerifyIndex(indexes, index) {
					indexes = append(indexes, index)
				}
			}
		}
		for _, index := range indexes {
			data.ToFind = fonctions.Replace(data.ToFind, input, index)
		}
		if data.Attemps == 0 {
			fmt.Println("Vous avez perdu, le mot était:")
		}
		if data.ToFind == data.Word {
			http.Redirect(w, r, "/victory", http.StatusFound)
			return
		}
		if data.Attemps == 0 {
			http.Redirect(w, r, "/defeat", http.StatusFound)
			return
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("/victory", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			http.Redirect(w, r, "/", http.StatusFound)
			data.Word = fonctions.RandomWord()
			data.ToFind = fonctions.RevealLetter(fonctions.RandomWord())
			data.Attemps = 10
		}
		tmpl1.Execute(w, data)
	})

	http.HandleFunc("/defeat", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			http.Redirect(w, r, "/", http.StatusFound)
			data.Word = fonctions.RandomWord()
			data.ToFind = fonctions.RevealLetter(fonctions.RandomWord())
			data.Attemps = 10
		}
		tmpl2.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}
