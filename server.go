package main

import (
	"asciiart/asciiart"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const port = "80"

var tmpl *template.Template

type Error struct {
	Code    int
	Message string
}

type Input struct {
	Input string
	Font  string
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method == http.MethodPost && r.URL.Path == "/ascii" {
	r.ParseForm()
	AsciiInput := Input{
		Input: r.FormValue("Input"),
		Font:  r.FormValue("Font"),
	}
	// fmt.Println(Input)
	tmpl = template.Must(template.ParseFiles("index.html"))
	ascii := asciiart.AsciiPrint(AsciiInput.Input, AsciiInput.Font)
	tmpl.ExecuteTemplate(w, "index.html", ascii)
}

func main() {
	fmt.Printf("Starting server at port %v\n", port)
	http.Handle("/ascii", http.FileServer(http.Dir("./")))
	http.HandleFunc("/", asciiHandler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Code: 500\nError starting server.")
		log.Fatal(err)
	}
}
