# goFlash
A simple flashcard app written in Go and JS.

The project contains three files: 
  cardadd.html is a template containing HTML form elements feeding input to a data struct, allowing users to add their own flashcards to a deck.
  main.html is another template containing the necessary Javascript for serving the user's custom cards on the front-end.
  flash.go contains the back-end code for parsing the templates and serving the user's card data as virtual flashcards.

TODO:
  Add the ability to create, select, and serve multiple decks of cards.
  Add user login/verification for maintaining separate decks on a "per-user" and/or "shared" basis.
