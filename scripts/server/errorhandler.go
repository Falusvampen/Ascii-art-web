package server

import (
	"net/http"
	"strings"
	"text/template"
)

type Error struct {
	Code    int
	Message string
}

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

func validInput(input string) bool {
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

func ValidForm(font string, input string) Error {
	var errForm Error
	if !validInput(input) {
		errForm.Code = 400
		errForm.Message = "Invalid Input"
		return errForm
	}
	if !validFont(font) {
		errForm.Code = 400
		errForm.Message = "Invalid Font"
		return errForm
	}
	return errForm
}
