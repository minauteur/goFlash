package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Card struct {
	Front string `json:"front"` 	//tagged JSON for clarity, will ease implementation/addition of new fields (eg "category") for sorting in the future 
	Back  string `json:"back"`
}

type Deck struct {
	Cards []Card 			//secondary struct allowing new "card" items to be stored in slices for serving/transport			
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
	t.Execute(w, template.JS(b))	//escapes stored JSON data within a script element for proper front-end consumption
}

func cardform(w http.ResponseWriter, r *http.Request) {	//serves form for adding custom cards when /cardform handler is invoked
	t, _ := template.ParseFiles("cardadd.html")
	t.Execute(w, nil)
}

func cardadd(w http.ResponseWriter, r *http.Request) {
	var newcard Card
	newcard.Front = r.FormValue("term")		//HTML form input hooks for storage
	newcard.Back = r.FormValue("definition")	//
	deck = append(deck, newcard)
	b, err := json.MarshalIndent(deck, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	os.Stdout.Write(b)				//each new card is written to the console on form submission to confirm proper form function and to test for Marshalling/JSON errors.
	t, _ := template.ParseFiles("cardadd.html")
	t.Execute(w, nil)
}

func main() {						//handlers
	http.HandleFunc("/", servemain)
	http.HandleFunc("/cardform", cardform)
	http.HandleFunc("/cardadd", cardadd)
	http.ListenAndServe(":8080", nil)
}
