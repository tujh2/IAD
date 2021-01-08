package handlers

import (
	"../utils"
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
