package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Card struct {
	Front string `json:"front"`
	Back  string `json:"back"`
}

type Deck struct {
	Cards []Card
}

var deck = []Card{}

func servemain(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("main.html")
	if err != nil {
		fmt.Println("There was an error:", err)
	}
	b, err := json.MarshalIndent(deck, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t.Execute(w, template.JS(b))
}

func cardform(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("cardadd.html")
	t.Execute(w, nil)
}

func cardadd(w http.ResponseWriter, r *http.Request) {
	var newcard Card
	newcard.Front = r.FormValue("term")
	newcard.Back = r.FormValue("definition")
	deck = append(deck, newcard)
	b, err := json.MarshalIndent(deck, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	os.Stdout.Write(b)
	t, _ := template.ParseFiles("cardadd.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", servemain)
	http.HandleFunc("/cardform", cardform)
	http.HandleFunc("/cardadd", cardadd)
	http.ListenAndServe(":8080", nil)
}
