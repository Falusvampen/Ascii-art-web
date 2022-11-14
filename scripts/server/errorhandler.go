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

// handles 500 errors
func error500(w http.ResponseWriter, r *http.Request) {
	d := Error{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
	errorHandler(w, r, &d)
}

func errorHandler(w http.ResponseWriter, r *http.Request, d *Error) {
	tmpl = template.Must(template.ParseFiles("templates/error.html"))
	w.WriteHeader(d.Code)
	tmpl.ExecuteTemplate(w, "error.html", d)
}

// looks through the input and returns true if it is valid
func validInput(input string) bool {
	input = strings.ReplaceAll(input, "\r\n", "")
	for _, c := range input {
		if c < 32 || c > 126 {
			return false
		}
	}
	return true
}

// Looks through the list of fonts and returns true if the font is valid
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

// uses the validInput and validFont functions to check if the input is valid
func ValidForm(font string, input string) Error {
	var errForm Error
	if !validInput(input) {
		errForm.Code = http.StatusBadRequest
		errForm.Message = "Invalid Input"
		return errForm
	}
	if !validFont(font) {
		errForm.Code = http.StatusBadRequest
		errForm.Message = "Invalid Font"
		return errForm
	}
	return errForm
}
