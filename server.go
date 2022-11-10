package main

import (
	"asciiart/asciiart"
	"fmt"
	"log"
	"net/http"
	"strings"
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

func errorHandler(w http.ResponseWriter, r *http.Request, d *Error) {
	tmpl = template.Must(template.ParseFiles("error.html"))
	w.WriteHeader(d.Code)
	tmpl.ExecuteTemplate(w, "error.html", d)
}

// make a function that checks if the input contains invalid characters
func valid(input string) bool {
	input = strings.ReplaceAll(input, "\r\n", "")
	for _, c := range input {
		if c < 32 || c > 126 {
			return false
		}
	}
	return true
}
func validFont(font string) bool {
	fonts := []string{"standard", "shadow", "thinkertoy", "300iqfont"}
	for _, f := range fonts {
		if font == f {
			return true
		}
	}
	return false
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	// if r.Method == http.MethodPost && r.URL.Path == "/ascii" {
	r.ParseForm()
	AsciiInput := Input{
		Input: r.FormValue("Input"),
		Font:  r.FormValue("Font"),
	}
	if !validFont(AsciiInput.Font) {
		errorHandler(w, r, &Error{Code: http.StatusBadRequest, Message: "Invalid font"})
		return
	}
	if !valid(AsciiInput.Input) {
		errorHandler(w, r, &Error{Code: http.StatusBadRequest, Message: "Invalid characters"})
		return
	} else {
		tmpl = template.Must(template.ParseFiles("index.html"))
		ascii := asciiart.AsciiPrint(AsciiInput.Input, AsciiInput.Font)
		tmpl.ExecuteTemplate(w, "index.html", ascii)
	}

}

func main() {

	// start the server on port 80 and make sure the css and js files are served and handle all errors
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", asciiHandler)
	fmt.Println("Server started on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// fmt.Printf("Starting server at port %v\n", port)
// http.Handle("/", http.FileServer(http.Dir("./")))
// http.HandleFunc("/ascii-art", asciiHandler)
// if err := http.ListenAndServe(":"+port, nil); err != nil {
// 	fmt.Println("Code: 500\nError starting server.")
// 	log.Fatal(err)
// }
// }
