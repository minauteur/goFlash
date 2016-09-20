package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Card struct {
	Front string `json:"front"` //tagged JSON for clarity, will ease implementation/addition  
	Back  string `json:"back"`	//of new fields (eg "category") to implement sorting in the future
}

type Deck struct {
	Cards []Card	//secondary struct allowing new "card" items to be stored in slices for serving/transport			
}

var deck = []Card{}	//global declaration of deck (of type Deck) allows its value to be accessed by functions below

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
	var newcard Card	//defining newcard (of type Card) that will contain our form data
	newcard.Front = r.FormValue("term")				//HTML form input hooks for storage
	newcard.Back = r.FormValue("definition")	//
	deck = append(deck, newcard)	//append form inputs to our globally stored data for template injection
	b, err := json.MarshalIndent(deck, "", "    ")	//format deck for legibility
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	os.Stdout.Write(b)	//each new card is written to the console on submission to confirm proper form function and to test for Marshalling/JSON errors.
	t, _ := template.ParseFiles("cardadd.html")	//re-renders blank form template 
	t.Execute(w, nil)
}

func main() {						//handlers
	http.HandleFunc("/", servemain) //renders main.html template with embedded JS for flipping/cycling through cards injected as formatted JSON
	http.HandleFunc("/cardform", cardform) //renders cardadd.html
	http.HandleFunc("/cardadd", cardadd)	//writes POST containing newcard data, adds it to globally stored deck, and re-renders cardadd.html
	http.ListenAndServe(":8080", nil)	//port info
	
}
