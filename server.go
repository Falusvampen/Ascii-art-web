package main

import (
	"asciiart/asciiart"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

const Port = "8080"

var tmpl *template.Template

type Error struct {
	Code    int
	Message string
}

type Input struct {
	Input string
	Font  string
}

var err error

func error500(w http.ResponseWriter, r *http.Request) {
	d := Error{
		Code:    500,
		Message: "Internal Server Error",
	}
	errorHandler(w, r, &d)
}

func errorHandler(w http.ResponseWriter, r *http.Request, d *Error) {
	tmpl = template.Must(template.ParseFiles("templates/error.html"))
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
	if font == "" {
		return true
	}
	fonts := []string{"standard", "shadow", "thinkertoy", "300iqfont"}
	for _, f := range fonts {
		if font == f {
			return true
		}
	}
	return false
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	d := Error{}
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		d.Code = 404
		d.Message = "Page Not Found"
		errorHandler(w, r, &d)
		return
	}
	err = r.ParseForm()
	if err != nil {
		error500(w, r)
		return
	}
	AsciiInput := Input{
		Input: r.FormValue("Input"),
		Font:  r.FormValue("Font"),
	}
	if !validFont(AsciiInput.Font) {
		errorHandler(w, r, &Error{Code: http.StatusBadRequest, Message: "Invalid font"})
		return
	} else if !valid(AsciiInput.Input) {
		errorHandler(w, r, &Error{Code: http.StatusBadRequest, Message: "Invalid characters"})
		return
	} else {
		nice, err := template.ParseFiles("templates/index.html")
		if err != nil {
			error500(w, r)
			return
		}
		tmpl = template.Must(nice, err)
		ascii := asciiart.AsciiPrint(AsciiInput.Input, AsciiInput.Font)
		tmpl.ExecuteTemplate(w, "index.html", ascii)
	}
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", asciiHandler)
	fmt.Println("Server started on port", Port)
	err := http.ListenAndServe(":"+Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
