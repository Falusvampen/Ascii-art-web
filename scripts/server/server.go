package server

import (
	"asciiart/scripts/asciiart"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"text/template"
)

var Port = "8080"

var tmpl *template.Template

type Input struct {
	Input string
	Font  string
}

var err error

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	d := Error{}
	// for 404 handling
	// if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
	if r.URL.Path != "/" {
		d.Code = http.StatusNotFound
		d.Message = "Page Not Found"
		errorHandler(w, r, &d)
		return
	}
	// check if there was a problem during parseForm
	err = r.ParseForm()
	if err != nil {
		error500(w, r)
		return
	}

	// take input from form
	AsciiInput := Input{
		Input: r.FormValue("Input"),
		Font:  r.FormValue("Font"),
	}
	// check if the input is valid
	errForm := ValidForm(AsciiInput.Font, AsciiInput.Input)
	if errForm.Code != 0 {
		errorHandler(w, r, &errForm)
		return
	}

	index, err := template.ParseFiles("templates/index.html")
	if err != nil {
		error500(w, r)
		return
	}
	tmpl = template.Must(index, err)
	ascii := asciiart.AsciiPrint(AsciiInput.Input, AsciiInput.Font)
	tmpl.ExecuteTemplate(w, "index.html", ascii)
}

// Start the server
func Start() {
	exec.Command("open", "http://localhost:"+Port).Start()
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.HandleFunc("/", asciiHandler)
	fmt.Println("Server started on port", Port)
	err := http.ListenAndServe(":"+Port, nil)
	if err != nil {
		log.Fatal(err)
	}
}