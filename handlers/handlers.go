package handlers

import (
	"ascii-art-web/printascii"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles("templates/index.html", "templates/error.html", "templates/style.css")
	if err != nil {
		log.Fatal(err)
	}
}

type ErrorBody struct {
	Status  int
	Message string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		headError(w, http.StatusNotFound)
		return
	}
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func ProcessorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		headError(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/ascii-art" {
		headError(w, http.StatusNotFound)
		return
	}
	fname := r.FormValue("string")
	f := r.FormValue("font")
	color := r.FormValue("color")
	ascii, err := printascii.AsciiWeb(fname, f)

	if err == printascii.ErrFont || err == printascii.ErrNonAscii || err == printascii.ErrString {
		headError(w, http.StatusBadRequest)
		return
	} else if err == printascii.ErrTxtFile || err == printascii.ErrRead {
		headError(w, http.StatusInternalServerError)
		return
	}
	if !(color == "white" || color == "black" || color == "red" || color == "pink" || color == "blue") {
		headError(w, http.StatusBadRequest)
		return
	}
	d := struct {
		AsciiPrint string
		AsciiColor string
	}{
		AsciiPrint: ascii,
		AsciiColor: color,
	}
	tpl.ExecuteTemplate(w, "index.html", d)
}

func headError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	eh := setError(status)
	tpl.ExecuteTemplate(w, "error.html", eh)
}

func setError(status int) *ErrorBody {
	return &ErrorBody{
		Status:  status,
		Message: http.StatusText(status),
	}
}
