package handlers

import (
	"../utils"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ErrorStruct struct {
	Code        int
	Description string
}

var HttpNotFound = ErrorStruct{Code: 404, Description: "Page not found"}
var HttpForbidden = ErrorStruct{Code: 403, Description: "Access denied"}
var HttpInternalServerError = ErrorStruct{Code: 500, Description: "Internal server error"}


func writeCode(error ErrorStruct, w http.ResponseWriter, ctx *utils.AppContext) {
	tmpl, err := template.ParseFiles(ctx.Config.Root + "templates/error.html")
	if err != nil {
		log.Println(err)
		return
	}
	err = tmpl.Execute(w, error)
	if err != nil {
		log.Println(err)
	}
}

func writeHttpCode(code int, w http.ResponseWriter)  {
	log.Println(fmt.Sprintf("Error(code): %d", code))
	w.WriteHeader(code)
	switch code {
	case http.StatusForbidden:
		_, _ = w.Write([]byte("403 Access denied."))
	case http.StatusInternalServerError:
		_, _ = w.Write([]byte("500 Internal server error."))
	case http.StatusNotFound:
		_, _ = w.Write([]byte("404 Not found."))
	case http.StatusBadRequest:
		_, _ = w.Write([]byte("400 Bad request. Check params."))
	}
}
